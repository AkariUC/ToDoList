package userHandler

import (
	"ToDoList/application/server/response"
	"encoding/json"
	"log"
	"net/http"
)

type UserSigninRequest struct {
	Name     string `json:"name"`
	PassWord string `json:"password"`
}

func HandleUserSignin() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

		// リクエストBodyから更新後情報を取得
		var requestBody UserSigninRequest
		if err := json.NewDecoder(request.Body).Decode(&requestBody); err != nil {
			log.Println(err)
			response.InternalServerError(writer, "Internal Server Error")
			return
		}

		// UUIDで認証トークンを生成する
		authToken, err := uuid.NewRandom()
		if err != nil {
			log.Println(err)
			response.InternalServerError(writer, "Internal Server Error")
			return
		}

		// データベースにユーザデータを登録する
		// ユーザデータの登録クエリを入力する
		err = userModel.InsertUser(&userModel.User{
			Name:      requestBody.Name,
			PassWord:  requestBody.PassWord,
			AuthToken: authToken.String(),
			Coin:      0,
		})
		if err != nil {
			log.Println(err)
			response.InternalServerError(writer, "Internal Server Error")
			return
		}
	}
}