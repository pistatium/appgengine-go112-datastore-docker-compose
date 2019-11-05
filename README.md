# appgengine-go112-datastore-docker-compose
AppEngine 1.12 + Datastore (+ gin ) の開発環境を Docker-compose でシュッと立ち上げるサンプル。

## ローカルでの動かし方 

```bash
docker-compose up -d
```

これで、 Goのサーバーと、Datastoreのエミュレータが起動します。

http://localhost:8000 を開けば Go のアプリケーションにアクセスできます。

fresh を導入しているので、main.go を変更すればホットリロードされます。

### 試し方
Entry の投稿と閲覧を仮で実装しています。

httpie を使って

```bash
# Entry を POST
http http://localhost:8000/entries title="タイトル" body="本文てすと"

# Entry 一覧を取得
http http://localhost:8000/entries
```

みたいな感じで試せます。

## デプロイ方法

```bash
gcloud app deploy
```
