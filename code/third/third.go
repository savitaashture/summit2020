package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func lastHandler(w http.ResponseWriter, r *http.Request) {
	go test()
	target := os.Getenv("TARGET")
	if target == "" {
		target = "RedHat Summit 2020"
	}
	fmt.Fprintf(w, "response from "+target)
}

func test() {
	for {
		fmt.Println("Running Infinite")
	}
}
func main() {
	log.Print("microservice C started")

	http.HandleFunc("/c", lastHandler)

	port := os.Getenv("TARGET_PORT")
	if port == "" {
		port = "80"
	}

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
