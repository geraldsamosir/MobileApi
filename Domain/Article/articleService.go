package Article

import (
	"encoding/json"
	"net/http"
)

type ArticleService struct {
	Title string
	Body  string
	Tags  []string
}

func (a *ArticleService) Get(res http.ResponseWriter, req *http.Request) {
	a.Title = "hello"
	a.Body = "this is body"
	var tags []string
	a.Tags = append(tags, "ok")
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(200)
	json.NewEncoder(res).Encode(a)
}
