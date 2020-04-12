build:
	go build -o bin/hello 001-hello/you-go.go
	go build -o bin/web 002-functional/functional.go
	go build -o bin/web 003-web/web.go
	GOOS=linux GOARCH=amd64 go build -o 004-once-more-unto-the-breach-dear-friends/bin/main ./004-once-more-unto-the-breach-dear-friends/main/main.go
test:
	cd 001-hello && go test
	cd 002-functional && go test
	cd 003-web && go test
	GOOS=linux GOARCH=amd64 cd ./004-once-more-unto-the-breach-dear-friends/main && go test
run-001:
	go run 001-hello/you-go.go

run-002:
	go run 002-functional/functional.go

run-003:
	go run 003-web/web.go

build-004:
	GOOS=linux GOARCH=amd64 go build -o 004-once-more-unto-the-breach-dear-friends/bin/main ./004-once-more-unto-the-breach-dear-friends/main

