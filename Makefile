IMAGE_NAME="ecareer_bot:latest"

build:
	docker build --rm -t ${IMAGE_NAME} .

run-dev:
	docker run --rm --name "ecareer_bot" -v "/usr/local/_data:/data" --env-file .env ${IMAGE_NAME} sh -c 'go run cmd/main.go'