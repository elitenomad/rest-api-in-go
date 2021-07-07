package http

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/elitenomad/rest-api/internal/comment"
	"github.com/gorilla/mux"
)

func (h *Handler) GetComment(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	vars := mux.Vars(r)
	id := vars["id"]

	idx, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		sendHttpResponseError(w, "Unable to parse the String sent", err)
	}

	comment, err := h.Service.GetComment(idx)
	if err != nil {
		sendHttpResponseError(w, "Error in returning Comment for the ID given", err)
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
		sendHttpResponseError(w, "Unable to Parse the id sent", err)
	}

	err = h.Service.DeleteComment(uint(idx))
	if err != nil {
		sendHttpResponseError(w, "Error in deleting Comment for the ID given", err)
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
		sendHttpResponseError(w, "Unable to Parse the id sent", err)
	}

	var comment comment.Comment
	if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
		sendHttpResponseError(w, "Failed to decode the JSON body - comment", err)
	}

	comment, err = h.Service.UpdateComment(idx, comment)
	if err != nil {
		sendHttpResponseError(w, "Error in Updating a Comment", err)
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
		sendHttpResponseError(w, "Failed to decode the JSON body - comment", err)
	}

	comment, err := h.Service.PostComment(comment)
	if err != nil {
		sendHttpResponseError(w, "Error in posting a Comment", err)
	}

	if err := json.NewEncoder(w).Encode(comment); err != nil {
		panic(err)
	}
}
