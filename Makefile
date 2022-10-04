env:
	cp ./.env.example ./.env
up:
	docker-compose up -d
down:
	docker-compose down
run:
	go run main.go httpserver
