# omochi

## migrate

`golang-migrate`のインストールからマイグレートの実行までを簡単に説明する

参考: <https://github.com/golang-migrate/migrate>

### golang-migrateのインストール

```sh
brew install golang-migrate
```

### マイグレートファイルの作成

upファイルにマイグレート時に実行するsqlを記述し、downファイルにロールバック時に実行するsqlを記述する

```sh
# 以下のファイルが作成される
# {timestamp}_{filename}.up.sql
# {timestamp}_{filename}.down.sql

migrate create -ext sql -dir db/migrations ファイル名
```

### マイグレートの実行

```sh
# マイグレートする場合
migrate -source file://db/migrations/ -database 'mysql://root:password@tcp(127.0.0.1:3306)/omochi' up

# ロールバックする場合
migrate -source file://db/migrations/ -database 'mysql://root:password@tcp(127.0.0.1:3306)/omochi' down
```

up, downの後に数字をつけて実行すると最新のファイルから数字分のファイルまでのマイグレート、ロールバックを実行する

```sh
# 一つだけロールバックしたい場合
migrate -source file://db/migrations/ -database 'mysql://root:password@tcp(127.0.0.1:3306)/omochi' down 1
```
