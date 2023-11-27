up:
	@docker compose up -d

down:
	@docker compose down
	docker system prune -af

migrate-up:
	migrate -path migrations -database "postgresql://root:secret@localhost:4000/mycloud?sslmode=disable" -verbose up

migrate-down:
	migrate -path migrations -database "postgresql://root:secret@localhost:4000/mycloud?sslmode=disable" -verbose down

