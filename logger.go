package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func logAnswer(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", req.Header.Get("Origin"))
	var img = req.FormValue("image")
	var label = req.FormValue("label")
	log.Printf("img: %q, label: %q", img, label)
	f, err := os.OpenFile("labels.txt", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		log.Print(err)
		http.Error(w, "Could not open the file with labels", 500)
		return
	}
	fmt.Fprintf(f, "%s %s\n", img, label)
	if err = f.Close(); err != nil {
		log.Print(err)
		http.Error(w, "Could not write to the labels file", 500)
		return
	}
}

func main() {
	http.HandleFunc("/log", logAnswer)
	log.Fatal(http.ListenAndServe(":10000", nil))
}
