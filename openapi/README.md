## openapi

[gh-pages](https://kokoichi206.github.io/awesome-chat-app/) で確認できます。  
（[github actions](../.github/workflows/deploy-openapi.yml) で更新されます。）

## Open API の preview 確認

以下のように、redoc を docker compose で使うことで、ブラウザ上で preview することが可能です。

``` sh
# http://localhost:12345 にアクセスする。
docker compose up -d

# make が使える場合。
make
```

## Stub server の起動

ws 以外のエンドポイントについては, [prism](https://github.com/stoplightio/prism) を使うことで stub server を立てることが可能です。

``` sh
# デフォルトでは 127.0.0.1:4010 で起動する。
prism mock openapi.yml

# port を指定する。
prism mock openapi.yml -p 2828
```

## Links

- [OpenAPI Specification v3.1.0](https://spec.openapis.org/oas/v3.1.0)
