package userModel

import (
	"ToDoList/infrastructure/db"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Name      string
	PassWord  string
	AuthToken string
	Coin      int32
}

// InsertUser データベースをレコードを登録する
func InsertUser(record *User) error {
	stmt, err := db.Conn.Prepare("INSERT INTO user (name, password, auth_token, coin) VALUES (?, ?, ?, ?); ")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(record.Name, record.PassWord, record.AuthToken, record.Coin)
	return err
}

// todo ログイン時にユーザから送られてきたデータをもとに、DBのユーザを探し、トークンを返す

// todo 変換