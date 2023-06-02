dc = docker compose

.PHONY: psql
psql:
	$(dc) exec postgres psql -U postgres awesome-chat-app

db-up:
	$(dc) up -d

db-down:
	$(dc) down -v
