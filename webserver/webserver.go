package main

import (
	"encoding/json"
	"net/http"

	"github.com/oskanberg/egoji"
)

type request struct {
	Input string `json:"input"`
}

type response struct {
	Output string `json:"output"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var t request
	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	translater := egoji.NewSimpleTranslate()
	output, _ := translater.Translate(t.Input)

	encoder := json.NewEncoder(w)
	encoder.Encode(response{Output: output})
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
