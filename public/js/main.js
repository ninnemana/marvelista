require.config({
	paths: {
		angular: '/js/vendor/angular/angular.min',
		jquery: '/js/vendor/jquery/dist/jquery.min',
		ngResource: '/js/vendor/angular-resource/angular-resource.min'
	},
	shim: {
		angular: {
			exports: 'angular'
		},
		jquery:{
			exports: '$'
		},
		ngResource: ['angular']
	}
});

require(['angular', 'app', 'jquery', 'ngResource', 'controllers/characters'], function (angular) {
	'use strict';
	angular.bootstrap(document, ['marvelista']);
});