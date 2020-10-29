package todoHandler

import (
	"encoding/json"
	"log"
	"net/http"

	"ToDoList/constant"
	"ToDoList/infrastructure/model/todoModel"
	"ToDoList/interface/server/response"
)

type TodoDeleteRequest struct {
	ID int `json:"id"`
}

func HandleTodoDelete() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

		// リクエストBodyから更新後情報を取得
		var requestBody TodoDeleteRequest
		if err := json.NewDecoder(request.Body).Decode(&requestBody); err != nil {
			log.Println(err)
			response.InternalServerError(writer, "Internal Server Error")
			return
		}

		// Todoのデータを書き換える
		err := todoModel.UpdateTodoExistence(constant.ExistenceNull, requestBody.ID)
		if err != nil {
			log.Println(err)
			response.InternalServerError(writer, "Internal Server Error")
			return
		}
	}
}
