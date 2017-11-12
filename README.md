## Go入門

### Tour of Go (Exercise)
+ ニュートン法
+ 画像の生成
+ Mapの値をカウントする
+ フィボナッチ数列
+ IPアドレスをドットで区切る(Stringers)
+ 平方根で負数渡したらerrorを返却する
+ ASCII文字 'A' の無限ストリームを出力する(Reader)
+ ROT13換字式暗号をアルファベットの文字に適用して読み出す

### 課題1. フィボナッチ
+ n番目のフィボナッチを返却する
+ TableDrivenTestsを使ったテスト
+ fibするコマンドラインツールの作成



## createClassFileTool
クラス名とパッケージを指定値に変更したコピー元ファイルを複製するツール
#### 使い方
1. コピー元ファイルをこのツールと同じディレクトリに配置する
2. コピー元ファイルの中身で置換したい文字列を「sample」に修正する
3. コマンドライン引数に下記の引数を渡して実行する
```
go run createCopyFile.go {コピー元ファイル} {作成するディレクトリパス} {作成するファイル名}

```
指定したディレクトリにファイルが作成される
