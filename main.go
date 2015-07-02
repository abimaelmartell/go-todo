package main

import (
    "net/http"
    "github.com/abimaelmartell/go-todo/lib"
)

func main() {
    fs := http.FileServer(http.Dir("static"))
    http.Handle("/", fs)


    // This handles GET and POST
    http.HandleFunc("/todos", lib.TodoIndex)

    http.ListenAndServe(":4567", nil)

}
