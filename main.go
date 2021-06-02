package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	fmt.Println("hello world")

	port := os.Getenv("PORT")

	if port == "" {
		fmt.Println("$PORT must be set")
	}

	http.HandleFunc("/", hello)
	http.ListenAndServe(":"+port, nil)

}

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello SKG Gophers!\n")
}
