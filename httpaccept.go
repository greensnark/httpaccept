package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
)

var address = flag.String("addr", ":8080", "bind address")

func main() {
	flag.Parse()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(os.Stderr, "\n"+r.Method, r.URL.Path)
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

	listener, err := net.Listen("tcp", *address)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to bind to", *address, err)
		return
	}

	fmt.Fprintln(os.Stderr, "Listening on", "http://"+listener.Addr().String())
	http.Serve(listener, nil)
}
