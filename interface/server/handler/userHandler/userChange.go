package userHandler

import (
	"encoding/json"
	"log"
	"net/http"

	"ToDoList/infrastructure/model/userModel"
	"ToDoList/interface/server/response"
)

type UserChangeRequest struct {
	Name     string `json:"name"`
	NewName  string `json:"newName"`
	PassWord string `json:"password"`
}

func HandleUserChange() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

		// リクエストBodyから更新後情報を取得
		var requestBody UserChangeRequest
		if err := json.NewDecoder(request.Body).Decode(&requestBody); err != nil {
			log.Println(err)
			response.InternalServerError(writer, "Internal Server Error")
			return
		}

		// データベースのユーザデータを書き換える
		err := userModel.UpdateUserName(
			requestBody.Name,
			requestBody.NewName,
			requestBody.PassWord,
		)
		if err != nil {
			log.Println(err)
			response.InternalServerError(writer, "Internal Server Error")
			return
		}
	}
}
