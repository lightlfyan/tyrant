package api

import (
	"cell/logics"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

const (
//fileServer = "http://10.104.108.102:8001/static/fake/"
//fileServer = "http://127.0.0.1:8001/static/fake/"
)

var FileServer string = "http://123.206.183.121:8001/static/fake/"

func Loadfakedata(fakeFilename string) {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Println(err)
		panic(err)
	}

	var l []string

	fakeBytes, err := ioutil.ReadFile(dir + "/" + fakeFilename)
	if err != nil {
		return
	}

	l = strings.Split(string(fakeBytes), "\n")
	for idx, v := range l {
		l[idx] = strings.TrimSpace(v)
	}
	l = Filter(l, func(x string) bool { return x != "" })

	logics.FakeList = make([][]string, 0)

	for _, item := range l {
		l1 := strings.Split(string(item), ",")
		for idx, v := range l1 {
			l1[idx] = strings.TrimSpace(v)
		}
		l1 = Filter(l1, func(x string) bool { return x != "" })
		logics.FakeList = append(logics.FakeList, l1)
	}
}

func Filter(vs []string, f func(string) bool) []string {
	vsf := make([]string, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

func DownloadFakeData(fakefilename string) {
	log.Println("download: ", fakefilename)
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}

	log.Println("download err: ", downloadFile(dir+"/"+fakefilename, FileServer+fakefilename))
}

func downloadFile(filePath string, url string) (err error) {
	// Create the file
	out, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return nil
	}

	// Writer the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}
