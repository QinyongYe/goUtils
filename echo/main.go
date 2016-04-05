package main

import (
	"net/http"
	"fmt"
)

func SayHello(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte(req.URL.String()))
	for k,v := range req.Header {
		w.Write([]byte(fmt.Sprintf("%s : %s \n", k,v)))
	}

}

func main() {
	http.HandleFunc("/", SayHello)
	http.ListenAndServe(":8001", nil)

}
