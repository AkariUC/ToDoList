package todoModel

import (
	_ "github.com/go-sql-driver/mysql"

	"ToDoList/constant"
	"ToDoList/infrastructure/db"
)

type Todo struct {
	TodoUserId  int64
	TodoArticle string
	TodoLimit   string
	TodoTagId   int64
}

// InsertTodo データベースをレコードを登録する
func InsertTodo(record *Todo) error {
	stmt, err := db.Conn.Prepare("INSERT INTO todo (todo_user_id, todo_article, todo_limit, todo_tag_id, todo_complete, todo_existence) VALUES (?, ?, ?, ?, ?, ?); ")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(record.TodoUserId, record.TodoArticle, record.TodoLimit, record.TodoTagId, constant.Complete, constant.Existence)
	return err
}
