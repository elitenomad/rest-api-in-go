package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/elitenomad/rest-api/internal/comment"
	"github.com/gorilla/mux"
)

type Handler struct {
	Router  *mux.Router
	Service *comment.Service
}

type Response struct {
	Message string
	Error   string
}

func NewHandler(service *comment.Service) *Handler {
	return &Handler{
		Service: service,
	}
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.WithFields(
			log.Fields{
				"Method": r.Method,
				"Path":   r.URL.Path,
			}).
			Info("handled request")
		next.ServeHTTP(w, r)
	})
}

// BasicAuth - a handy middleware function that logs out incoming requests
func BasicAuth(original func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		user, pass, ok := r.BasicAuth()
		if user == "admin" && pass == "password" && ok {
			original(w, r)
		} else {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			sendHttpResponseError(w, "not authorized", errors.New("not authorized"))
		}
	}
}

func (h *Handler) SetupRoutes() {
	fmt.Println("Setting up routes...")
	h.Router = mux.NewRouter()
	h.Router.Use(LoggingMiddleware)

	h.Router.HandleFunc("/api/comments", h.GetAllComments).Methods("GET")
	h.Router.HandleFunc("/api/comments", BasicAuth(h.PostComment)).Methods("POST")
	h.Router.HandleFunc("/api/comments/{id}", h.GetComment).Methods("GET")
	h.Router.HandleFunc("/api/comments/{id}", h.UpdateComment).Methods("PUT")
	h.Router.HandleFunc("/api/comments/{id}", h.DeleteComment).Methods("DELETE")
	h.Router.HandleFunc("/api/health", func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Add("Content-Type", "application/json")
		rw.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(rw).Encode(Response{Message: "I am Alive"}); err != nil {
			panic(err)
		}
	})

}

func sendHttpResponseError(w http.ResponseWriter, message string, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	if err := json.NewEncoder(w).Encode(Response{Message: message, Error: err.Error()}); err != nil {
		panic(err)
	}
}
