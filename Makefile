db-up:
	@docker compose up -d

db-down:
	@docker compose down

clean:
	@docker compose down -v

up:
	go run main.go