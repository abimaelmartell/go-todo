(function() {

    var app = angular.module('todoApp', []);

    app.factory('Todo', function($http) {

        var Todo = {};

        Todo.query = function() {
            return $http.get('todos');
        };

        Todo.create = function(name) {
            return $http.post('todos', { name: name });
        }

        Todo.update = function(id, params) {
            return $http.put('todos?id=' + id, params);
        }

        Todo.delete = function(id) {
            return $http.delete('todos?id=' + id);
        }

        return Todo;
    });

    app.controller('MainController', function(Todo, $scope) {
        $scope.fetchTodos = function() {
            Todo.query()
                .success(function(response) {
                    $scope.todos = response;
                });
        };

        $scope.addTodo = function() {
            Todo.create($scope.newTodo)
                .success(function() {
                    $scope.fetchTodos();
                    $scope.newTodo = null;
                });
        }

        $scope.updateTodo = function(id) {
            Todo.update(id, { done: true })
        }

        $scope.deleteTodo = function(id) {
            Todo.delete(id);
        }

        $scope.fetchTodos();
    });

})();
