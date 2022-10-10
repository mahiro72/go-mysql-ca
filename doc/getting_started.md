# Getting Started

docker 環境を立ち上げる

```
make up
```

<br>

APIが動いているか確認する

http://localhost:8080/task/all


<br>

### ⚠️よくみるエラーについて

### ```make: command not found```

windowsの方はmake をインストールする必要があります

[こちら](https://gnuwin32.sourceforge.net/packages/make.htm)からインストールしてください

### ```listen tcp 0.0.0.0:3306: bind: Only one usage of each socket address (protocol/network address/port) is normally permitted.```

mysqld.exeがすでに動いているので、PIDを特定しタスクを終了します

詳しくは以下の記事をご覧ください

参考記事 : [dockerコンテナ起動時にOnly one usage of each socket addressというエラー](https://qiita.com/tbzgkzejyk1ef5s91wuf/items/cc2be6263523149d7cbe)
