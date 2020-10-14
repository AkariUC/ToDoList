package todoModel

import (
	"log"

	_ "github.com/go-sql-driver/mysql"

	"ToDoList/infrastructure/db"
)

type Todo struct {
	TodoUserId    int64
	TodoArticle   string
	TodoLimit     string
	TodoTagId     int64
	TodoComplete  int
	TodoExistence int
}

type TodoArray []*Todo

// InsertTodo データベースをレコードを登録する
func InsertTodo(record *Todo) error {
	stmt, err := db.Conn.Prepare("INSERT INTO todo (todo_user_id, todo_article, todo_limit, todo_tag_id, todo_complete, todo_existence) VALUES (?, ?, ?, ?, ?, ?); ")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(record.TodoUserId, record.TodoArticle, record.TodoLimit, record.TodoTagId, record.TodoComplete, record.TodoExistence)
	return err
}

// UpdateTodoExistence Todoの存在を削除する
func UpdateTodoExistence(existenceNull, id int) error {
	_, err := db.Conn.Exec("UPDATE todo SET todo_existence = ? WHERE id = ?", existenceNull, id)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// SelectTodoList ユーザのIDからTodoのリストを取得する
func SelectTodoList(id int64, existenceFull int) (*TodoArray, error) {
	rows, err := db.Conn.Query("SELECT id, todo_article, todo_limit, todo_tag_id, todo_complete, todo_existence FROM todo WHERE todo_user_id = ? AND todo_existence = ?", id, existenceFull)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var todoList TodoArray
	for rows.Next() {
		var todo Todo
		if err := rows.Scan(&todo.TodoUserId, &todo.TodoArticle, &todo.TodoLimit, &todo.TodoTagId, &todo.TodoComplete, &todo.TodoExistence); err != nil {
			log.Println(err)
			return nil, err
		}
		todoList = append(todoList, &todo)
	}
	return &todoList, nil
}

// ユーザから送られたtag idを用いて，指定したタグのみを返す
func SelectTodoTag(id int64, tagID, existenceFull int) (*TodoArray, error) {
	rows, err := db.Conn.Query("SELECT id, todo_article, todo_limit, todo_tag_id, todo_complete, todo_existence FROM todo WHERE todo_user_id = ? AND todo_tag_id = ? AND todo_existence = ?", id, tagID, existenceFull)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var todoList TodoArray
	for rows.Next() {
		var todo Todo
		if err := rows.Scan(&todo.TodoUserId, &todo.TodoArticle, &todo.TodoLimit, &todo.TodoTagId, &todo.TodoComplete, &todo.TodoExistence); err != nil {
			log.Println(err)
			return nil, err
		}
		todoList = append(todoList, &todo)
	}
	return &todoList, nil
}
