package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Animechan struct {
	Anime     string `json:"anime"`
	Character string `json:"character"`
	Quote     string `json:"quote"`
}

func ClientGet() ([]Animechan, error) {
	client := http.Client{}
	// Hit API https://animechan.vercel.app/api/quotes/anime?title=naruto with method POST:
	res, err := client.Get("https://animechan.vercel.app/api/quotes/anime?title=naruto")
	if err != nil {
		return []Animechan{}, err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return []Animechan{}, err
	}

	var data []Animechan
	err = json.Unmarshal(body, &data)
	if err != nil {
		return []Animechan{}, err
	}
	return data, nil // TODO: replace this
}

type data struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

type Postman struct {
	Data data
	Url  string `json:"url"`
}

func ClientPost() (Postman, error) {
	postBody, _ := json.Marshal(map[string]string{
		"name":  "Dion",
		"email": "dionbe2022@gmail.com",
	})
	responseBody := bytes.NewBuffer(postBody)
	fmt.Println(responseBody)

	// Hit API https://postman-echo.com/post with method POST:
	res, err := http.Post("https://postman-echo.com/post", "application/json", responseBody)
	if err != nil {
		return Postman{}, err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return Postman{}, err
	}

	var postman Postman
	err = json.Unmarshal(body, &postman)
	if err != nil {
		return Postman{}, err
	}
	return postman, nil // TODO: replace this
}

func main() {
	get, _ := ClientGet()
	fmt.Println(get)

	post, _ := ClientPost()
	fmt.Println(post)
}
