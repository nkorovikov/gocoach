package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("get request")
		io.WriteString(rw, "This is my website!\n")
	})

	err := http.ListenAndServe(":3333", nil)
	if err != nil {
		log.Fatal("cannot start server")
	}
}
