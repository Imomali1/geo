build:
	env GOOS=linux GOARCH=amd64 go build -o bin/linux && \
	env GOOS=windows GOARCH=amd64 go build -o bin/windows