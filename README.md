你好！
很冒昧用这样的方式来和你沟通，如有打扰请忽略我的提交哈。我是光年实验室（gnlab.com）的HR，在招Golang开发工程师，我们是一个技术型团队，技术氛围非常好。全职和兼职都可以，不过最好是全职，工作地点杭州。
我们公司是做流量增长的，Golang负责开发SAAS平台的应用，我们做的很多应用是全新的，工作非常有挑战也很有意思，是国内很多大厂的顾问。
如果有兴趣的话加我微信：13515810775  ，也可以访问 https://gnlab.com/，联系客服转发给HR。
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
