package api

import (
	"a21hc3NpZ25tZW50/model"
	"encoding/json"
	"net/http"
	"path"
	"text/template"
	"time"

	"github.com/google/uuid"
)

func (api *API) Register(w http.ResponseWriter, r *http.Request) {
	// Read username and password request with FormValue.
	if r.Body == http.NoBody {
		w.WriteHeader(http.StatusBadRequest)
		var error model.ErrorResponse
		error.Error = "Username or Password empty"
		jsonData, err := json.Marshal(error)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(jsonData)
		return
	}
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
        return
	}
	creds := model.Credentials{} // TODO: replace this
	creds.Username = r.FormValue("username")
	creds.Password = r.FormValue("password")

	// Handle request if creds is empty send response code 400, and message "Username or Password empty"
	// TODO: answer here
	if creds.Username == "" && creds.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		var error model.ErrorResponse
		error.Error = "Username or Password empty"
		jsonData, err := json.Marshal(error)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(jsonData)
		return
	}

	w.WriteHeader(http.StatusOK)
	err := api.usersRepo.AddUser(creds)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: "Internal Server Error"})
		return
	}

	filepath := path.Join("views", "status.html")
	tmpl, err := template.ParseFiles(filepath)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: "Internal Server Error"})
		return
	}

	var data = map[string]string{"name": creds.Username, "message": "register success!"}
	err = tmpl.Execute(w, data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: "Internal Server Error"})
	}
}

func (api *API) Login(w http.ResponseWriter, r *http.Request) {
	// Read usernmae and password request with FormValue.
	if r.Body == http.NoBody {
		w.WriteHeader(http.StatusBadRequest)
		var error model.ErrorResponse
		error.Error = "Username or Password empty"
		jsonData, err := json.Marshal(error)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(jsonData)
		return
	}
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
        return
	}
	creds := model.Credentials{} // TODO: replace this
	creds.Username = r.FormValue("username")
	creds.Password = r.FormValue("password")

	// Handle request if creds is empty send response code 400, and message "Username or Password empty"
	// TODO: answer here
	if creds.Username == "" && creds.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		var error model.ErrorResponse
		error.Error = "Username or Password empty"
		jsonData, err := json.Marshal(error)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(jsonData)
		return
	}

	err := api.usersRepo.LoginValid(creds)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: err.Error()})
		return
	}

	// Generate Cookie with Name "session_token", Path "/", Value "uuid generated with github.com/google/uuid", Expires time to 5 Hour.
	// TODO: answer here
	value := uuid.New().String()
	expiration := time.Now().Add(time.Hour * 5)
	http.SetCookie(w, &http.Cookie{Name: "session_token", Value: value, Expires: expiration, Path: "/"})
	
	w.WriteHeader(http.StatusOK)
	session := model.Session{} // TODO: replace this
	session.Token = value
	session.Username = creds.Username
	session.Expiry = expiration
	err = api.sessionsRepo.AddSessions(session)

	api.dashboardView(w, r)
}

func (api *API) Logout(w http.ResponseWriter, r *http.Request) {
	//Read session_token and get Value:
	cookie, _ := r.Cookie("session_token")
	sessionToken := cookie.Value // TODO: replace this

	api.sessionsRepo.DeleteSessions(sessionToken)

	//Set Cookie name session_token value to empty and set expires time to Now:
	// TODO: answer here
	http.SetCookie(w, &http.Cookie{Name: "session_token", Value: "", Expires: time.Now()})
	w.WriteHeader(http.StatusOK)
	filepath := path.Join("views", "login.html")
	tmpl, err := template.ParseFiles(filepath)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: "Internal Server Error"})
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(model.ErrorResponse{Error: "Internal Server Error"})
	}
}
