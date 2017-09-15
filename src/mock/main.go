package main

import (
	"mock/cfg"
	"mock/handler"

	"fly"
	//"fly/midware"
)

func main() {
	/*
		file1 := os.NewFile(3, "")
		listenerTCP1, _ := net.FileListener(file1)
	*/

	cfg.Load()

	router := fly.IWillFly()

	//router.MidWare(midware.Logger, midware.Recovery)
	//router.MidWare(midware.Recovery)
	router.GET("/*url", handler.MockHandler)
	router.POST("/*url", handler.MockHandler)
	router.PATCH("/*url", handler.MockHandler)
	router.PUT("/*url", handler.MockHandler)

	fly.ReloadRun(router, ":8889")

	//fly.ReloadRunUseGiveFd(router, listenerTCP1, listenerTCP1.Addr().String())
}
