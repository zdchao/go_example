package main

/**
web服务器
在浏览器中输入:http://127.0.0.1:8080
*/

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handle)
	http.ListenAndServe("localhost:8080", nil)
}

func handle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%q\n", r.URL.Path)
	fmt.Fprintf(w, "sss")
}
