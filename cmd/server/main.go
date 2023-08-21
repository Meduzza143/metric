package main

import (
	"net/http"

	handlers "github.com/Meduzza143/metric/cmd/server/internal"
)

func main() {

	mux := http.NewServeMux()
	//http://<АДРЕС_СЕРВЕРА>/update/<ТИП_МЕТРИКИ>/<ИМЯ_МЕТРИКИ>/<ЗНАЧЕНИЕ_МЕТРИКИ>
	mux.HandleFunc(`/update/`, handlers.UpdateHandle)

	err := http.ListenAndServe(`localhost:8080`, mux)
	if err != nil {
		panic(err)
	}
}
