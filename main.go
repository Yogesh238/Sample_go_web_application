package main

import (
    "fmt"
    "net/http"
)

fun main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello World from Go Application from main branch 1\n")
    })

    http.ListenAndServe(":9000", nil)
}

