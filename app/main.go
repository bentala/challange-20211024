package main

import (
	"challange-20211024/app/entrypoints"
	"fmt"
	"net/http"
)

func getRoutes() {
	http.HandleFunc("/", entrypoints.GetHandlers)

	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			w.Write([]byte("pong"))
		default:
			http.Error(w, "", http.StatusInternalServerError)
		}
	})
}

func main() {
	fmt.Println("server started at localhost:9000")
	getRoutes()
	http.ListenAndServe(":9000", nil)
}
