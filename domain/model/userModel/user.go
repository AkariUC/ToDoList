package userModel

import (
	"ToDoList/infrastructure/db"
	"database/sql"
	"log"

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
func SelectUser(name, password string) (error, *string) {
	row := db.Conn.QueryRow("SELECT auth_token FROM user WHERE id = ? AND name = ? AND password = ? AND existence = 1")

	user := User{}
	err := row.Scan(&user.Name, &user.AuthToken)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		log.Println(err)
		return err, nil
	}
	return nil, &user.AuthToken
}

// todo 変換
