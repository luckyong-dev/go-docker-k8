package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	r := httprouter.New()

	r.GET("/hello", getHello)

	server := &http.Server{Addr: ":9090", Handler: r}

	server.ListenAndServe()
}

func getHello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Write([]byte("Hello World"))
}
