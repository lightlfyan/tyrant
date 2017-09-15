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

package main

import (
	"cell/function"
	_ "cell/services"
	"flag"
	"runtime"

	"cell/api"
	"fmt"
	"log"
	"math/rand"
	"time"
)

var addr *string = flag.String("addr", "10.135.45.154", "master address")

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	flag.Parse()
	runtime.GOMAXPROCS(-1)

	rand.Seed(int64(time.Now().UTC().UnixNano()))

	//pid := os.Getpid()
	//filename := fmt.Sprintf("/tmp/cell%d.pid", pid)
	//ioutil.WriteFile(filename, []byte(strconv.Itoa(pid)), 0755)

	api.FileServer = fmt.Sprintf("http://%s:8001/static/fake/", *addr)

	addrWithPort := *addr + ":9999"
	log.Println(api.FileServer, addrWithPort)

	api.DownloadFakeData("file1.txt")
	api.Loadfakedata("file1.txt")

	function.ProtocolManage.Start()

	function.NetworkManage.Start(&addrWithPort)
}
