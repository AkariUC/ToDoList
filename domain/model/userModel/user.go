package userModel

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"

	"ToDoList/infrastructure/db"
)

type User struct {
	Name      string
	PassWord  string
	AuthToken string
	Existence int32
}

// InsertUser データベースをレコードを登録する
func InsertUser(record *User) error {
	stmt, err := db.Conn.Prepare("INSERT INTO user (name, password, auth_token, existence) VALUES (?, ?, ?, ?); ")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(record.Name, record.PassWord, record.AuthToken, record.Existence)

	fmt.Println(record.Name)
	fmt.Println(record.PassWord)
	fmt.Println(record.AuthToken)
	fmt.Println(record.Existence)

	return err
}

// todo ログイン時にユーザから送られてきたデータをもとに、DBのユーザを探し、トークンを返す
func SelectUser(name, password string) (error, *string) {
	// todo 2 existence = 1 は良くないので `constant` を活用
	// todo 2 SELECT 文の `?` が どこにも定義されていないので WHEREが死んでいる
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
