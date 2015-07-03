package lib

import (
    "net/http"
    "encoding/json"
    "strconv"
)

func TodoController(w http.ResponseWriter, r *http.Request) {
    if r.Method == "GET" {
        todos := GetAllTodos();

        json.NewEncoder(w).Encode(todos)
    }

    if r.Method == "PUT" {

        decoder := json.NewDecoder(r.Body)

        var todoParams = Todo{};

        _err := decoder.Decode(&todoParams)

        if _err != nil {
            panic(_err)
        }

        UpdateTodo(getTodoId(r), todoParams)
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

    if r.Method == "DELETE" {
        DeleteTodo(getTodoId(r))
    }
}

func getTodoId(r *http.Request) int {
    idStr := r.FormValue("id")

    if idStr != "" {
        id, err := strconv.Atoi(idStr);

        if err != nil {
            panic(err)
        }

        return id
    }
    return 0
}
