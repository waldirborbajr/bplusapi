package handler

import "net/http"

func HandlerCustom(w http.ResponseWriter, r *http.Request) {
	reponseWithJSON(w, http.StatusOK, struct{}{})
}
