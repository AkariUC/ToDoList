# 後から振り替えれるためのメモ

## 開発

### レイアードアーキテクチャ
Interfaces層  
↓  
usecase層  
↓  
domain層  
↓  
infrastructure層  

### DDD
Interfaces層  
↓  
usecase層  
↓  
domain層  
↑  
infrastructure層

### ディレクトリ構成
レイヤードアーキテクチャの層は本来は一番上から、
Presentation → Application → Domain → Infrastructure
という名前が付いていますが責務のイメージが浮かびずらいと思うので今回は
Interfaces → Usecase → Domain → Infrastructure
という名前で進めさせていただきます。


## その他

### go.modについて

[Go Modules](https://qiita.com/propella/items/e49bccc88f3cc2407745)  
[Go言語の依存モジュール管理ツール Modules の使い方](https://qiita.com/uchiko/items/64fb3020dd64cf211d4e)  

### gofmt
[【Go】 go fmt でコード整形](https://qiita.com/taji-taji/items/6d286bf4483a4c6ceed6)  

### goimports
[goimportsを使ってIntelliJ IDEA でGoのコードのimportを手軽に整理する](https://qiita.com/dmnlk/items/c423d5853cc129ab77a4)


## 個人的なメモ(もり)
[MySQLのrootパスワードを忘れた時のリセット方法](https://qiita.com/miriwo/items/1880e9d2ebcfd3c0e60d)
