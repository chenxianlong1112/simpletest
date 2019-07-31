package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("start")
	fmt.Println("add test")
	//hello
	http.HandleFunc("/hello", hello)
	//health
	http.HandleFunc("/health", health)
	//test
	http.HandleFunc("/test", test)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello koderover")
}

func health(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "healthy")
}

func test(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "test")
}
