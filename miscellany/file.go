package main

import "io/ioutil"

func main(){
    texto := []byte("Conteudo")
    err := ioutil.WriteFile("/tmp/content.txt", texto, 0644)

    if err != nil {
        panic(err)
    }
}
