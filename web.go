package main

import (
	"fmt"
	"net/http"
	"runtime"
)

func main() {
	//http.HandleFunc("/", hello)
	http.HandleFunc("/generate_204", generate204)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		panic(err)
	}
}

func hello(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "hello, world from %s", runtime.Version())
}

func generate204(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(204)
}
