package main
import (
  "fmt"
  "log"
  "io/ioutil"
  "gopkg.in/yaml.v2"
)


// Yaml -
type Yaml struct {
  data interface{}
}

func main(){

    example, err := ioutil.ReadFile("example.yml")

    if err != nil {
      log.Println("Error to load example.yml")
    }

    var data interface{}
    err = yaml.Unmarshal(example, &data)

    if err != nil {
      log.Println("Unmarshal error.")
    }

    y := Yaml{data}
    fmt.Println(y)
}
