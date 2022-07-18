package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	port := os.Getenv("PORT")

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("get request")
		io.WriteString(rw, "This is my website!\n")
	})

	err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
	if err != nil {
		log.Fatal("cannot start server")
	}

}
