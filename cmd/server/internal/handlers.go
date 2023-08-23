package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// type value interface {
// 	GetValue() interface{}
// }
// type int64Var struct{ value int64 }
// type float64Var struct{ value float64 }
// type stringVar struct{ value string }

// func (v int64Var) GetValue() interface{} {
// 	return v.value
// }
// func (v float64Var) GetValue() interface{} {
// 	return v.value
// }
// func (v stringVar) GetValue() interface{} {
// 	return v.value
// }

type MemStruct struct {
	metricType string
	value      interface{}
	//gaugeValue   float64
	//counterValue int64
}

var MemStorage = make(map[string]MemStruct)

// //http://<АДРЕС_СЕРВЕРА>/update/<ТИП_МЕТРИКИ>/<ИМЯ_МЕТРИКИ>/<ЗНАЧЕНИЕ_МЕТРИКИ>
func UpdateHandle(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("content-type", "text/plain")
	vars := mux.Vars(req)

	if vars["name"] == "" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	switch vars["type"] {
	case "gauge":
		currValue, err := strconv.ParseFloat(vars["value"], 64)
		if err == nil {
			MemStorage[vars["name"]] = MemStruct{vars["type"], currValue}
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	case "counter":
		currValue, err := strconv.ParseInt(vars["value"], 0, 64)
		if err == nil {
			thisValue := MemStorage[vars["name"]].value
			if thisValue == nil { //new value
				MemStorage[vars["name"]] = MemStruct{vars["type"], currValue}
			} else {
				switch i := thisValue.(type) {
				case int64:
					MemStorage[vars["name"]] = MemStruct{vars["type"], currValue + i}
				default:
					w.WriteHeader(http.StatusBadRequest) //wrong counter type
					return
				}
			}
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	default:
		w.WriteHeader(http.StatusBadRequest)
	}
}

func GetMetric(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("content-type", "text/plain")
	vars := mux.Vars(req)
	if val, ok := MemStorage[vars["name"]]; ok {
		if val.metricType == vars["type"] {
			switch val.metricType {
			case "gauge":
				w.Write([]byte(fmt.Sprint(val.value)))
			case "counter":
				w.Write([]byte(fmt.Sprint(val.value)))
			default:
				w.WriteHeader(http.StatusNotFound)
			}
		}
	} else {
		w.WriteHeader(http.StatusNotFound)
	}

}

func GetAll(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("content-type", "text/plain")
	body := "Current values: \n"
	for k, v := range MemStorage {
		switch v.metricType {
		case "gauge":
			body += fmt.Sprintf("%v = %v \n", k, v.value)
		case "counter":
			body += fmt.Sprintf("%v = %v \n", k, v.value)
		}
	}
	w.Write([]byte(body))
}
