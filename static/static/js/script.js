var mamidApp = angular.module('mamidApp', ['ngRoute']);

mamidApp.config(function ($routeProvider) {
   $routeProvider
       .when('/', {
           templateUrl: 'pages/home.html',
           controller: 'mainController'
       })
       .when('/slaves', {
           templateUrl: 'pages/slaves.html',
           controller: 'slaveIndexController'
       })
       .when('/slave/:slaveId', {
           templateUrl: 'pages/slave.html',
           controller: 'slaveByIdController'
       })
});

mamidApp.controller('mainController', function($scope) {
    $scope.message = 'Greetings from the controller';
});

mamidApp.controller('slaveIndexController', function($scope, $http) {
    $http.get('/api/slaves').
        success(function(data){
            $scope.slaves = data;
    });
});

mamidApp.controller('slaveByIdController', function($scope, $http, $routeParams, $location) {
    $http.get('/api/slave/' + $routeParams['slaveId']).success(
        function(data) {
            $scope.slave = data;
        }
    );
    
    $scope.deleteSlave = function (id) {
        $http.delete('/api/slave/' + id).success(
            function (data) {
                $location.path('#/slaves');
            }
        );
    };
});