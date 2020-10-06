package userModel

import (
	"ToDoList/infrastructure/db"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Name      string
	PassWord  string
	AuthToken string
}

// InsertUser データベースをレコードを登録する
func InsertUser(record *User) error {
	stmt, err := db.Conn.Prepare("INSERT INTO user (name, password, auth_token) VALUES (?, ?, ?); ")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(record.Name, record.PassWord, record.AuthToken)
	return err
}

// todo ログイン時にユーザから送られてきたデータをもとに、DBのユーザを探し、トークンを返す
func SelectUser(record *User) error {
	stmt, err := db.Conn.Prepare("SELECT auth_token FROM user WHERE id = ? AND name = ? AND password = ? AND auth_token = ? AND existence = 1")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(record.Name, record.PassWord, record.AuthToken)
	// 仮return
	return err
}

// todo 変換
