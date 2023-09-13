package main

import (
	"fmt"
	"net/http"
	"strconv"
	"log"
)

const (
	PORT = 9001
)

func rootHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "yo")
}

func findingsHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w,"findings")
}

func filtersHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w,"filters")
}

func casesHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w,"cases")
}

func toolDefectsHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w,"tool-defects")
}

func wikiHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w,"wiki")
}


func main() {
	port := strconv.Itoa(PORT)

	mux := http.NewServeMux()
	mux.HandleFunc("/",rootHandle)
	mux.HandleFunc("/findings",findingsHandle)
	mux.HandleFunc("/filters",filtersHandle)
	mux.HandleFunc("/cases",casesHandle)
	mux.HandleFunc("/tool-defects",toolDefectsHandle)
	mux.HandleFunc("/wiki",wikiHandle)
	

	fmt.Println("server at http://localhost:"+port)
	if http.ListenAndServe(":"+port,mux) != nil {
		log.Fatal("could not listen and serve")
	}

}
