.PHONY: build help bootstrap godoc
.DEFAULT_GOAL := help

EXTERNAL_TOOLS := \
	github.com/golangci/golangci-lint/cmd/golangci-lint@v1.51.1 \
	github.com/golang/mock/mockgen@v1.6.0 \
	golang.org/x/pkgsite/cmd/pkgsite@latest # latest は go 1.19 以上が必要: https://github.com/golang/pkgsite#requirements

help:	## https://postd.cc/auto-documented-makefile/
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

bootstrap: ## 外部ツールをインストールする。
	for t in $(EXTERNAL_TOOLS); do \
		echo "Installing $$t ..." ; \
		go install $$t ; \
	done

mock:	## interface に沿った mock を生成する。
	mockgen -source=usecase/usecase.go -destination=handler/mock_usecase_test.go -package=handler_test
	mockgen -source=repository/repository.go -destination=usecase/mock_repository_test.go -package=usecase_test

DC = docker compose
psql:	## docker compose で起動した postgresql の db に接続する。
	$(DC) exec postgres psql -U root postgresql 

godoc:	## godoc をローカルで表示する。http://localhost:8080/{module_name}
	pkgsite

.PHONY: lint lint-fix serve

lint:	## golangci を使って lint を走らせる。
	golangci-lint run -v

lint-fix:	## lint 実行時, gofumpt のエラーが出たらやると良い。
	golangci-lint run --fix

serve:	## サーバーを起動する。
	go run app/*

build-local:	## バイナリをビルドする（race オプションがついているため、ローカル実行専用とする）。
	go build -race -o app-local app/*

# カバレッジが低い場合は build-loacl でも動かしてみて競合の確認をしたい。
test:	## 全テストを実行する。
	go test -race -cover -shuffle=on ./... -v

jaeger:	## docker で jaeger を起動する。
	docker run -d --name jaeger \
		-e COLLECTOR_ZIPKIN_HTTP_PORT=9411 \
		-p 5775:5775/udp \
		-p 6831:6831/udp \
		-p 6832:6832/udp \
		-p 5778:5778 \
		-p 16686:16686 \
		-p 14268:14268 \
		-p 9411:9411 \
		jaegertracing/all-in-one:latest
	@echo see: http://localhost:16686/search
