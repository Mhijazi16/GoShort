package main

import (
	"fmt"
	"net/http"
)

func redirect(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	fmt.Println(id)
	http.Redirect(w, r, "https://coldrelay.com", http.StatusMovedPermanently)
}

func main() {
	router := http.NewServeMux()
	router.HandleFunc("GET /{id}", redirect)
	fmt.Println("server listening on 8080")
	http.ListenAndServe(":8080", router)
}
