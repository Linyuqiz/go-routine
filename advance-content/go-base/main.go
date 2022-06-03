package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"unicode/utf8"
)

func main() {
	http.HandleFunc("/", hi)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func hi(w http.ResponseWriter, r *http.Request) {
	// r.Body 只能读取一次
	body, _ := io.ReadAll(r.Body)
	println("request body content: ", string(body))
	_, _ = fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func hello() {
	// bytes strings utf8
	println(len("Hello, world."))
	println(utf8.RuneCountInString("Hello, world."))
}
