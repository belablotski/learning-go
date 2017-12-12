// http://localhost:8080/test1/
// http://localhost:8080/test2/

package main

import (
  "net/http"
  "html/template"
)

func handler1(w http.ResponseWriter, r *http.Request) {
  tmpl, _ := template.ParseFiles("32-net_http-html_template.html")
  data := struct {Handler string; Url string} {"Handler 1", r.URL.Path}
  tmpl.Execute(w, data)
}

func handler2(w http.ResponseWriter, r *http.Request) {
  tmpl, _ := template.ParseFiles("32-net_http-html_template.html")
  data := struct {Handler string; Url string} {"Handler 2", r.URL.Path}
  tmpl.Execute(w, data)
}

func handlerOthers(w http.ResponseWriter, r *http.Request) {
  http.Error(w, "Unexpected request", http.StatusInternalServerError)	// Replace standard 404 with 500
}

func main() {
  http.HandleFunc("/test1/", handler1)
  http.HandleFunc("/test2/", handler2)
  http.HandleFunc("/", handlerOthers)
  http.ListenAndServe(":8080", nil)
}