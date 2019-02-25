package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func RegiserHandlers() *httprouter.Router {
	router := httprouter.New()
	router.GET("/videos/:vid-id", streamHeader)
	router.POST("/upload/:vid-id", uploadHeader)
	return router
}
func main() {
	r := RegiserHandlers()
	serve := http.ListenAndServe("9000", r)
}
