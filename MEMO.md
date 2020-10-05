# 後から振り替えれるためのメモ

## 開発
* ファイル構成
https://qiita.com/hmarf/items/7f4d39c48775c205b99b


├── application    // Request、Responseを行うだけ  
│   └── server  
│       ├── handler  
│       ├── response  
│       └── server.go  
├── domain         // データを加工する。必要とあらば infra を通じてDB情報を使う  
│   ├── model      // structの塊  
│   ├── repository // Interfaceの塊（あとで説明します）  
│   └── service    // 処理をごちゃごちゃ書くこところ  
├── infrastructure // DBと通信する  
└── cmd  
    └── main.go  


## その他

### go.modについて

[Go Modules](https://qiita.com/propella/items/e49bccc88f3cc2407745)  
[Go言語の依存モジュール管理ツール Modules の使い方](https://qiita.com/uchiko/items/64fb3020dd64cf211d4e)  

### gofmt
[【Go】 go fmt でコード整形](https://qiita.com/taji-taji/items/6d286bf4483a4c6ceed6)  

### goimports
[goimportsを使ってIntelliJ IDEA でGoのコードのimportを手軽に整理する](https://qiita.com/dmnlk/items/c423d5853cc129ab77a4)