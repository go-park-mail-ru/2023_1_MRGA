.PHONY: runserver test cov
host=127.0.0.1
port=8080

runserver:
	go run cmd/meetme/main.go -h $(host) -p $(port)

test:
	go test ./...

cov:
	go test -coverpkg=./... -coverprofile=cover ./... && cat cover | grep -v "mock" | grep -v  "easyjson" | grep -v "proto" > cover.out && go tool cover -func=cover.out