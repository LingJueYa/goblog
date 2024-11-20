package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Hello, 这里是 goblog</h1>")
}
func main() {
	http.HandleFunc("/", handlerFunc)
	err := http.ListenAndServe(":3001", nil)
	if err != nil {
		fmt.Println("服务器启动失败:", err)
	}
}
