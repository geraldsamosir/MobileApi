package Article

import (
	"github.com/gorilla/mux"
)

type Router struct {
}

func (routes *Router) AddSignHandler(r *mux.Router) {
	var articleService ArticleService
	r.HandleFunc("/article", articleService.Get).Methods("GET")
}
