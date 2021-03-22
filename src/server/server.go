package server

import (
	"net/http"

	"github.com/gorilla/mux"
)

type AuthHandler interface {
	CheckAuth(next http.Handler) http.Handler
}

type UserHandler interface {
	SignUp(rw http.ResponseWriter, r *http.Request)
}

type TasksHandler interface {
	Create(rw http.ResponseWriter, r *http.Request)
	GetAll(rw http.ResponseWriter, r *http.Request)
}

type TaskCommentsHandler interface {
	Create(rw http.ResponseWriter, r *http.Request)
	Delete(rw http.ResponseWriter, r *http.Request)
}

type Server struct {
	Auth         AuthHandler
	User         UserHandler
	Tasks        TasksHandler
	TaskComments TaskCommentsHandler
}

func (s *Server) NewRouter() http.Handler {
	r := mux.NewRouter()

	root := r.PathPrefix("/todolist").Subrouter()
	root.Path("/users").HandlerFunc(s.User.SignUp).Methods(http.MethodPost)

	withAuth := root.PathPrefix("/").Subrouter()
	withAuth.Use(s.Auth.CheckAuth)

	withAuth.Path("/tasks").HandlerFunc(s.Tasks.Create).Methods(http.MethodPost)
	withAuth.Path("/tasks").HandlerFunc(s.Tasks.GetAll).Methods(http.MethodGet)

	taskComments := withAuth.PathPrefix("/tasks").Subrouter()
	taskComments.Path("/comments").HandlerFunc(s.TaskComments.Create).Methods(http.MethodPost)
	taskComments.Path("/{task_id}/comments/{comment_id}").HandlerFunc(s.TaskComments.Delete).Methods(http.MethodDelete)

	return r
}
