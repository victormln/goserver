package main

import (
	"io"
	"net/http"
)

// We put "*" in http.Request to performance reason and to call the pointer and not copy another variable
func handler(writer http.ResponseWriter, request *http.Request) {
	io.WriteString(writer, "Go server listen on :8000")
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/ii", func(writer http.ResponseWriter, request *http.Request) {
		io.WriteString(writer, "You are on /ii")
	})
	// This line is with an anonim function that is the same that "handler" that we created in the top of the file
	/*
		http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
			io.WriteString(writer, "Go server listen on :8000")
		})
	*/
	http.ListenAndServe(":8000", nil)
}
