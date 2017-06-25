package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	//http.HandleFunc("/", handler)
	http.HandleFunc("/feed/", feed)
	http.Handle("/prometheus/", promhttp.Handler())
	println("Starting")
	http.ListenAndServe(":8080", nil)

}

type Feedstruct struct {
	Id int64
	Caption string
	Url string
}

func feed(w http.ResponseWriter, r *http.Request) {
	if r.Method =="GET" {

		//
		f := make([]Feedstruct, 0)
		//println(len(r.URL.Path))
		if len(r.URL.Path) <= 6 {
			f1 := Feedstruct{1, "my cap", "xxx"}
			f2 := Feedstruct{2, "what's up", "yyy"}
			f = append(f, f1)
			f = append(f, f2)
		} else {
			f1 := Feedstruct{1, "your cap", "xxx"}
			f2 := Feedstruct{2, "hey "+r.URL.Path[6:], "yyy"}
			f = append(f, f1)
			f = append(f, f2)
		}


		//f := Message{"Welcome to the SandovalEffect API, build v0.0.001.992, 6/22/2015 0340 UTC."}
		b, err := json.Marshal(f)

		if err != nil {
			panic(err)
		}
		w.Write(b)
	} else if r.Method == "PUT" {
		fmt.Fprintf(w, "Success")
	}
}
