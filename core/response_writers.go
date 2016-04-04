package core

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

func WriteResponse(w http.ResponseWriter, raw interface{}) {
	response, err := json.Marshal(raw)
	if err != nil {
		WriteError(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if _, err = w.Write(response); err != nil {
		WriteError(w, err)
	}
}

func WriteErrorBadRequest(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusBadRequest)
	WriteError(w, err)
}

func WriteErrorInternal(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	WriteError(w, err)
}

func WriteError(w http.ResponseWriter, err error) {
	log.Println("error occured", err)
	var errStruct struct {
		Msg string `json:"msg"`
	}
	errStruct.Msg = strings.Replace(err.Error(), `"`, "'", -1)
	response, err := json.Marshal(errStruct)
	if err != nil {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}
