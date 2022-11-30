.PHONY: gen-protobuf run-demo-docker-compose stop-demo-docker-compose run-server run-client run-rabbit-mq run-tests

gen-protobuf:
	protoc --proto_path=models --go_out=models --go_opt=paths=source_relative models/command.proto

run-rabbit-mq:
	docker compose run --service-ports rabbit