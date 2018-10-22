package Controller

import (
	"encoding/json"
	"net/http"
)

type Article struct {
	Title string
	Body  string
	Tags  []string
}

func (a *Article) Get(res http.ResponseWriter, req *http.Request) {
	a.Title = "hello"
	a.Body = "this is body"
	a.Tags = append(a.Tags, "ok")
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(200)
	json.NewEncoder(res).Encode(a)
}