package main

import (
	"encoding/json"
	"net/http"

	"wsbrasil.com/simulate/kangu/kangu"
	"wsbrasil.com/simulate/kangu/pluggto"
	"wsbrasil.com/simulate/kangu/yampi"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		decoder := json.NewDecoder(r.Body)

		origem := r.Header.Get("origem")

		if origem == "pluggto" {
			paramsKangu := pluggto.Parse(decoder, r)
			body, _ := json.Marshal(paramsKangu.Params)
			responseBody := kangu.Request(body, paramsKangu.Token)
			result := pluggto.MakeResult(responseBody)
			w.Header().Set("Content-Type", "application/json")
			content, _ := json.Marshal(result)
			w.Write(content)
		} else {
			paramsKangu := yampi.Parse(decoder, r)
			body, _ := json.Marshal(paramsKangu.Params)
			responseBody := kangu.Request(body, paramsKangu.Token)
			result := yampi.MakeResult(responseBody)
			w.Header().Set("Content-Type", "application/json")
			content, _ := json.Marshal(result)
			w.Write(content)
		}
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
