package lib

import (
)

type Todo struct {
    Id int                `json:"id"`
    Done bool             `json:"done"`
    Name string           `json:"name"`
}

type Todos []Todo;

var todos = Todos{};

func GetAllTodos () Todos{
    return todos;
}

func CreateNewTodo (todo Todo) bool {

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
            } else if currentTodo.Done == true && todoParams.Done == false{
                currentTodo.Done = false
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
