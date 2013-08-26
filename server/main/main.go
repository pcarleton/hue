package main

import (
  "log"
  "net/http"
  "fmt"
  "github.com/pcarleton/hue/remote"
)

var x = 0
var l = remote.LightState{}

func handler(w http.ResponseWriter, r *http.Request) {
  x += 1
  fmt.Fprintf(w, "Hello %d!",x)
  log.Println("loggged")

}

func main() {
  http.HandleFunc("/", handler)
  http.ListenAndServe(":8080", nil )
}

