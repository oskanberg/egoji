package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// {
//   input: "lorem ipsum"
// }

type request struct {
	Input string `json:"input"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var t request
	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	fmt.Println(t)
	encoder := json.NewEncoder(w)
	encoder.Encode(t)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
