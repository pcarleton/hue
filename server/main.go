package main

import (
  "log"
  "net/http"
  "fmt"
  "github.com/pcarleton/hue/remote"
)

var x = 0
var lights = [3]remote.LightState {
  remote.DefaultState,
  remote.DefaultState,
  remote.DefaultState,
}

func handler(w http.ResponseWriter, r *http.Request) {
  x += 1
  fmt.Fprintf(w, "Hello %d!\n",x)
  fmt.Fprintf(w, "%+v", lights)
  log.Println("loggged")

}

func main() {
  http.HandleFunc("/", handler)
  http.ListenAndServe(":8080", nil )
}

