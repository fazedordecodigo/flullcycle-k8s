package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", Hello)
	http.HandleFunc("/config", ConfigMap)
	http.HandleFunc("/secret", Secret)
	http.ListenAndServe(":80", nil)
}

func Hello(w http.ResponseWriter, r *http.Request) {
	name := os.Getenv("NAME")
	age := os.Getenv("AGE")

	fmt.Fprintf(w, "Hello, %s! You are %s years old.", name, age)
}

func ConfigMap(w http.ResponseWriter, r *http.Request) {
	data, err := os.ReadFile("myfamily/family.txt")
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	fmt.Fprintf(w, "My Family:\n%s", string(data))
}

func Secret(w http.ResponseWriter, r *http.Request) {
	secret := os.Getenv("SECRET")
	fmt.Fprintf(w, "The secret is: %s", secret)
}
