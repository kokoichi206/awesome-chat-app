# server

## Environments

| key | description | default |
| :---: | :---: | :---: |
| SERVER_HOST |  | localhost |
| SERVER_PORT |  | 8080 |
| AGENT_HOST |  |  |
| AGENT_PORT |  |  |
| DB_DRIVER |  | postgres |
| DB_HOST |  | localhost |
| DB_PORT |  | 5432 |
| DB_USER |  | root |
| DB_PASSWORD |  | root |
| DB_NAME |  | sakamichi |
| DB_SSL_MODE |  | disable |
| FIREBASE_ADMIN_CONFIG_PATH | Firebase admin sdk の設定ファイルへのパス |  |

### ローカル実行

``` go
func init() {
	os.Setenv("AGENT_HOST", "localhost")
	os.Setenv("AGENT_PORT", "5775")

	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_PORT", "5432")
	os.Setenv("DB_USER", "postgres")
	os.Setenv("DB_PASSWORD", "postgres")
	os.Setenv("DB_NAME", "awesome-chat-app")

	os.Setenv("FIREBASE_ADMIN_CONFIG_PATH", "config/awesome-chat-app-c9e81-firebase-adminsdk-gdtpk-randomdesu.json")
}
```

## License

This repository is under [MIT License](./LICENSE).
