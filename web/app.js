'use strict';

var urlapi = "http://127.0.0.1:3014/";
//var urlapi = "http://51.255.193.106:3014/";

// Declare app level module which depends on views, and components
angular.module('webApp', [
    'ngRoute',
    'ngMessages',
    'angularBootstrapMaterial',
    'angular-svg-round-progressbar',
    'app.navbar',
    'app.main',
    'app.block',
    'app.tx',
    'app.address',
    'app.network',
    'app.addressNetwork',
    'app.sankey',
    'app.dateAnalysis',
    'app.blocks',
    'app.txs',
    'app.addresses',
    'app.beta'
]).
config(['$locationProvider', '$routeProvider', function($locationProvider, $routeProvider) {
        $locationProvider.hashPrefix('!');

        $routeProvider.otherwise({
            redirectTo: '/main'
        });
    }]);
