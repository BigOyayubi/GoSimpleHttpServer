# gosimplehttpd

指定フォルダ以下のファイルをGetで取得するだけの開発用サーバー
exeを実行するだけで実行ディレクトリ以下のファイルがGetできます。

# QuickStart 

```
$ pwd
/home/hoge

$ ls ./test/
file1.txt file2.txt

$ cat ./test/file1.txt
hoge

$ $GOPATH/bin/gosimplehttpd
root dir : /home/hoge
port     : 8080

$ curl http://localhost:8080/test/file1.txt
hoge

$ $GOPATH/bin/gosimplehttpd -d test -p 8081
root dir : /home/hoge/test
port     : 8081
```
