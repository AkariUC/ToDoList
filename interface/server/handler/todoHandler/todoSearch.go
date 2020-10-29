package todoHandler

import (
	"encoding/json"
	"log"
	"net/http"

	"ToDoList/constant"
	"ToDoList/infrastructure/model/todoModel"
	"ToDoList/infrastructure/model/userModel"
	"ToDoList/interface/server/response"
)

type TodoSearchRequest struct {
	TodoTagID int `json:"todo_tag_id"`
}

func HandleTodoSearch() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

		// リクエストBodyから更新後情報を取得
		var requestBody TodoSearchRequest
		if err := json.NewDecoder(request.Body).Decode(&requestBody); err != nil {
			log.Println(err)
			response.InternalServerError(writer, "Internal Server Error")
			return
		}

		// リクエストヘッダからuserを取得
		token := request.Header.Get("x-token")
		err, user := userModel.SelectUserData(token)
		if err != nil {
			log.Println(err)
			response.InternalServerError(writer, "Internal Server Error")
			return
		}

		// データベースから指定したユーザ，タグIDのTodoを取得
		todoArray, err := todoModel.SelectTodoTag(user.ID, requestBody.TodoTagID, constant.ExistenceFull)
		if err != nil {
			log.Println(err)
			response.InternalServerError(writer, "Internal Server Error")
			return
		}

		// つめる
		todos := make([]TodoList, len(*todoArray))
		for i, todo := range *todoArray {
			todos[i] = TodoList{
				TodoArticle:   todo.TodoArticle,
				TodoLimit:     todo.TodoLimit,
				TodoTagID:     todo.TodoTagId,
				TodoComplete:  todo.TodoComplete,
				TodoExistence: todo.TodoExistence,
			}
		}

		// レスポンス
		response.Success(writer, &TodoListResponse{
			TodoList: todos,
		})
	}
}
