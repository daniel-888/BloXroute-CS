.PHONY: gen-protobuf run-demo-docker-compose stop-demo-docker-compose run-server run-client run-rabbit-mq run-tests

gen-protobuf:
	protoc --proto_path=models --go_out=models --go_opt=paths=source_relative models/command.proto

run-demo-docker-compose:
	docker compose --profile demo up

stop-demo-docker-compose:
	docker compose down --remove-orphans

run-server:
	cd server && go run main.go server

run-client-add-item:
	cd client && go run main.go AddItem

run-client-remove-item:
	cd client && go run main.go RemoveItem

run-client-get-item:
	cd client && go run main.go GetItem

run-client-get-all-items:
	cd client && go run main.go GetAllItems

run-rabbit-mq:
	docker compose run --service-ports rabbit

run-tests:
	go test ./server/... && go test ./client/...