package userHandler

import (
	"encoding/json"
	"log"
	"net/http"

	"ToDoList/constant"
	"ToDoList/infrastructure/model/userModel"
	"ToDoList/interface/server/response"
)

type UserLoginRequest struct {
	Name     string `json:"name"`
	PassWord string `json:"password"`
}

type UserLoginResponse struct {
	Token string `json:"token"`
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
		err, user := userModel.SelectUser(requestBody.Name, requestBody.PassWord, constant.ExistenceFull)
		if err != nil {
			log.Println(err)
			response.InternalServerError(writer, "Internal Server Error")
			return
		}

		// レスポンス
		response.Success(writer, &UserLoginResponse{
			Token: user.AuthToken,
		})
	}
}
