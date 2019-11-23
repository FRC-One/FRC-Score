package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
)

func returnHomePage(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("html/home.html","html/pageSections/pageStart.html","html/pageSections/pageEnd.html")

	if err == nil{
		w.WriteHeader(200)
		err = t.ExecuteTemplate(w, "home.html", nil)
		if err != nil{
			panic(err)
		}


	} else{
		fmt.Println(err)
	}

}

func returnTrackPage(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("html/track.html", "html/pageSections/pageStart.html","html/pageSections/pageEnd.html")

	if err == nil{
		w.WriteHeader(200)
		err = t.ExecuteTemplate(w, "track.html", nil)
		if err != nil{
			panic(err)
		}


	} else{
		fmt.Println(err)
	}

}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", returnHomePage)
	r.HandleFunc("/track", returnTrackPage)
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))


	if err := http.ListenAndServe(":8080", r); err != nil {
		panic(err)
	}
	
}
