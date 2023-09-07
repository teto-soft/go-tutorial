# 書き方

はじめに、 `package main` と書く

Go のプログラムは何らかのパッケージに依存している必要がある。  
また Go のプログラムのうちひとつは main パッケージに依存する必要がある。  
Go は main パッケージ内の main 関数からプログラムがスタートするようになっていて、プログラムがスタートする場所をエントリーポイントという。

次に fmt (format) パッケージをインポートする。

fmt パッケージの Println でログを出力する。

# ビルドと実行

` go build main.go` コマンドでビルドする。
`./main`コマンドでビルドされたファイルを実行する。

`go run main.go`コマンドでビルドと実行を行う。

```
package main
import ("fmt")

func main(){
    fmt.Println("Good Morning")
}
```
