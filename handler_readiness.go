package main

import "net/http"

func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	respondWithJson(w, http.StatusOK, map[string]string{"status": "ok"})
}

func handlerErr(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, http.StatusInternalServerError, "Internal Server Error")
}
