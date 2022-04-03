package main

import (
  "fmt"

  "api/config"
)

func init() {
  config.Load()
}

func main() {
  fmt.Println("hello")
  fmt.Println(config.Global.RESTPort)
}
