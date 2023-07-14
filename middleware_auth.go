package main

import (
	"net/http"

	"github.com/rodmedeiross/scratch/internal/auth"
	"github.com/rodmedeiross/scratch/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
		respondWithError(w, 403, err.Error())
		return
	}

		user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			respondWithError(w, 400, err.Error())
			return
		}
		handler(w, r, user)
	}
}
