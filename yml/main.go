package main
import (
  "fmt"
  "log"
  "io/ioutil"
  "gopkg.in/yaml.v2"
)
// People -
type People struct {
  Name string `yaml:name"`
  LastName string `yaml:"last_name"`
  Age int `yaml:"age"`
}

func main(){

    example, err := ioutil.ReadFile("example.yml")

    if err != nil {
      log.Println("Error to load example.yml")
    }

    p := People{}
    err = yaml.Unmarshal(example, &p)

    if err != nil {
      log.Println("Unmarshal error.")
    }

    fmt.Println(p)
}
