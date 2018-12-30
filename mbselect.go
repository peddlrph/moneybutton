// websockets.go
package main

import (
	"html/template"
	"net/http"
)

type MBDetails struct {
	DataAmount string
}

func main() {

	tmpl := template.Must(template.ParseFiles("showmb.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "mbselect.html")
	})

	http.HandleFunc("/showmb", func(w http.ResponseWriter, r *http.Request) {
		var mbdetails MBDetails
		mbdetails.DataAmount = r.FormValue("mbvalue")
		tmpl.Execute(w, mbdetails)
	})

	http.ListenAndServe(":80", nil)
}
