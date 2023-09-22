package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

const (
	PORT = 9001
)

type NavLink struct {
	Name string
	Link string
}

// type TemplateSelector struct {
// 	Name string
// 	Opt string
// }

type Page struct {
	Nav            []NavLink
	TemplateSelect string
	Dynamic        bool
}

func Navbar() []NavLink {
	return []NavLink{
		{
			Name: "/",
			Link: "/",
		},
		{
			Name: "Finding",
			Link: "/findings",
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
			Link: "/tool-defects",
		},
		{
			Name: "Wiki",
			Link: "/wiki",
		},
		{
			Name: "Settings",
			Link: "/settings",
		},
	}
}

func rootHandle(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseGlob("templates/*.htmx"))

	page := Page{
		Nav:            Navbar(),
		TemplateSelect: "root",
	}

	tpl.ExecuteTemplate(w, "main", page)

}

func findingsHandle(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseGlob("templates/*.htmx"))

	page := Page{
		Nav:            Navbar(),
		TemplateSelect: "findings",
	}

	tpl.ExecuteTemplate(w, "main", page)
}

func filtersHandle(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseGlob("templates/*.htmx"))
	page := Page{
		Nav:            Navbar(),
		TemplateSelect: "filters",
	}

	tpl.ExecuteTemplate(w, "main", page)
}

func casesHandle(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseGlob("templates/*.htmx"))
	page := Page{
		Nav:            Navbar(),
		TemplateSelect: "cases",
	}

	tpl.ExecuteTemplate(w, "main", page)
}

func toolDefectsHandle(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseGlob("templates/*.htmx"))
	page := Page{
		Nav:            Navbar(),
		TemplateSelect: "tooldefect",
	}

	tpl.ExecuteTemplate(w, "main", page)
}

func wikiHandle(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseGlob("templates/*.htmx"))
	page := Page{
		Nav:            Navbar(),
		TemplateSelect: "wiki",
	}

	tpl.ExecuteTemplate(w, "main", page)
}

func settingsHandle(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseGlob("templates/*.htmx"))

	page := Page{
		Nav:            Navbar(),
		TemplateSelect: "settings",
	}
	
	tpl.ExecuteTemplate(w, "main", page)
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
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))

	// URL Routes
	mux.HandleFunc("/", rootHandle)
	mux.HandleFunc("/findings", findingsHandle)
	mux.HandleFunc("/filters", filtersHandle)
	mux.HandleFunc("/cases", casesHandle)
	mux.HandleFunc("/tool-defects", toolDefectsHandle)
	mux.HandleFunc("/wiki", wikiHandle)

	mux.HandleFunc("/settings", settingsHandle)
	mux.HandleFunc("/graphql", graphqlHandle)

	fmt.Println("server at http://localhost:" + port)
	if http.ListenAndServe(":"+port, mux) != nil {
		log.Fatal("could not listen and serve")
	}

}
