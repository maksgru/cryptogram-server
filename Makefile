build:
	go build -v ./cmd/app

test:
	go test -v -race -timeout 30s ./...

swagger:
	swag init -g ./cmd/app/main.go -o ./docs
