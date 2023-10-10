package senders

import (
	"bytes"
	"fmt"
	"net/http"
)

var APIURL string

func SendData(value, name string, valueType string) (retCode int) {

	finalURL := fmt.Sprintf("%s/update/%s/%s/%s", APIURL, valueType, name, value)
	r := bytes.NewReader([]byte("test"))
	resp, err := http.Post(finalURL, "text/plain", r)
	//var retCode = -1
	retCode = 1
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resp)
		retCode = resp.StatusCode
	}

	defer resp.Body.Close()
	//return retCode
	return
}
