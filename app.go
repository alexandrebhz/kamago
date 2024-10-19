package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(":80", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<html><body><h1>Hello from kamal and digital ocean! </h1></body></html>")
}
