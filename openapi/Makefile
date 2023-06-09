dc = docker compose

bootstrap: ## 外部ツールをインストールする。
	# https://redocly.com/docs/cli/installation/
	npm i -g @redocly/cli@latest
	# https://github.com/deepmap/oapi-codegen#overview
	go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest

.PHONY: redoc
redoc:	## openapi のプレビューをブラウザで開く。
	$(dc) up -d
	@echo preview openapi: http://localhost:12345

.PHONY: down
down:
	$(dc) down -v

.PHONY: gen
gen:	# openapi.yaml から go のコードを生成する。
	redocly bundle ./openapi.yml --output ./gen/generated.yml
	oapi-codegen -config gen/go-model.config.yml -generate gin-server ./gen/generated.yml
