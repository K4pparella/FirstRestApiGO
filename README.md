# FirstRestApiG
Hi, this is my first RESTful API made in golang. Thanks to [Ossan]([https://website-name.com](https://github.com/ossan-dev)https://github.com/ossan-dev) for helping me out learning go
***
# Features
|Function|Status|
|-----|--------|
|Get all todos|Functioning|
|Get a todo via ID|Functioning|
|Add a new todo|Functioning|
|Update a todo|Functioning|
|Complete a todo|Functioning|
|Delete a todo|Functioning|
***

# How to run
Firstly make sure you have GO installed on your computer, if not download it from [HERE!](https://go.dev/dl/).
Then download and exctract everything into a folder of your choice.
After that just open a terminal inside the folder and write
```
go run main.go
```
***
# Postman collection
If you want to use postman to check the api just use the collection below!
```json
{
	"info": {
		"_postman_id": "304eb948-2416-40f5-84f3-62aec6093442",
		"name": "FirstGOLANGEX",
		"schema": "https://schema.getpostman.com/json/collection/v2.1.0/collection.json",
		"_exporter_id": "30626451"
	},
	"item": [
		{
			"name": "Get a json of all todos",
			"request": {
				"method": "GET",
				"header": [],
				"url": {
					"raw": "localhost:2555/todos/",
					"host": [
						"localhost"
					],
					"port": "2555",
					"path": [
						"todos",
						""
					]
				}
			},
			"response": []
		},
		{
			"name": "Get a todo via his ID",
			"protocolProfileBehavior": {
				"disableBodyPruning": true
			},
			"request": {
				"method": "GET",
				"header": [],
				"body": {
					"mode": "raw",
					"raw": "{\r\n    \"id\": \"4\",\r\n    \"content\": \"test44324\",\r\n    \"finished\": false,\r\n    \"endYear\": \"2023\",\r\n    \"endMonth\": \"12\",\r\n    \"endDay\": \"5\"\r\n}",
					"options": {
						"raw": {
							"language": "json"
						}
					}
				},
				"url": {
					"raw": "localhost:2555/todos/3",
					"host": [
						"localhost"
					],
					"port": "2555",
					"path": [
						"todos",
						"3"
					]
				}
			},
			"response": []
		},
		{
			"name": "Add a todo (With all checks done)",
			"request": {
				"method": "POST",
				"header": [],
				"body": {
					"mode": "raw",
					"raw": "{\r\n    \"id\": \"4\",\r\n    \"content\": \"test4\",\r\n    \"finished\": false,\r\n    \"endYear\": \"2023\",\r\n    \"endMonth\": \"12\",\r\n    \"endDay\": \"5\"\r\n}",
					"options": {
						"raw": {
							"language": "json"
						}
					}
				},
				"url": {
					"raw": "localhost:2555/todos/",
					"host": [
						"localhost"
					],
					"port": "2555",
					"path": [
						"todos",
						""
					]
				}
			},
			"response": []
		},
		{
			"name": "Edit a TODO content",
			"request": {
				"method": "PATCH",
				"header": [],
				"url": {
					"raw": "localhost:2555/todos/3/E",
					"host": [
						"localhost"
					],
					"port": "2555",
					"path": [
						"todos",
						"3",
						"E"
					]
				}
			},
			"response": []
		},
		{
			"name": "CompleteTODO",
			"request": {
				"method": "PATCH",
				"header": [],
				"url": {
					"raw": "localhost:2555/todos/3",
					"host": [
						"localhost"
					],
					"port": "2555",
					"path": [
						"todos",
						"3"
					]
				}
			},
			"response": []
		},
		{
			"name": "DeleteTODO",
			"request": {
				"method": "DELETE",
				"header": [],
				"url": {
					"raw": "localhost:2555/todos/3",
					"host": [
						"localhost"
					],
					"port": "2555",
					"path": [
						"todos",
						"3"
					]
				}
			},
			"response": []
		}
	]
}
```
