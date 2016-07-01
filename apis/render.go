package apis

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type ErrorNode struct {
	Status string `json:"status,omitempty"`
	Source string `json:"source,omitempty"`
	Title  string `json:"title,omitempty"`
	Detail string `json:"detail,omitempty"`
}

type ErrorPayload struct {
	Errors []*ErrorNode `json:"errors"`
}

func render(rw http.ResponseWriter, data []byte) {
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(data)
}

func renderError(rw http.ResponseWriter, err error, errorCode int) {
	jsonErr := ErrorNode{Status: strconv.Itoa(errorCode), Title: err.Error(), Detail: err.Error()}
	payload := ErrorPayload{Errors: []*ErrorNode{&jsonErr}}

	b, err := json.Marshal(payload)
	if err != nil {
		log.Println("error:", err)
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(errorCode)
	rw.Write(b)
}
