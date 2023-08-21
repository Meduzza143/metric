package main

import (
	"flag"
	"fmt"
	"net/http"

	handlers "github.com/Meduzza143/metric/cmd/server/internal"
	"github.com/gorilla/mux"
)

type NetAddress struct {
	Host string
	Port int
}

func main() {

	ipAddr := flag.String("a", "localhost:8080", "address:port")
	flag.Parse()
	//http://<АДРЕС_СЕРВЕРА>/update/<ТИП_МЕТРИКИ>/<ИМЯ_МЕТРИКИ>/<ЗНАЧЕНИЕ_МЕТРИКИ>
	r := mux.NewRouter()
	r.HandleFunc(`/update/{type}/{name}/{value}`, handlers.UpdateHandle).Methods("POST")
	r.HandleFunc(`/value/{type}/{name}`, handlers.GetMetric).Methods("GET")
	r.HandleFunc(`/`, handlers.GetAll).Methods("GET")

	fmt.Printf("starting server... at %v \n", *ipAddr)

	err := http.ListenAndServe(*ipAddr, r)
	if err != nil {
		panic(err)
	}

}
