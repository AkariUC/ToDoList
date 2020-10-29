package todoHandler

import (
	"log"
	"net/http"

	"ToDoList/constant"
	"ToDoList/infrastructure/model/todoModel"
	"ToDoList/infrastructure/model/userModel"
	"ToDoList/interface/server/response"
)

type TodoListResponse struct {
	TodoList []TodoList `json:"todo_list"`
}

type TodoList struct {
	TodoArticle   string `json:"todo_article"`
	TodoLimit     string `json:"todo_limit"`
	TodoTagID     int64  `json:"todo_tag_id"`
	TodoComplete  int    `json:"todo_complete"`
	TodoExistence int    `json:"todo_existence"`
}

func HandleTodoGetList() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		// リクエストヘッダからauth_tokenを取得
		token := request.Header.Get("x-token")
		err, user := userModel.SelectUserData(token)
		if err != nil {
			log.Println(err)
			response.InternalServerError(writer, "Internal Server Error")
			return
		}

		// データベースから指定したユーザのTodoを取得
		todoArray, err := todoModel.SelectTodoList(user.ID, constant.ExistenceFull)
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
