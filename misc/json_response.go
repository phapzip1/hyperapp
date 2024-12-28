package misc

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type errResp struct {
	Error string `json:"error"`
}

func ResponseWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)

	data, err := json.Marshal(payload)
	if err != nil {
		ResponseWithError(w, http.StatusInternalServerError, err.Error())
	}

	w.Write(data)
}

func ResponseWithError(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		fmt.Println("internal server error: ", msg)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)

 	data, _ := json.Marshal(errResp{
		Error: msg,
	})

	w.Write(data)
}
