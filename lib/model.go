package lib

import (
    "time"
)

type Todo struct {
    Id int                `json:"id"`
    Done bool             `json:"done"`
    Name string           `json:"name"`
    CreatedAt time.Time   `json:"created_at"`
    CompletedAt time.Time `json:"completed_at"`
}

type Todos []Todo;

var todos = Todos{};

func GetAllTodos () Todos{
    return todos;
}

func CreateNewTodo (todo Todo) bool {
    todo.CreatedAt = time.Now();

    if len(todos) == 0 {
        todo.Id = 1
    } else {
        lastTodo := todos[len(todos) - 1]
        todo.Id = lastTodo.Id + 1
    }

    todos = append(todos, todo)

    return true
}

func UpdateTodo (id int, todoParams Todo) bool {

    for i, todo := range todos {
        if todo.Id == id {
            currentTodo := &todos[i]
            if currentTodo.Done == false && todoParams.Done == true {
                currentTodo.Done = true
                currentTodo.CompletedAt = time.Now()
            }

            break;
        }
    }


    return true;
}

func DeleteTodo (id int) bool {
    for i, todo := range todos {
        if todo.Id == id {
            todos = append(todos[:i], todos[i+1:]...)
        }
    }

    return true
}
