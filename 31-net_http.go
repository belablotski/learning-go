package main

import (
    "fmt"
    "net/http"
)

func handler1(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "<h1>Handler #1</h1> Relative path: %s", r.URL.Path)
}

func handler2(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "<h1>Handler #2</h1> Relative path: %s", r.URL.Path)
}

func main() {
    http.HandleFunc("/test1/", handler1)
    http.HandleFunc("/test2/", handler2)
    http.ListenAndServe(":8080", nil)
}