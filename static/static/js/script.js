var mamidApp = angular.module('mamidApp', ['ngRoute']);

mamidApp.config(function ($routeProvider) {
   $routeProvider
       .when('/', {
           templateUrl: 'pages/home.html',
           controller: 'mainController',
           controllerAs: 'main'
       })
       .when('/slaves', {
           templateUrl: 'pages/slaves.html',
           controller: 'slavesController',
           controllerAs: 'slaves'
       })
});

mamidApp.controller('mainController', function() {
    this.message = 'Greetings from the controller';
});

mamidApp.controller('slavesController', function() {
    this.message = 'List of slaves'
});