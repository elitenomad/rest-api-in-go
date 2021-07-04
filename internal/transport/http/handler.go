package http

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct {
	Router *mux.Router
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) SetupRoutes() {
	fmt.Println("Setting up routes...")
	h.Router = mux.NewRouter()
	h.Router.HandleFunc("/api/health", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(rw, "I am loading alright...")
	})
}
