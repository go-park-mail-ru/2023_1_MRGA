.PHONY: runserver test cov
host=127.0.0.1
port=8080

port_storage=8081


runserver:
	docker-compose up
	go run cmd/meetme/main.go -h $(host) -p $(port)
	go run services/files_storage/cmd/main.go -p $(port_storage)

test:
	go test ./...

cov:
	go test -coverpkg=./... -coverprofile=cover ./... && cat cover | grep -v "mock" | grep -v  "easyjson" | grep -v "proto" > cover.out && go tool cover -func=cover.out

run:
	 go run ./cmd/meetme/main.go -h 0.0.0.0 -p 8080 &
	go run ./services/files_storage/cmd/main.go -p 8081 &
	go run ./services/auth/cmd/auth/main.go &
	go run ./services/complaints/cmd/migrate/main.go &
	go run ./services/complaints/cmd/app/main.go &
	go run ./services/chat/cmd/main.go &

