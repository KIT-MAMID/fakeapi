var mamidApp = angular.module('mamidApp', ['ngRoute']);

mamidApp.config(function ($routeProvider) {
   $routeProvider
       .when('/', {
           templateUrl: 'pages/home.html',
           controller: 'mainController'
       })
       .when('/slaves', {
           templateUrl: 'pages/slaves.html',
           controller: 'slavesController'
       })
});

mamidApp.controller('mainController', function($scope) {
    $scope.message = 'Greetings from the controller';
});

mamidApp.controller('slavesController', function($scope, $http) {
    $scope.message = 'List of slaves'
    $http.get('/api/slaves').
        success(function(data){
            $scope.slaves = data
    });
});