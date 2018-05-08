package main

// https://github.com/go-br/estudos/tree/master/wikipedia

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"
)

func main() {

	resp, err := http.Get("https://pt.wikipedia.org/w/api.php?action=opensearch&format=json&search=Go_(linguagem_de_programação)")

	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
	}

	var m []interface{}

	err = json.Unmarshal(body, &m)

	if err != nil {
		fmt.Println(err)
		return
	}

	r := make(map[string]int)

	walker(m, r)

	for k, v := range r {
		fmt.Println(k, v)
	}
}

func walker(m interface{}, r map[string]int) {

	for _, v := range m.([]interface{}) {
		switch reflect.TypeOf(v).Kind() {
		case reflect.String:
			a := strings.Split(v.(string), " ")

			for _, s := range a {
				aux := r[s]
				aux++
				r[s] = aux
			}
		case reflect.Slice:
			walker(v, r)
		default:
			fmt.Println(v, reflect.TypeOf(v))
		}
	}
}
