package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Message struct {
	Mymessage string
}

func (m *Message) Test(res http.ResponseWriter, req *http.Request) {
	m.Mymessage = "message ok from modules"
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(200)
	json.NewEncoder(res).Encode(m)
}

func main() {
	fmt.Println("hello this is from  package golang using go mod ")
	var m Message
	r := mux.NewRouter()

	r.HandleFunc("/", m.Test)

	http.ListenAndServe("0.0.0.0:5000", r)
}
