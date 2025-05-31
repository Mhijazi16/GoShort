package main

import (
	"log"
	"net/http"
	"os"

	"github.com/google/uuid"
)

const filename string = "/home/ka1ser/github/GoShort/data.txt"

func createRedirection(w http.ResponseWriter, r *http.Request) {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal("⚠️ failed to open file")
	}
	defer file.Close()

	target := r.PathValue("target")
	data := uuid.New().String() + "=>" + target
	file.WriteString(data)
	w.WriteHeader(202)
	log.Println(target, " was registered in the service")
}

func main() {
	router := http.NewServeMux()
	router.HandleFunc("POST /create/{target}", createRedirection)
	log.Println("starting registeration service with port 8081")
	http.ListenAndServe(":8081", router)
}
