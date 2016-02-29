package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

var address = ":8080"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(os.Stderr, "\n", r.Method, r.URL.Path)
		for header := range r.Header {
			value := r.Header.Get(header)
			fmt.Fprintln(os.Stderr, ">", header, "=", value)
		}
		if r.Body != nil {
			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				fmt.Fprintln(os.Stderr, "Error reading request body:", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			fmt.Println(string(body))
		}
		w.WriteHeader(http.StatusNoContent)
	})

	fmt.Fprintln(os.Stderr, "Listening on", address)
	http.ListenAndServe(address, nil)
}
