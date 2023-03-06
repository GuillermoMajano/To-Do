package router

import (
	"net/http"

	"github.com/GuillermoMajano/todo-app/middleware"
	"github.com/gorilla/mux"
)

func NewRouter() http.Handler

func Router() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("api/task", middleware.GetAllTasks).Methods("GET", "OPTIONS")
	r.HandleFunc("api/tasks", middleware.CreateTask).Methods("POST", "OPTIONS")
	r.HandleFunc("api/tasks/{id}", middleware.TaskComplete).Methods("PUT", "OPTIONS")
	r.HandleFunc("api/undoTask/{id}", middleware.UndoTask).Methods("PUT", "OPTIONS")
	r.HandleFunc("api/delateTasks", middleware.DeleteTask).Methods("DELETE", "OPTIONS")
	r.HandleFunc("api/delateAllTasks", middleware.DeleteAllTasks).Methods("DELETE", "OPTIONS")
	return r
}
