package api

import (
	"log"
	"net/http"

	"ToDoList/interface/server/handler/todoHandler"
	"ToDoList/interface/server/handler/userHandler"
)

// Serve HTTPサーバを起動する
func Serve(addr string) {

	http.HandleFunc("/user/signin", post(userHandler.HandleUserSignin()))
	http.HandleFunc("/user/login", post(userHandler.HandleUserLogin()))
	http.HandleFunc("/user/change", post(userHandler.HandleUserChange()))
	http.HandleFunc("/user/delete", post(userHandler.HandleUserDelete()))
	http.HandleFunc("/todo/add", post(todoHandler.HandleTodoRegistration()))
	http.HandleFunc("/todo/delete", post(todoHandler.HandleTodoDelete()))
	http.HandleFunc("/todo/list", post(todoHandler.HandleTodoGetList()))
	http.HandleFunc("/todo/search", post(todoHandler.HandleTodoSearch()))
	http.HandleFunc("/todo/complete", post(todoHandler.HandleTodoComplete()))

	/* ===== サーバの起動 ===== */
	log.Println("Server running...")
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatalf("Listen and serve failed. %+v", err)
	}
}

// get GETリクエストを処理する
func get(apiFunc http.HandlerFunc) http.HandlerFunc {
	return httpMethod(apiFunc, http.MethodGet)
}

// post POSTリクエストを処理する
func post(apiFunc http.HandlerFunc) http.HandlerFunc {
	return httpMethod(apiFunc, http.MethodPost)
}

// httpMethod 指定したHTTPメソッドでAPIの処理を実行する
func httpMethod(apiFunc http.HandlerFunc, method string) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

		// CORS対応
		writer.Header().Add("Access-Control-Allow-Origin", "*")
		writer.Header().Add("Access-Control-Allow-Headers", "Content-Type,Accept,Origin,x-token")

		// プリフライトリクエストは処理を通さない
		if request.Method == http.MethodOptions {
			return
		}
		// 指定のHTTPメソッドでない場合はエラー
		if request.Method != method {
			writer.WriteHeader(http.StatusMethodNotAllowed)
			writer.Write([]byte("Method Not Allowed"))
			return
		}

		// 共通のレスポンスヘッダを設定
		writer.Header().Add("Content-Type", "interface/json")
		apiFunc(writer, request)
	}
}
