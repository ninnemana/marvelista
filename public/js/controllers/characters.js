define(['app', 'services/characters'], function (app) {


	return app.controller('CharactersController', ['CharacterService','$scope', '$location',
		function CharactersController(characterService, $scope, $location) {
			$scope.next_offset = 0;
			$scope.characters = [];

			$scope.loadCharacters = function(e){
				characterService.find($scope.next_offset,100).then(function(resp){

					$scope.characters = $scope.characters.concat(resp.data.results);
					$scope.next_offset = $scope.next_offset + 100;
					$scope.limit = 100;
				});
			}
		}
	]);
});