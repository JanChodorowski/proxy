package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		link := r.URL.Query().Get("link")

		if link == "" {
			http.Error(w, "no link to asset provided", http.StatusBadRequest)
		}

		resp, err := http.Get(link)

		if err != nil {
			fmt.Fprint(w, "Error")
		}

		bytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)

		}

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/octet-stream")
		w.WriteHeader(http.StatusOK)
		w.Write(bytes)

	})

	port := os.Getenv("PORT")
	if port == "" {
		panic("Port not specified")
	}

	log.Fatal(http.ListenAndServe(":"+port, nil))

}
