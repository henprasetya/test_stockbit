# test_stockbit
Access with postman :
{
	"info": {
		"_postman_id": "2ec40117-2758-4668-b9c5-31d4cae9a66f",
		"name": "omdbapi",
		"schema": "https://schema.getpostman.com/json/collection/v2.1.0/collection.json"
	},
	"item": [
		{
			"name": "search_movie",
			"request": {
				"method": "GET",
				"header": [],
				"url": {
					"raw": "localhost:8080/api/search_movie?search=Batman&page=2",
					"host": [
						"localhost"
					],
					"port": "8080",
					"path": [
						"api",
						"search_movie"
					],
					"query": [
						{
							"key": "search",
							"value": "Batman"
						},
						{
							"key": "page",
							"value": "2"
						}
					]
				}
			},
			"response": []
		},
		{
			"name": "movie_detail",
			"request": {
				"method": "GET",
				"header": [],
				"url": {
					"raw": "localhost:8080/api/movie_detail?omdbID=tt4853102",
					"host": [
						"localhost"
					],
					"port": "8080",
					"path": [
						"api",
						"movie_detail"
					],
					"query": [
						{
							"key": "omdbID",
							"value": "tt4853102"
						}
					]
				}
			},
			"response": []
		}
	]
}
