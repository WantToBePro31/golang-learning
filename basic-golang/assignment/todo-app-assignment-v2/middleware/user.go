package middleware

import (
	"encoding/json"
	"net/http"
	"time"

	"a21hc3NpZ25tZW50/model"
)

func isExpired(s model.Session) bool {
	return s.Expiry.Before(time.Now())
}

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := r.Cookie("session_token")
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
		} else {
			next.ServeHTTP(w, r)
		}
	}) // TODO: replace this
}
