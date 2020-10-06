CREATE DATABASE IF NOT EXISTS todo_db;

CREATE TABLE IF NOT EXISTS `todo_db`.`user`
(
    `id`         INT AUTO_INCREMENT NOT NULL COMMENT 'id',
    `name`       VARCHAR(64)        NOT NULL COMMENT 'ユーザ名',
    `password`   VARCHAR(64)        NOT NULL COMMENT 'パスワード',
    `auth_token` VARCHAR(128)       NOT NULL COMMENT '認証トークン',
    PRIMARY KEY (`id`)
)
ENGINE = InnoDB
COMMENT = 'ユーザ';


CREATE TABLE IF NOT EXISTS `todo_db`.`todo`
(
    `id`            INT AUTO_INCREMENT NOT NULL COMMENT 'id',
    `user_id`       INT                NOT NULL COMMENT 'todoリストを作成したユーザのid',
    `todo_sentence` VARCHAR(256)       NOT NULL COMMENT '内容',
    `todo_tag`      INT                NOT NULL COMMENT 'タグ',
    `todo_deadline` DATETIME                    COMMENT 'TODOの期限',
    `todo_complete` INT                NOT NULL COMMENT '終了したかどうか',
    `todo_delete`   INT                NOT NULL COMMENT '削除したかどうか',
  PRIMARY KEY (`id`),
  FOREIGN KEY('user_id')  REFERENCES user('id') ON UPDATE CASCADE ON DELETE CASCADE,
  FOREIGN KEY('todo_tag') REFERENCES tag('id')  ON UPDATE CASCADE ON DELETE CASCADE,
)
ENGINE = InnoDB
COMMENT = 'ToDo List';


CREATE TABLE IF NOT EXISTS `todo_db`.`tag`
(
    `id`  INT          NOT NULL COMMENT 'id',
    `tag` VARCHAR(256) NOT NULL COMMENT '内容',
    PRIMARY KEY (`id`)
)
ENGINE = InnoDB
COMMENT = 'タグ';

INSERT INTO todo_db.tag(id, tag) VALUES( 1, '家事');
INSERT INTO todo_db.tag(id, tag) VALUES( 2, '勉強');
INSERT INTO todo_db.tag(id, tag) VALUES( 3, '就活');
INSERT INTO todo_db.tag(id, tag) VALUES( 4, '課題');
INSERT INTO todo_db.tag(id, tag) VALUES( 5, 'その他');