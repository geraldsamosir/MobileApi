package Controller

import (
	"encoding/json"
	"net/http"
)

type Article struct {
	Title string
	body  string
	tags  []string
}

func (a *Article) Get(res http.ResponseWriter, req *http.Request) {
	a.Title = "hello"
	a.body = "this is body"
	a.tags[0] = "ok"
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(200)
	json.NewEncoder(res).Encode(a)
}
