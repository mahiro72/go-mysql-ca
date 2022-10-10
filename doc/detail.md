# 各技術の詳細について

## Gin
今回はルーティング処理などをGinを使って行いました

理由は以下があげられます
- ルーティング処理が簡単
- middlewareの導入が簡単
- かなり有名なライブラリなので、知見がネット上に多くある

懸念点としては gin.Contextに依存しているため、
gin以外に置き換える時に変更する量がかなり多くなってしまうことです

他に検討したライブラリである[go-chi](https://github.com/go-chi/chi)はその点移行しやすいです

## sqlx
今回はDBとの通信処理をsqlxを使って行いました

理由は以下があげられます
- 素のSQLがかける
- 最低限の欲しい機能がそろっている

database/sqlだとScanで記述する処理が長くなりますが、
sqlxだとそのあたりをよりシンプルにかけます

他のライブラリにGoで有名なORMである[gorm](https://github.com/go-gorm/gorm)などありますが
こちらは書きやすい分、かなりライブラリに依存してしまうため今回は利用しませんでした

## MySQL
今回はDBにMySQLを採用しました


## Docker, Docker Compose
インフラまわりはDockerで構築します

## Makefile
makeでコマンドをいろいろまとめています

makeを使うことで普段使っているコマンドをまとめて実行できます

とくにdocker環境においてコンテナの削除やマウントしたデータを削除、イメージまで削除するとなると
かなり手間です

そこでmakeを利用することで一つのコマンドでまとめて実行できます

デメリットとしては、makefileの中身に詳しい人がいなくなったとき
内部でどのようなコマンドが走っているか誰もわからなくなり、ブラックボックス化することです


今回のプロジェクトのmakeについて、
どのようなコマンドがあるかは、```make```または```make help```で確認してください

macの方はもともとmakeが入ってますが、windowsの方は[こちら](https://gnuwin32.sourceforge.net/packages/make.htm)からinstallする必要があります


## GitHub Actions
簡易的ですが、CI(継続的インテグレーション)を導入しました

このプロジェクトではbuildができるかと、testを走らせてFAILしないかを確認しています