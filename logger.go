package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func serveMain(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "mrlabel.html")
}

func readLines(fname string, mustExist bool) ([]string, error) {
	data, err := ioutil.ReadFile(fname)
	if err != nil && (!os.IsNotExist(err) || mustExist) {
		return nil, err
	}
	if err != nil {
		// File does not exist, empty list.
		return nil, nil
	}
	lines := strings.Split(string(data), "\n")
	var res []string
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line != "" {
			res = append(res, line)
		}
	}
	return res, nil
}

func serveImages(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")
	images, err := readLines("images.txt", true)
	if err != nil {
		log.Print(err)
		http.Error(w, "Could not read images.txt", 500)
		return
	}
	labels, err := readLines("labels.txt", false)
	if err != nil {
		log.Print(err)
		http.Error(w, "Could not read labels.txt", 500)
		return
	}
	was := make(map[string]bool)
	for _, label := range labels {
		key := strings.Split(label, " ")[0]
		was[key] = true
	}
	fmt.Fprintln(w, "images_str = `")
	for _, image := range images {
		if !was[image] {
			fmt.Fprintln(w, image)
		}
	}
	fmt.Fprintln(w, "`")
}

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
	http.HandleFunc("/", serveMain)
	http.HandleFunc("/images.txt", serveImages)
	http.HandleFunc("/log", logAnswer)
	log.Fatal(http.ListenAndServe(":10000", nil))
}
