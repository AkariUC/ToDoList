# ToDoList

## 構想
### Todoリスト

クラアンのすること
* サインインログイン
* ToDoの登録(内容、期限日、タグ)
* ToDoの一覧の確認(終わっていないもの、終わっているもの)
* ToDoの完了
* ToDoの削除

### MySQLWorkbenchの設定
MySQLへの接続設定をします。
1. MySQL Connections の + を選択
2. 以下のように接続設定を行う
    ```
    Connection Name: 任意
    Connection Method: Standard (TCP/IP)
    Hostname: 127.0.0.1 (localhost)
    Port: 3306
    Username: moriakari
    Password: pass123
    Default Schema: todo_db
    ```

### API用のデータベースの接続情報を設定する
環境変数にデータベースの接続情報を設定します。<br>
ターミナルのセッション毎に設定したり、.bash_profileで設定を行います。

Macの場合
```
$ export MYSQL_USER=moriakari \
    MYSQL_PASSWORD=pass123 \
    MYSQL_HOST=127.0.0.1 \
    MYSQL_PORT=3306 \
    MYSQL_DATABASE=todo_db
```

## MYSQL
[ユーザを作成する理由](https://techacademy.jp/magazine/5110)

#### 1 rootでMySQLコマンドラインツールの起動  
$`mysql -u root -p`

#### 2 rootで'todo_db'作成
mysql>`CREATE DATABASE IF NOT EXISTS todo_db;`

#### 3 rootでユーザー作成
mysql>`create user 'moriakari'@'localhost' identified by 'pass123';` 

#### 3 rootでユーザーに権限付与
mysql>```GRANT SELECT, INSERT, UPDATE, DELETE, CREATE, DROP, RELOAD, PROCESS, REFERENCES, INDEX, ALTER, SHOW DATABASES, CREATE TEMPORARY TABLES, LOCK TABLES, EXECUTE, REPLICATION SLAVE, REPLICATION CLIENT, CREATE VIEW, SHOW VIEW, CREATE ROUTINE, ALTER ROUTINE, CREATE USER, EVENT, TRIGGER ON *.* TO 'moriakari'@'localhost' WITH GRANT OPTION;``` 

#### 4 出る
mysql>`exit`

#### 5 作成したユーザでmysqlに接続する
$`mysql -u moriakari -p`  
PASS=`pass123`

#### 6 DB、テーブル作成
```
CREATE DATABASE IF NOT EXISTS todo_db;
```
```
USE `todo_db`;
```

```
CREATE TABLE IF NOT EXISTS `todo_db`.`user`
(
    `id`           INT AUTO_INCREMENT NOT NULL COMMENT 'id',
    `name`         VARCHAR(255)       NOT NULL COMMENT 'ユーザ名',
    `password`     VARCHAR(255)       NOT NULL COMMENT 'パスワード',
    `auth_token`   VARCHAR(255)       NOT NULL COMMENT '認証トークン',
    `existence`    INT                NOT NULL COMMENT 'ユーザの存在の有無',
    PRIMARY KEY (`id`)
)
ENGINE = InnoDB
COMMENT = 'ユーザ';
```

```
CREATE TABLE IF NOT EXISTS `todo_db`.`tag`
(
    `id`   INT          NOT NULL COMMENT 'id',
    `tag_data` VARCHAR(255) NOT NULL COMMENT '内容',
    PRIMARY KEY (`id`)
)
ENGINE = InnoDB
COMMENT = 'タグ';
```

```
INSERT INTO todo_db.tag(id, tag_data) VALUES( 1, '家事');
INSERT INTO todo_db.tag(id, tag_data) VALUES( 2, '勉強');
INSERT INTO todo_db.tag(id, tag_data) VALUES( 3, '就活');
INSERT INTO todo_db.tag(id, tag_data) VALUES( 4, '課題');
INSERT INTO todo_db.tag(id, tag_data) VALUES( 5, 'その他');
```

```
CREATE TABLE IF NOT EXISTS `todo_db`.`todo`
(
    `id`             INT AUTO_INCREMENT NOT NULL COMMENT 'id',
    `todo_user_id`   INT                NOT NULL COMMENT 'todoリストを作成したユーザのid',
    `todo_article`   VARCHAR(255)       NOT NULL COMMENT '内容',
    `todo_limit`     DATETIME                    COMMENT 'TODOの期限',
    `todo_tag_id`    INT                NOT NULL COMMENT 'タグ',
    `todo_complete`  INT                NOT NULL COMMENT '終了したかどうか',
    `todo_existence` INT                NOT NULL COMMENT '削除したかどうか',
    PRIMARY KEY (`id`),
    FOREIGN KEY(`todo_user_id`) REFERENCES user(`id`) ON UPDATE CASCADE ON DELETE CASCADE,
    FOREIGN KEY(`todo_tag_id`) REFERENCES tag(`id`) ON UPDATE CASCADE ON DELETE CASCADE
)
ENGINE = InnoDB
COMMENT = 'ToDoList';
```