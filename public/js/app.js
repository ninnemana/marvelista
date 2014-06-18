define(['angular'], function (angular) {
	'use strict';
	return angular.module('marvelista', [], function($interpolateProvider){
		$interpolateProvider.startSymbol('[[');
		$interpolateProvider.endSymbol(']]');
	});
});