(function() {

    var app = angular.module('todoApp', []);

    app.factory('Todo', function($http) {

        var Todo = {};

        Todo.query = function() {
            return $http.get('todos');
        };

        return Todo;
    });

    app.controller('MainController', function(Todo) {
        Todo.query()
            .success(function(response) {
            });
    });

})();
