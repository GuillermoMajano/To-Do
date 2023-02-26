package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter() http.Handler

func Router() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("api/task", func(http.ResponseWriter, *http.Request) {}).Methods("GET", "OPTIONS")
	r.HandleFunc("api/tasks", func(http.ResponseWriter, *http.Request) {}).Methods("POST", "OPTIONS")
	r.HandleFunc("api/tasks/{id}", func(http.ResponseWriter, *http.Request) {}).Methods("PUT", "OPTIONS")
	r.HandleFunc("api/undoTask/{id}", func(http.ResponseWriter, *http.Request) {}).Methods("PUT", "OPTIONS")
	r.HandleFunc("api/delateTasks", func(http.ResponseWriter, *http.Request) {}).Methods("DELETE", "OPTIONS")
	r.HandleFunc("api/delateAllTasks", func(http.ResponseWriter, *http.Request) {}).Methods("DELETE", "OPTIONS")
	return r
}
