package http

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/elitenomad/rest-api/internal/comment"
	"github.com/gorilla/mux"
)

type Handler struct {
	Router  *mux.Router
	Service *comment.Service
}

func NewHandler(service *comment.Service) *Handler {
	return &Handler{
		Service: service,
	}
}

func (h *Handler) SetupRoutes() {
	fmt.Println("Setting up routes...")
	h.Router = mux.NewRouter()
	h.Router.HandleFunc("/api/comments", h.GetAllComments).Methods("GET")
	h.Router.HandleFunc("/api/comments", h.PostComment).Methods("POST")
	h.Router.HandleFunc("/api/comments/{id}", h.GetComment).Methods("GET")
	h.Router.HandleFunc("/api/comments/{id}", h.UpdateComment).Methods("PUT")
	h.Router.HandleFunc("/api/comments/{id}", h.DeleteComment).Methods("DELETE")
	h.Router.HandleFunc("/api/health", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(rw, "I am loading not alright...")
	})

}

func (h *Handler) GetComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	idx, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		fmt.Println("Unable to parse the String sent")
	}

	comment, err := h.Service.GetComment(idx)
	if err != nil {
		fmt.Fprintf(w, "Error in returning Comment for the ID given")
	}

	fmt.Fprintf(w, "%+v", comment)
}

func (h *Handler) DeleteComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	idx, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		fmt.Println("Unable to Parse the id sent")
	}

	err = h.Service.DeleteComment(uint(idx))
	if err != nil {
		fmt.Fprintf(w, "Error in deleting Comment for the ID given")
	}

	fmt.Fprintf(w, "Successfully deleted the comment for ID given")
}

func (h *Handler) GetAllComments(w http.ResponseWriter, r *http.Request) {
	comments := h.Service.GetAllComments()
	fmt.Fprintf(w, "%+v", comments)
}

func (h *Handler) UpdateComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	idx, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		fmt.Println("Unable to Parse the id sent")
	}

	comment, err := h.Service.UpdateComment(idx, comment.Comment{
		Slug: "/",
	})

	if err != nil {
		fmt.Fprintf(w, "Error in Updating a Comment")
	}

	fmt.Fprintf(w, "%+v", comment)
}

func (h *Handler) PostComment(w http.ResponseWriter, r *http.Request) {
	comment, err := h.Service.PostComment(comment.Comment{
		Slug: "/",
	})

	if err != nil {
		fmt.Fprintf(w, "Error in posting a Comment")
	}

	fmt.Fprintf(w, "%+v", comment)
}
