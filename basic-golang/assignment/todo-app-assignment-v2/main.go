package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"a21hc3NpZ25tZW50/db"
	"a21hc3NpZ25tZW50/middleware"
	"a21hc3NpZ25tZW50/model"

	"github.com/google/uuid"
)

func Register(w http.ResponseWriter, r *http.Request) {
	// TODO: answer here
	if r.Body == http.NoBody {
		w.WriteHeader(http.StatusBadRequest)
		var error model.ErrorResponse
		error.Error = "Internal Server Error"
		jsonData, err := json.Marshal(error)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(jsonData)
		return
	}
	var user model.Credentials
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if user.Username == "" && user.Password == "" {
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
	for val := range db.Users {
		if user.Username == val {
			w.WriteHeader(http.StatusConflict)
			var error model.ErrorResponse
			error.Error = "Username already exist"
			jsonData, err := json.Marshal(error)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Write(jsonData)
			return
		}
	}
	w.WriteHeader(http.StatusOK)
	var res model.SuccessResponse
	res.Username = user.Username
	res.Message = "Register Success"
	db.Users[user.Username] = user.Password
	jsonData, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(jsonData)
}

func Login(w http.ResponseWriter, r *http.Request) {
	// TODO: answer here
	if r.Body == http.NoBody {
		w.WriteHeader(http.StatusBadRequest)
		var error model.ErrorResponse
		error.Error = "Internal Server Error"
		jsonData, err := json.Marshal(error)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(jsonData)
		return
	}
	var user model.Credentials
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if user.Username == "" && user.Password == "" {
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
	for username, password := range db.Users {
		if username == user.Username && password == user.Password {
			value := uuid.New().String()
			expiration := time.Now().Add(time.Hour * 5)
			http.SetCookie(w, &http.Cookie{Name: "session_token", Value: value, Expires: expiration, Path: "/"})

			w.WriteHeader(http.StatusOK)
			var res model.SuccessResponse
			res.Username = username
			res.Message = "Login Success"
			db.Sessions[value] = model.Session{
				Username: username,
				Expiry:   expiration,
			}
			jsonData, err := json.Marshal(res)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Write(jsonData)
			return
		}
	}
	w.WriteHeader(http.StatusUnauthorized)
	var error model.ErrorResponse
	error.Error = "Wrong User or Password!"
	jsonData, err := json.Marshal(error)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(jsonData)
}

func AddToDo(w http.ResponseWriter, r *http.Request) {
	// TODO: answer here
	cookie, _ := r.Cookie("session_token")
	var todo model.Todo
	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	var res model.SuccessResponse
	res.Username = db.Sessions[cookie.Value].Username
	res.Message = fmt.Sprintf("Task %s added!", todo.Task)
	db.Task[res.Username] = append(db.Task[res.Username], model.Todo{
		Id:   uuid.New().String(),
		Task: todo.Task,
		Done: todo.Done,
	})
	jsonData, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(jsonData)
}

func ListToDo(w http.ResponseWriter, r *http.Request) {
	// TODO: answer here
	cookie, _ := r.Cookie("session_token")
	username := db.Sessions[cookie.Value].Username
	userTodos, ok := db.Task[username]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		var error model.ErrorResponse
		error.Error = "Todolist not found!"
		jsonData, err := json.Marshal(error)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(jsonData)
		return
	}
	w.WriteHeader(http.StatusOK)
	var todos []model.Todo
	for _, val := range userTodos {
		todos = append(todos, model.Todo{
			Id:   val.Id,
			Task: val.Task,
			Done: val.Done,
		})
	}
	jsonData, err := json.Marshal(todos)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(jsonData)
}

func ClearToDo(w http.ResponseWriter, r *http.Request) {
	// TODO: answer here
	cookie, _ := r.Cookie("session_token")
	w.WriteHeader(http.StatusOK)
	var res model.SuccessResponse
	res.Username = db.Sessions[cookie.Value].Username
	res.Message = "Clear ToDo Success"
	db.Task[res.Username] = []model.Todo{}
	jsonData, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(jsonData)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	// TODO: answer here
	cookie, _ := r.Cookie("session_token")
	w.WriteHeader(http.StatusOK)
	var res model.SuccessResponse
	res.Username = db.Sessions[cookie.Value].Username
	res.Message = "Logout Success"
	delete(db.Sessions, cookie.Value)
	jsonData, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(jsonData)
}

func ResetToDo(w http.ResponseWriter, r *http.Request) {
	db.Task = map[string][]model.Todo{}
	w.WriteHeader(http.StatusOK)
}

type API struct {
	mux *http.ServeMux
}

func NewAPI() API {
	mux := http.NewServeMux()
	api := API{
		mux,
	}

	mux.Handle("/user/register", middleware.Post(http.HandlerFunc(Register)))
	mux.Handle("/user/login", middleware.Post(http.HandlerFunc(Login)))
	mux.Handle("/user/logout", middleware.Get(middleware.Auth(http.HandlerFunc(Logout))))

	// TODO: answer here
	mux.Handle("/todo/create", middleware.Post(middleware.Auth(http.HandlerFunc(AddToDo))))
	mux.Handle("/todo/read", middleware.Get(middleware.Auth(http.HandlerFunc(ListToDo))))
	mux.Handle("/todo/clear", middleware.Delete(middleware.Auth(http.HandlerFunc(ClearToDo))))

	mux.Handle("/todo/reset", http.HandlerFunc(ResetToDo))

	return api
}

func (api *API) Handler() *http.ServeMux {
	return api.mux
}

func (api *API) Start() {
	fmt.Println("starting web server at http://localhost:8080")
	http.ListenAndServe(":8080", api.Handler())
}

func main() {
	mainAPI := NewAPI()
	mainAPI.Start()
}
