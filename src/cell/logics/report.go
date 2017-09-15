// Copyright 2014 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package logics

import (
	"cell/function"
	"cell/protocol"
	"fmt"
	"log"
	"master/settings"
	"sort"
	"strings"
	"time"
)

const (
	barChar = "o"
)

type report struct {
	avgTotal float64
	fastest  float64
	slowest  float64
	average  float64
	rps      float64

	results chan *result
	total   time.Duration

	errorDist      map[string]int
	statusCodeDist map[int]int
	lats           []float64
	sizeTotal      int64

	output string
	taskid int64

	C          int64
	ErrorCount int64
}

func newReport(size int, C int, results chan *result, output string, total time.Duration, taskid int64) *report {
	return &report{
		output:         output,
		results:        results,
		total:          total,
		statusCodeDist: make(map[int]int),
		errorDist:      make(map[string]int),
		taskid:         taskid,
		C:              int64(C),
	}
}

func (r *report) finalize() {
	for {
		select {
		case res := <-r.results:
			if res.err != nil {
				r.ErrorCount++
				r.errorDist[res.err.Error()]++
			} else {
				r.lats = append(r.lats, res.duration.Seconds())
				r.avgTotal += res.duration.Seconds()
				r.statusCodeDist[res.statusCode]++
				if res.contentLength > 0 {
					r.sizeTotal += res.contentLength
				}
			}
		default:
			r.rps = float64(len(r.lats)) / r.total.Seconds()
			r.average = r.avgTotal / float64(len(r.lats))
			r.print()
			r.send2Master()
			return
		}
	}
}

func (r *report) send2Master() {

	sort.Float64s(r.lats)
	ts := r.total.Seconds()
	var sizereq int64 = 0

	if len(r.lats) > 0 {
		r.fastest = r.lats[0]
		r.slowest = r.lats[len(r.lats)-1]
		sizereq = r.sizeTotal / int64(len(r.lats))
	}

	var c2 int64 = 0
	var c3 int64 = 0
	var c4 int64 = 0
	var c5 int64 = 0
	var co int64 = 0

	for code, num := range r.statusCodeDist {
		num64 := int64(num)
		if code < 200 {
		} else if code < 300 {
			c2 += num64
		} else if code < 400 {
			c3 += num64
		} else if code < 500 {
			c4 += num64
		} else if code < 600 {
			c5 += num64
		} else {
			co += num64
		}
	}

	pctls := []int{10, 25, 50, 75, 90, 95, 99}
	data := make([]float64, len(pctls))
	j := 0
	for i := 0; i < len(r.lats) && j < len(pctls); i++ {
		current := i * 100 / len(r.lats)
		if current >= pctls[j] {
			data[j] = r.lats[i]
			j++
		}
	}

	pb := &protocol.Task_Report_C2S{
		Fastest:     &r.fastest,
		Slowest:     &r.slowest,
		Totalsecond: &ts,
		Average:     &r.average,
		Rps:         &r.rps,
		Sizetotal:   &r.sizeTotal,
		Sizereq:     &sizereq,
		Taskid:      &r.taskid,
		Code2:       &c2,
		Code3:       &c3,
		Code4:       &c4,
		Code5:       &c5,
		CodeOther:   &co,
		P10:         &data[0],
		P25:         &data[1],
		P50:         &data[2],
		P75:         &data[3],
		P90:         &data[4],
		P95:         &data[5],
		P99:         &data[6],
		Conc:        &r.C,
		Error:       &r.ErrorCount,
		Cellid:      &CellId,
	}

	log.Println("send report ", pb.Taskid)
	function.NetworkManage.Send(settings.Task_Report_C2S, pb)
}

func (r *report) print() {
	sort.Float64s(r.lats)

	if r.output == "csv" {
		r.printCSV()
		return
	}

	if len(r.lats) > 0 {
		r.fastest = r.lats[0]
		r.slowest = r.lats[len(r.lats)-1]
		fmt.Printf("Summary:\n")
		fmt.Printf("  Total:\t%4.4f secs\n", r.total.Seconds())
		fmt.Printf("  Slowest:\t%4.4f secs\n", r.slowest)
		fmt.Printf("  Fastest:\t%4.4f secs\n", r.fastest)
		fmt.Printf("  Average:\t%4.4f secs\n", r.average)
		fmt.Printf("  Requests/sec:\t%4.4f\n", r.rps)
		if r.sizeTotal > 0 {
			fmt.Printf("  Total data:\t%d bytes\n", r.sizeTotal)
			fmt.Printf("  Size/request:\t%d bytes\n", r.sizeTotal/int64(len(r.lats)))
		}
		r.printStatusCodes()
		r.printHistogram()
		r.printLatencies()
	}

	if len(r.errorDist) > 0 {
		r.printErrors()
	}
}

func (r *report) printCSV() {
	for i, val := range r.lats {
		fmt.Printf("%v,%4.4f\n", i+1, val)
	}
}

// Prints percentile latencies.
func (r *report) printLatencies() {
	pctls := []int{10, 25, 50, 75, 90, 95, 99}
	data := make([]float64, len(pctls))
	j := 0
	for i := 0; i < len(r.lats) && j < len(pctls); i++ {
		current := i * 100 / len(r.lats)
		if current >= pctls[j] {
			data[j] = r.lats[i]
			j++
		}
	}
	fmt.Printf("\nLatency distribution:\n")
	for i := 0; i < len(pctls); i++ {
		if data[i] > 0 {
			fmt.Printf("  %v%% in %4.4f secs\n", pctls[i], data[i])
		}
	}
}

func (r *report) printHistogram() {
	bc := 10
	buckets := make([]float64, bc+1)
	counts := make([]int, bc+1)
	bs := (r.slowest - r.fastest) / float64(bc)
	for i := 0; i < bc; i++ {
		buckets[i] = r.fastest + bs*float64(i)
	}
	buckets[bc] = r.slowest
	var bi int
	var max int
	for i := 0; i < len(r.lats); {
		if r.lats[i] <= buckets[bi] {
			i++
			counts[bi]++
			if max < counts[bi] {
				max = counts[bi]
			}
		} else if bi < len(buckets)-1 {
			bi++
		}
	}
	fmt.Printf("\nResponse time histogram:\n")
	for i := 0; i < len(buckets); i++ {
		// Normalize bar lengths.
		var barLen int
		if max > 0 {
			barLen = counts[i] * 40 / max
		}
		fmt.Printf("  %4.3f [%v]\t|%v\n", buckets[i], counts[i], strings.Repeat(barChar, barLen))
	}
}

// Prints status code distribution.
func (r *report) printStatusCodes() {
	fmt.Printf("\nStatus code distribution:\n")
	for code, num := range r.statusCodeDist {
		fmt.Printf("  [%d]\t%d responses\n", code, num)
	}
}

func (r *report) printErrors() {
	fmt.Printf("\nError distribution:\n")
	for err, num := range r.errorDist {
		fmt.Printf("  [%d]\t%s\n", num, err)
	}
}
