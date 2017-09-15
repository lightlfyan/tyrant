package main

import (
	"fly"

	"master/midware"

	"master/data"

	"master/function"
	"master/handler"
	_ "master/services"

	"log"
	"os"
	"path/filepath"
	"syscall"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	Pid := syscall.Getpid()
	log.Println("pid: ", Pid)

	defer data.BoltClose()

	function.ProtocolManage.Start()
	data.Start()
	go function.NetworkManage.Start()

	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	fakedir := dir + "/../public/fake"
	if b, _ := exists(fakedir); !b {
		err := os.MkdirAll(fakedir, os.ModePerm)
		if err != nil {
			panic(err)
		}
	}

	router := fly.IWillFly()
	//router.MidWare(midware.Logger, midware.Recovery)
	router.MidWare(midware.CheckLogin)

	router.GET("/static/*filename", handler.StaticHandler)

	router.GET("/", handler.Index)

	router.GET("/show", handler.ShowResult)
	router.GET("/stats", handler.ShowStats)
	router.GET("/chart", handler.ShowChart)
	router.GET("/taskqueue", handler.GetQueueHandler)
	router.GET("/deletetask", handler.DeleteTask)
	router.GET("/deletereport", handler.DeleteReport)

	router.GET("/login", handler.LoginGet)
	router.POST("/login", handler.LoginPost)
	router.POST("/register", handler.Register)

	router.GET("/logout", handler.Logout)

	router.POST("/create", handler.CreatTaskPostHandler)
	router.POST("/delete", handler.DeleteHandler)
	router.POST("/upload", handler.Uploadfile)
	router.GET("/upload", handler.UploadfileGet)

	router.GET("/arch", handler.ARCH)

	log.Println("http on:", "8001")
	fly.Run(router, ":8001")
	//fly.ReloadRun(router, ":8001")
}

func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}
