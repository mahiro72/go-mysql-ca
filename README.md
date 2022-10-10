# TODO LIST API

## About

このリポジトリはGoで書かれたTodo List APIです

<br>

実装について、技術はGoやMySQL、アーキテクチャはCleanArchitectureを採用しています

アプリケーションの起動については、[こちら](./doc/getting_started.md)を参考にしてください

コマンドの一覧は`make`か`make help`で確認できます

<br>

## ディレクトリ構成

このリポジトリのディレクトリ構成です

```
|--.github
|  |--workflows   # workflowsは、github actionsがまとめてあります
|--db             # dbは、mysqlのスキーマやコンテナのマウントしたデータがあります
|--doc            # docは、このリポジトリのドキュメントが管理してあります
|--docker         # dockerは、docker-composeファイルがまとめてあります
|--Makefile       # Makefileには環境構築用のコマンドがまとめてあります
|--src            # srcには、アプリケーションのコードがあります
```

<br>


### 利用技術やツールについて

このアプリケーションで使われている技術やツールについての詳細な説明があります

詳しくは、[こちら](./doc/detail.md)からご覧ください

<br>

### アーキテクチャについて

このアプリケーションで採用したアーキテクチャについて詳細が書いてあります

詳しくは、[こちら](./doc/architecture.md)からご覧ください

<br>

### APIについて

APIのドキュメントについては、[こちら](./doc/api.md)からご覧ください

<br>
