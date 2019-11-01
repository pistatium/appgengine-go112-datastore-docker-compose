# appgengine-go112-datastore-docker-compose
AppEngine 1.12 + Datastore (+ gin ) の開発環境を Docker-compose でシュッと立ち上げるサンプル。

## ローカルでの動かし方 

```bash
docker-compose up -d
```

これで、 Goのサーバーと、Datastoreのエミュレータが起動します。

http://localhost:8000 を開けば Go のアプリケーションにアクセスできます。


## デプロイ方法

```bash
gcloud app deploy
```
