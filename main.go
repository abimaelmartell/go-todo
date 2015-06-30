package main

import (
    "net/http"
)

func main() {
    fs := http.FileServer(http.Dir("static"))
    http.Handle("/", fs)

    http.ListenAndServe(":4567", nil)
}
