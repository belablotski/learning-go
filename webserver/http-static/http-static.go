package main

	
import (
    "fmt"
    "log"
    "net/http"
    "time"
)

func staticWebServer(path string, port uint) error {
    addr := fmt.Sprintf(":%d", port)
    log.Printf("Starting web-server on address '%s'", addr)
    return http.ListenAndServe(addr, http.FileServer(http.Dir(path)))
}

func main() {
    const MAX_RESTARTS = 100
    for restart_n := 0; restart_n < MAX_RESTARTS; restart_n++ {
        err := staticWebServer("./reports_html", 8080)
        log.Printf("ERROR: %s", err.Error())
        time.Sleep(10 * time.Second)
    }
}