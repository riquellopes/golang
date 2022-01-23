package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type ResquestResultDollar struct {
	Usdbrl struct {
		Code       string `json:"code"`
		Codein     string `json:"codein"`
		Name       string `json:"name"`
		High       string `json:"high"`
		Low        string `json:"low"`
		VarBid     string `json:"varBid"`
		PctChange  string `json:"pctChange"`
		Bid        string `json:"bid"`
		Ask        string `json:"ask"`
		Timestamp  string `json:"timestamp"`
		CreateDate string `json:"create_date"`
	} `json:"USDBRL"`
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

func request() (string, error) {
	log.Println("Iniciando request")
	response, err := http.Get(
		"https://economia.awesomeapi.com.br/last/USD-BRL")

	if err != nil {
		return "", err
	}

	log.Println("Recuperado request")

	// Close request
	defer func() {
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

		log.Println(parse.Usdbrl.Bid)
	}

	log.Println("Show")
}
