package userHandler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"ToDoList/application/server/response"
	"ToDoList/domain/model/userModel"
)

// todo ログインハンドラ作成
// todo domain/mode/userModelに ユーザから入力された値を送る
//  返り値はtokenを返すが、DBから送られてきた形式は、Goではそのまま使えないので(ライブラリを導入すればそのまま使える)そこまででOK

type UserLoginRequest struct {
	Name     string `json:"name"`
	PassWord string `json:"password"`
}

func HandleUserLogin() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		// リクエストBodyから更新後情報を取得
		var requestBody UserLoginRequest
		if err := json.NewDecoder(request.Body).Decode(&requestBody); err != nil {
			log.Println(err)
			response.InternalServerError(writer, "Internal Server Error")
			return
		}

		// データベースにユーザデータが存在するか確認する
		err, token := userModel.SelectUser(requestBody.Name, requestBody.PassWord)
		if err != nil {
			log.Println(err)
			response.InternalServerError(writer, "Internal Server Error")
			return
		}

		fmt.Println(token)
	}
}
