run:
	go run ./cmd/url-shortener

clean-restart-db:
	docker compose down && docker volume rm backend_postgres_data && docker compose up -d

restart-db:
	docker compose down && docker compose up -d
