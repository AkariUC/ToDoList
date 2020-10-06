CREATE DATABASE IF NOT EXISTS todo_db;

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
    FOREIGN KEY('todo_user_id')  REFERENCES user('id') ON UPDATE CASCADE ON DELETE CASCADE,
    FOREIGN KEY('todo_tag_id') REFERENCES tag('id')  ON UPDATE CASCADE ON DELETE CASCADE,
)
ENGINE = InnoDB
COMMENT = 'ToDo List';


CREATE TABLE IF NOT EXISTS `todo_db`.`tag`
(
    `id`   INT          NOT NULL COMMENT 'id',
    `tag_data` VARCHAR(255) NOT NULL COMMENT '内容',
    PRIMARY KEY (`id`)
)
ENGINE = InnoDB
COMMENT = 'タグ';

INSERT INTO todo_db.tag(id, tag_data) VALUES( 1, '家事');
INSERT INTO todo_db.tag(id, tag_data) VALUES( 2, '勉強');
INSERT INTO todo_db.tag(id, tag_data) VALUES( 3, '就活');
INSERT INTO todo_db.tag(id, tag_data) VALUES( 4, '課題');
INSERT INTO todo_db.tag(id, tag_data) VALUES( 5, 'その他');