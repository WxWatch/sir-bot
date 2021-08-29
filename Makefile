build:
	GOOS=linux GOARCH=amd64 go build -o ./dist/sir-bot ./src/main.go