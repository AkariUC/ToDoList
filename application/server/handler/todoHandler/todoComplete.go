package todoHandler

import (
	"encoding/json"
	"log"
	"net/http"

	"ToDoList/application/server/response"
	"ToDoList/constant"
	"ToDoList/domain/model/todoModel"
)

type TodoCompleteRequest struct {
	ID int `json:"id"`
}

func HandleTodoComplete() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

		// リクエストBodyから更新後情報を取得
		var requestBody TodoCompleteRequest
		if err := json.NewDecoder(request.Body).Decode(&requestBody); err != nil {
			log.Println(err)
			response.InternalServerError(writer, "Internal Server Error")
			return
		}

		// データベースから指定したユーザ，タグIDのTodoを取得
		err := todoModel.UpdateTodoComplete(requestBody.ID, constant.Complete)
		if err != nil {
			log.Println(err)
			response.InternalServerError(writer, "Internal Server Error")
			return
		}
	}
}
