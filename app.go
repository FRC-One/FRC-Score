package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

)

func returnGlobalHTML(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadFile("./html/global.html")
	if err == nil{
		w.WriteHeader(200)
		fmt.Fprintf(w, string(body))


	}

}

func main() {
	http.HandleFunc("/", returnGlobalHTML)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
	
}
