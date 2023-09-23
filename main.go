package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"wsbrasil.com/simulate/kangu/kangu"
	"wsbrasil.com/simulate/kangu/pluggto"
	"wsbrasil.com/simulate/kangu/yampi"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		debug := r.Header.Get("Debug")
		if debug == "true" {
			var bodyBytes []byte
			var err error

			if r.Body != nil {
				bodyBytes, err = ioutil.ReadAll(r.Body)
				if err != nil {
					fmt.Printf("Body reading error: %v", err)
					return
				}
				// defer r.Body.Close()
			}

			fmt.Printf("Headers: %+v\n", r.Header)

			if len(bodyBytes) > 0 {
				var prettyJSON bytes.Buffer
				if err = json.Indent(&prettyJSON, bodyBytes, "", "\t"); err != nil {
					fmt.Printf("JSON parse error: %v", err)
					return
				}
				fmt.Println(string(prettyJSON.Bytes()))
			} else {
				fmt.Printf("Body: No Body Supplied\n")
			}
		}

		decoder := json.NewDecoder(r.Body)

		origem := r.Header.Get("Origem")

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
