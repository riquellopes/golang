package main
import (
  "fmt"
  "log"
  "io/ioutil"
  "github.com/smallfish/simpleyaml"
)

func main(){
  example, err := ioutil.ReadFile("example.yml")

  if err != nil {
    log.Println("Error to load example.yml")
  }

  yaml, err := simpleyaml.NewYaml(example)

  if err != nil {
    log.Println("NewYaml error.")
  }

  fmt.Println(yaml.Get("name").String())
  fmt.Println(yaml.Get("last_name").String())
  fmt.Println(yaml.Get("age").Int())
}
