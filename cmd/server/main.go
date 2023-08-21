package main

import (
	"fmt"
	"net/http"

	handlers "github.com/Meduzza143/metric/cmd/server/internal"
	"github.com/gorilla/mux"
)

func main() {

	//http://<АДРЕС_СЕРВЕРА>/update/<ТИП_МЕТРИКИ>/<ИМЯ_МЕТРИКИ>/<ЗНАЧЕНИЕ_МЕТРИКИ>
	r := mux.NewRouter()
	r.HandleFunc(`/update/{type}/{name}/{value}`, handlers.UpdateHandle).Methods("POST")
	r.HandleFunc(`/value/{type}/{name}`, handlers.GetMetric).Methods("GET")
	r.HandleFunc(`/`, handlers.GetAll).Methods("GET")

	// mux.HandleFunc(`/update/`, handlers.UpdateHandle).Methods("POST")
	// mux.HandleFunc(`/value/`, handlers.GetHandle)
	// mux.HandleFunc(`/value/`, handlers.GetHandle)

	fmt.Println("starting server... ")

	err := http.ListenAndServe(`localhost:8080`, r)
	if err != nil {
		panic(err)
	}

}
