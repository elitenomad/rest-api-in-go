package http

import (
	"encoding/json"
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

type Response struct {
	Message string
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
		rw.Header().Add("Content-Type", "application/json")
		rw.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(rw).Encode(Response{Message: "I am Alive"}); err != nil {
			panic(err)
		}
	})

}

func (h *Handler) GetComment(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

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

	if err := json.NewEncoder(w).Encode(comment); err != nil {
		panic(err)
	}
}

func (h *Handler) DeleteComment(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

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

	if err := json.NewEncoder(w).Encode(Response{Message: "Successfully deleted the comment for the given ID"}); err != nil {
		panic(err)
	}
}

func (h *Handler) GetAllComments(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	comments := h.Service.GetAllComments()
	if err := json.NewEncoder(w).Encode(comments); err != nil {
		panic(err)
	}
}

func (h *Handler) UpdateComment(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	vars := mux.Vars(r)
	id := vars["id"]

	idx, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		fmt.Println("Unable to Parse the id sent")
	}

	var comment comment.Comment
	if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
		fmt.Fprintf(w, "Failed to decode the JSON body - comment")
	}

	comment, err = h.Service.UpdateComment(idx, comment)
	if err != nil {
		fmt.Fprintf(w, "Error in Updating a Comment")
	}

	if err := json.NewEncoder(w).Encode(comment); err != nil {
		panic(err)
	}
}

func (h *Handler) PostComment(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	var comment comment.Comment
	if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
		fmt.Fprintf(w, "Failed to decode the JSON body - comment")
	}

	comment, err := h.Service.PostComment(comment)

	if err != nil {
		fmt.Fprintf(w, "Error in posting a Comment")
	}

	if err := json.NewEncoder(w).Encode(comment); err != nil {
		panic(err)
	}
}
