# Golang-Fastcgi-Try

Go言語でFastCGI x Apacheをした

## 動かし方

``` sh
make docker-build
docker-compose up
```

以下のように返ってくれば正しく動いています
``` sh
$ curl localhost:8888
{"message":"pong"}
#
