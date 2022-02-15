package main

import (
	"log"
	"net/http"
	"fmt"
)

type server int

func (h *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path)
	w.Write([]byte("Hello beego1"))
	fmt.Println()
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!\n", r.URL.Path[1:])
	w.Write([]byte("Hello beego11111"))
}

func main() {
	//var s server
	http.HandleFunc("/", handler)
	http.ListenAndServe("localhost:8081", nil)
}