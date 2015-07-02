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

        return Todo;
    });

    app.controller('MainController', function(Todo, $scope) {
        Todo.query()
            .success(function(response) {
            });

        $scope.addTodo = function() {
            Todo.create($scope.newTodo)
        }
    });

})();
