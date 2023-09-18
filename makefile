build-api:
	go build -o ./app/services/api/ ./app/services/api 

run-api: build-api
	app/services/api/./api
