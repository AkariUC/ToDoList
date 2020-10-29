#

infrastructure層は、DBアクセスなどの技術的関心を記述します。  
この層はdomain層に依存しています。  
純粋なレイヤードアーキテクチャの場合、依存の向きがdomain → infrastructureですが、今回はDDDを取り込んだ設計になるので、依存の向きが逆転します。  
そのためinfrastructure層はdomain層のrepositoryで定義したインタフェースを実装します。

技術的  