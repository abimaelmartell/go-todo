package main

import (
    "net/http"
    "github.com/abimaelmartell/go-todo/lib"
    "os"
)

func main() {
    fs := http.FileServer(http.Dir("static"))
    http.Handle("/", fs)


    // This handles GET and POST and PUT
    http.HandleFunc("/todos", lib.TodoController)

    port := "4567"

    if os.Getenv("PORT") != "" {
        port = os.Getenv("PORT")
    }

    http.ListenAndServe(":" + port, nil)
}
