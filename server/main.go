package main

import (
  "log"
  "net/http"
  "fmt"
  "github.com/pcarleton/hue/remote"
  //"github.com/pcarleton/hue/color"
)

var x = 0
var lights = [3]remote.LightState {
  remote.DefaultState,
  remote.DefaultState,
  remote.DefaultState,
}

func handler(w http.ResponseWriter, r *http.Request) {
  x += 1
  m := remote.GetStatusMsg()
  remote_lights := m.Lights
  fmt.Fprintf(w, "<style>.circle { text-align: center; padding: 2em; width: 7em; height: 7em; border-radius: 7em;}</style>")

  for _, light_msg := range remote_lights {
    name := light_msg.Name
    light_state := light_msg.State
    hex_color := light_state.GetRGB().Hex()
    fmt.Fprintf(w, "<div class='circle' style='background-color:%s'>%s  %+v</div>", hex_color, name, light_state.GetRGB())
  }
  light1 := remote_lights["1"].State


  fmt.Fprintf(w, "Hello %d!\n",x)
  fmt.Fprintf(w, "%+v", light1)
  fmt.Fprintf(w, "%+v", light1.GetRGB())
  log.Println("loggged")

}

func main() {
  http.HandleFunc("/", handler)
  http.ListenAndServe(":8080", nil )
}

