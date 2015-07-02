package lib

import (
    "net/http"
    "encoding/json"
)

func TodoIndex(w http.ResponseWriter, r *http.Request) {
    if r.Method == "GET" {
        todos := GetAllTodos();

        json.NewEncoder(w).Encode(todos)
    }

    if r.Method == "POST" {
        decoder := json.NewDecoder(r.Body)

        var newTodo = Todo{};

        err := decoder.Decode(&newTodo)

        if err != nil {
            panic(err)
        }

        CreateNewTodo(newTodo)
    }
}
