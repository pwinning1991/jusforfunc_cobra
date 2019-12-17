.PHONY: build

build:
	@go build
install:
	@go install
run_server:
	@docker run -d -p 8888:8888 --name todo_server pwinnington/todo_server:v0.1.0
clean_server:
	@docker stop todo_server && docker rm todo_server
