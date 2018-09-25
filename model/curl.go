package model

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

type Responses struct {
	City    string  `json:"city"`
	Country string  `json:"country"`
	Lat     float64 `json:"lat"`
	Lon     float64 `json:"lon"`
}

func GetDataCurl(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	ip := vars["ip"]

	url := "http://ip-api.com/json/"
	url += ip

	req, _ := http.NewRequest("GET", url, nil)

	res, _ := http.DefaultClient.Do(req)

	//var resp Responses
	body, _ := ioutil.ReadAll(res.Body)
	m := []byte(body)
	rsp := bytes.NewReader(m)
	val := &Responses{}
	json.NewDecoder(rsp).Decode(val)
	json.NewEncoder(w).Encode(val)
}
