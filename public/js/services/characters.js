/*global define*/
'use strict';

/**
 * Services that persists and retrieves TODOs from localStorage.
 */
define(['app'], function (app) {

	app.service('CharacterService', function ($http, $q) {
		var find = function(offset, limit){
			var def = $q.defer();
			var api_key = '344dd9e52bab584a2b99ce0c60a3e99b';
			var base_url = 'http://gateway.marvel.com/v1/';

			$http.get(base_url+'public/characters?limit='+limit+'&offset='+offset+'&apikey='+api_key).success(def.resolve).error(def.reject);
			return def.promise;
		};

		return {
			find: find
		};
	});
});