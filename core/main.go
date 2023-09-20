package main

import (
	"fmt"
	"net/http"
	"strconv"
	"log"
	"text/template"
)

const (
	PORT = 9001
)


type NavLink struct {
	Name string
	Link string
}

type Page struct {
	Nav []NavLink
}

func Navbar() []NavLink {
	return []NavLink{
		{
			Name: "Home",
			Link: "/",
		},
		{
			Name: "Finding",
			Link: "/finding",
		},
		{
			Name: "Filters",
			Link: "/filters",
		},
		{
			Name: "Cases",
			Link: "/cases",
		},
		{
			Name: "Tool Defects",
			Link: "/tool_defects",
		},
		{
			Name: "Wiki",
			Link: "/wiki",
		},
	}
}


func rootHandle(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseGlob("templates/*.html"))

	page := Page{
		Nav: Navbar(),
	}
	
	tpl.ExecuteTemplate(w, "main",page)	

}

func findingsHandle(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseGlob("templates/*.html"))

	page := Page{
		Nav: Navbar(),
	}
	
	tpl.ExecuteTemplate(w, "main",page)	
}

func filtersHandle(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseGlob("templates/*.html"))
	page := Page{
		Nav: Navbar(),
	}
	
	tpl.ExecuteTemplate(w, "main",page)	
}

func casesHandle(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseGlob("templates/*.html"))
	page := Page{
		Nav: Navbar(),
	}
	
	tpl.ExecuteTemplate(w, "main",page)	
}

func toolDefectsHandle(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseGlob("templates/*.html"))
	page := Page{
		Nav: Navbar(),
	}
	
	tpl.ExecuteTemplate(w, "main",page)	
}

func wikiHandle(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseGlob("templates/*.html"))
	page := Page{
		Nav: Navbar(),
	}
	
	tpl.ExecuteTemplate(w, "main",page)	
}

func graphqlHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "put graphql api here")
}

func main() {
	port := strconv.Itoa(PORT)

	mux := http.NewServeMux()

	// add files in ./assets
	fs := http.FileServer(http.Dir("assets"))

	// route files in ./assets to /assets/ in the html
	mux.Handle("/assets/",http.StripPrefix("/assets/",fs))

	// URL Routes
	mux.HandleFunc("/",rootHandle)
	mux.HandleFunc("/findings",findingsHandle)
	mux.HandleFunc("/filters",filtersHandle)
	mux.HandleFunc("/cases",casesHandle)
	mux.HandleFunc("/tool-defects",toolDefectsHandle)
	mux.HandleFunc("/wiki",wikiHandle)
	mux.HandleFunc("/graphql",graphqlHandle)
	
	fmt.Println("server at http://localhost:"+port)
	if http.ListenAndServe(":"+port,mux) != nil {
		log.Fatal("could not listen and serve")
	}

}
