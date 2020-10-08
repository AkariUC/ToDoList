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
func SelectUser(name string, password string, existence int) (error, *User) {
	row := db.Conn.QueryRow("SELECT auth_token FROM user WHERE name = ? AND password = ? AND existence = ?", name, password, existence)

	user := User{}
	err := row.Scan(&user.AuthToken)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		log.Println(err)
		return err, nil
	}
	return nil, &user
}
