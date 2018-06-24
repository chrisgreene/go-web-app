package main

import (
        "fmt"
        "log"
        "net/http"
)

func main() {
        http.HandleFunc("/", homeHandler)
        log.Fatal(http.ListenAndServe(":8181", nil))
}

func homeHandler(w http.ResponseWriter, request *http.Request) {
        fmt.Fprintln(w, "home.")
}
