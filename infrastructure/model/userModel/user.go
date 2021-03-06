package userModel

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"

	"ToDoList/infrastructure/db"
)

type User struct {
	ID        int64
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
	return err
}

// SelectUser ユーザを取得
func SelectUser(name, password string, existence int) (error, *User) {
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

// SelectUserData トークンからユーザを取得
func SelectUserData(token string) (error, *User) {
	row := db.Conn.QueryRow("SELECT id, name, password, auth_token, existence FROM user WHERE auth_token = ?", token)

	user := User{}
	err := row.Scan(&user.ID, &user.Name, &user.PassWord, &user.AuthToken, &user.Existence)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		log.Println(err)
		return err, nil
	}
	return nil, &user
}

// UpdateUserName ユーザの名前を変更
func UpdateUserName(name, newName, password string) error {
	_, err := db.Conn.Exec("UPDATE user SET name = ? WHERE name = ? AND password = ?", newName, name, password)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// UpdateUserExistence ユーザの存在の有無
func UpdateUserExistence(name, password string, existenceNull int) error {
	_, err := db.Conn.Exec("UPDATE user SET existence = ? WHERE name = ? AND password = ?", existenceNull, name, password)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
