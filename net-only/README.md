## Net only
標準ライブラリのnetを使ったtcpレベルの接続テスト

サーバー側で
``` bash
go run ./main.go
```
を実施後、クライアント側も接続する
``` bash
go run ./client/main.go
```