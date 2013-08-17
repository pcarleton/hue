package main

import (
  "fmt"
  remote "github.com/pcarleton/hue/remote"
)

func main() {
  fmt.Println("HEllo!")
  fmt.Println(remote.GetStatus())
}
