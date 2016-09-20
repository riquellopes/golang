package main

import (
    "html/template"
    "io"
    "log"
    "net/http"
)

func home(res http.ResponseWriter, req *http.Request){
    io.WriteString(res, "Home")
}

func hello(res http.ResponseWriter, req *http.Request){
    tpl, err := template.ParseFiles("index.html")
    if err != nil {
        log.Fatalln(err)
    }

    tpl.Execute(res, nil)
}

func main(){
    http.HandleFunc("/", home)
    http.HandleFunc("/hello.html", hello)
    http.ListenAndServe(":5000", nil)
}
