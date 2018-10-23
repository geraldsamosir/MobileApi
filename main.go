package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/geraldsamosir/MobileApi/Domain/Article"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
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
	viper.AddConfigPath("configs/.")
	viper.SetConfigType("json")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
		fmt.Println("Config file not found...")
	}
	var m Message
	var article Article.Router
	r := mux.NewRouter()
	r.HandleFunc("/", m.Test).Methods("GET", "POST")

	article.AddSignHandler(r)
	fmt.Println("Your api run in port " + viper.GetString("port"))
	http.ListenAndServe("0.0.0.0:"+viper.GetString("port")+"", r)
}
