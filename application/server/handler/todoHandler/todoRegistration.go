package todoHandler

import (
	"encoding/json"
	"log"
	"net/http"

	"ToDoList/application/server/response"
	"ToDoList/constant"
	"ToDoList/domain/model/tagModel"
	"ToDoList/domain/model/todoModel"
	"ToDoList/domain/model/userModel"
)

type TodoRegistrationRequest struct {
	TodoArticle string `json:"todo_article"`
	TodoLimit   string `json:"todo_limit"`
	TodoTagID   int    `json:"todo_tag_id"`
}

func HandleTodoRegistration() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

		// リクエストBodyから更新後情報を取得
		var requestBody TodoRegistrationRequest
		if err := json.NewDecoder(request.Body).Decode(&requestBody); err != nil {
			log.Println(err)
			response.InternalServerError(writer, "Internal Server Error")
			return
		}

		// リクエストヘッダからauth_tokenを取得
		token := request.Header.Get("x-token")
		err, user := userModel.SelectUserData(token)
		if err != nil {
			log.Println(err)
			response.InternalServerError(writer, "Internal Server Error")
			return
		}

		// tagデータを取得
		err, tag := tagModel.SelectTag(requestBody.TodoTagID)
		if err != nil {
			log.Println(err)
			response.InternalServerError(writer, "Internal Server Error")
			return
		}

		// データベースにTodoデータを登録する
		err = todoModel.InsertTodo(
			&todoModel.Todo{
				TodoUserId:    user.ID,
				TodoArticle:   requestBody.TodoArticle,
				TodoLimit:     requestBody.TodoLimit,
				TodoTagId:     tag.ID,
				TodoComplete:  constant.CompleteNot,
				TodoExistence: constant.ExistenceFull,
			},
		)
		if err != nil {
			log.Println(err)
			response.InternalServerError(writer, "Internal Server Error")
			return
		}
	}
}
