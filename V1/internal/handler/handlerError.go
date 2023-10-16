package handler

import (
	"net/http"
)

func HandlerError(w http.ResponseWriter, r *http.Request) {
	reponseWithError(w, http.StatusBadRequest, "Error found")
}
