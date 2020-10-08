package userHandler

import (
	"encoding/json"
	"log"
	"net/http"

	"ToDoList/application/server/response"
	"ToDoList/domain/model/userModel"
)

type UserDeleteRequest struct {
	Name     string `json:"name"`
	PassWord string `json:"password"`
}

func HandleUserDelete() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

		// リクエストBodyから更新後情報を取得
		var requestBody UserSigninRequest
		if err := json.NewDecoder(request.Body).Decode(&requestBody); err != nil {
			log.Println(err)
			response.InternalServerError(writer, "Internal Server Error")
			return
		}

		// ユーザデータを検索する
		err := userModel.UpdateUserExistence(
			requestBody.Name,
			requestBody.PassWord,
		)
		if err != nil {
			log.Println(err)
			response.InternalServerError(writer, "Internal Server Error")
			return
		}
	}
}
