package main

import "net/http"

func main() {
	r := http.NewServeMux()
	r.HandleFunc("GET /{echo}", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(r.PathValue("echo")))
	})

	http.ListenAndServe(":8080", r)
}
