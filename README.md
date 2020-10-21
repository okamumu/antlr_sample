必要なツールをインストール

```sh
sudo apt install antlr4
go get github.com/antlr/antlr4/runtime/Go/antlr
```

カレントディレクトリを移動

```sh
cd parser
```

以下の実行して ANTLR でファイルを生成

```sh
curl -O https://raw.githubusercontent.com/antlr/grammars-v4/master/c/C.g4
antlr4 -Dlanguage=Go -listener C.g4
```

自分で作るファイル（今回は適当に作っています）
- `mylistener.go`: `c_base_listener.go` を埋め込んでメソッドオーバーライドをする

テストファイル `mylistener_test.go` にパースの実行例

```sh
go test
```

で実行

