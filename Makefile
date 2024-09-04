run:
	@echo :: Starting the app
	@go run cmd/main/main.go

build :
	@echo :: Building the app
	@go build -o bin/app.exe cmd/main/main.go