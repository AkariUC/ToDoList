#
usecase層の責務はinterfaces層から情報を受け取り、domain層で定義してある関数を用いて任意のビジネスロジックを実行することです。



このように、usecase層はバリデーション、ユーザIDの生成などのビジネスロジックを記述したり、infrastructure層で実装したDBアクセスに関する処理をdomain層を介して間接的に呼んだりします。
ここに関してはイメージがしずらいと思います。
「実際に利用するのはinfrastructure層の具体的な内容が実装されている関数だから結局usecase層はinfrastructure層に依存するんじゃないの？」と思う人が多いと思います。
そこで先ほどから出てきているNewXxxXxxやここでもUserUsecaseでinterfaceを定義していることによって間接的に呼びだしているように実装できるんです。
アーキテクチャの話をするとよく聞くDI（依存性の注入）というものに利用します。
それらについては後で説明します。

バリデーション