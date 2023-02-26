package main

import (
	"fmt"
	"net/http"

	"github.com/GuillermoMajano/todo-app/router"
)

func main() {
	router := router.NewRouter()
	fmt.Println("listen and serverd")
	http.ListenAndServe(":8080", router)
}
