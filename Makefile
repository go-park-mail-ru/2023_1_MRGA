.PHONY: runserver test cov
host=localhost
runserver:
	go run cmd/meetme/main.go $(host) 8080

test:
	go test ./...

cov:
	go test -coverpkg=./... -coverprofile=cover ./... && cat cover | grep -v "mock" | grep -v  "easyjson" | grep -v "proto" > cover.out && go tool cover -func=cover.out