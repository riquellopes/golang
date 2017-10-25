package main

import (
    "net/http"
    "log"
    "encoding/json"
    "io/ioutil"
)

// ResquestResultDollar -
type ResquestResultDollar struct {
    Status bool `json:"status"`
    Valores struct {
        USD struct {
            Fonte string `json:"fonte"`
            Nome string `json:"nome"`
            UltimaConsulta int `json:"ultima_consulta"`
            Valor float64 `json:valor`
        } `json:"USD"`
    } `json:"valores"`
}

// Parse -
func Parse(response string) *ResquestResultDollar {
    result := new(ResquestResultDollar)
    err := json.Unmarshal([]byte(response), result)

    if err != nil {
        log.Println(err)
    }

    log.Println(result)
    return result
}

func request() (string, error){
    log.Println("Iniciando request")
    response, err := http.Get(
            "http://api.promasters.net.br/cotacao/v1/valores?moedas=USD&alt=json")

    if err != nil {
        return "", err
    }

    // Close request
    defer func () {
        err := response.Body.Close()

        if err != nil {
            log.Println("Error to close request.")
        }
    }()

    contents, err := ioutil.ReadAll(response.Body)

    if err != nil {
        return "", err
    }

    return string(contents), nil
}


func main() {
    if response, err := request(); err == nil {
        parse := Parse(response)

        log.Println(parse.Valores.USD.Valor)
    }

    log.Println("Show")
}
