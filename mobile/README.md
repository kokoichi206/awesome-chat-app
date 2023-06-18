## mobile クライアント

Compose Multiplatform で作成されてます。

## ローカルと通信確認時

1. [shared/build.gradle.kts](./shared/build.gradle.kts) 内の `BASE_URL` を変更する。
2. [network_security_config.xml](./androidApp/src/androidMain/res/xml/network_security_config.xml) で、ローカル IP への HTTP 通信を許可する。 
