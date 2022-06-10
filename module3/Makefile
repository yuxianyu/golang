build:
	mkdir bin/amd64
	GOOS=linux GOARCH=amd64 go build -o bin/adm64/httpserver .

release: build
	docker build -t 896914858/httpserver:v1.0

push:
	docker push 896914858/httpserver:v1.0
