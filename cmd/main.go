package main

import (
	"ToDoList/cmd/api"
	"flag"
)

var (
	// Listenするアドレス+ポート
	addr string
)

func init() {
	flag.StringVar(&addr, "addr", ":8080", "tcp host:port to connect")
	flag.Parse()
}

func main() {
	api.Serve(addr)
}
