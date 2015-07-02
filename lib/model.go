package lib

import (
)

type Todo struct {
    Done bool           `json:"done"`
    Name string         `json:"name"`
}

type Todos []Todo;

var todos = Todos{};

func GetAllTodos () Todos{
    return todos;
}

func CreateNewTodo (todo Todo) bool {

    todos = append(todos, todo)

    return true
}
