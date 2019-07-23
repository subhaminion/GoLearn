package main

import (
	"fmt"
	"net/http"
)

func HelloHandler(rw http.ResponseWriter, rq *http.Request) {
	fmt.Fprint(rw, "Hello World!")
}

func PingHandler(rw http.ResponseWriter, rq *http.Request) {
	fmt.Fprint(rw, "Pong!")
}

func main() {
	http.HandleFunc("/ping", PingHandler)
	http.HandleFunc("/", HelloHandler)
	http.ListenAndServe(":8080", nil)
}
