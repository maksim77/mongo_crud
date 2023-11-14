start:
	docker compose -p mongo -f deployments/docker-compose.yml up -d

stop:
	docker compose -p mongo down

clean:
	rm -f ./mongo_crud

build: clean
	go build .

run: build
	./mongo_crud

mongosh:
	mongosh "mongodb://root:example@127.0.0.1:27017/?authSource=admin"

aggregate:
	mongosh "mongodb://root:example@127.0.0.1:27017/?authSource=admin" --file aggregate.js

import:
	mongoimport --db=strava --collection=workout "mongodb://root:example@127.0.0.1:27017/?authSource=admin" backup.json