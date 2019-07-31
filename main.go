package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("start")
	//hello
	http.HandleFunc("/hello", hello)
	//health
	http.HandleFunc("/health", health)

	log.Fatal(http.ListenAndServe(":8080", nil))
	close := make(chan bool)
	<-close
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello koderover")
}

func health(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "healthy")
}
