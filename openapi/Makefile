dc = docker compose

.PHONY: redoc
redoc:	## openapi のプレビューをブラウザで開く。
	$(dc) up -d
	@echo preview openapi: http://localhost:12345

.PHONY: down
down:
	$(dc) down -v
