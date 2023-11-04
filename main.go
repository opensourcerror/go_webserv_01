package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	if r.Method == http.MethodGet {
		w.Write([]byte("wlcm home \n"))
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed) //405
	}
}

// curl -i -X POST http://localhost:4000/snippet/create
func snippetCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", "POST")
		http.Error(w, "meth not allowed", 405)
		return
	}
	w.Write([]byte("create new snippet \n"))
}

// http://localhost:4000/snippet/view?id=123
func snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "showing snippet with id:  %d", id)

}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/create", snippetCreate)
	mux.HandleFunc("/snippet/view", snippetView)

	port := ":4000"
	log.Printf("srv up on %s", port)
	err := http.ListenAndServe(port, mux)
	log.Fatal(err)
}
