package api

import (
	"a21hc3NpZ25tZW50/model"
	"context"
	"encoding/json"
	"net/http"
)

func (api *API) Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("session_token")
		if err != nil {
			if err == http.ErrNoCookie {
				w.WriteHeader(http.StatusUnauthorized)
				var error model.ErrorResponse
				error.Error = "http: named cookie not present"
				jsonData, err := json.Marshal(error)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				w.Write(jsonData)
				return
			}
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		sessionToken := cookie.Value // TODO: replace this

		sessionFound, err := api.sessionsRepo.CheckExpireToken(sessionToken)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(model.ErrorResponse{Error: err.Error()})
			return
		}

		ctx := context.WithValue(r.Context(), "username", sessionFound.Username)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (api *API) Get(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO: answer here
		if r.Method != "GET" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			var error model.ErrorResponse
			error.Error = "Method is not allowed!"
			jsonData, err := json.Marshal(error)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Write(jsonData)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func (api *API) Post(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO: answer here
		if r.Method != "POST" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			var error model.ErrorResponse
			error.Error = "Method is not allowed!"
			jsonData, err := json.Marshal(error)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Write(jsonData)
			return
		}
		next.ServeHTTP(w, r)
	})
}
