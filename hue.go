// Library functions for working with HUE
package hue

import  (
  "fmt"
  "net/http"
  "encoding/json"
  "io"
)

type MsgWrapper struct {
  msg StatusMessage
}

type StatusMessage struct {
  Config map[string] interface{}
  Lights map[string] LightMessage
  Groups map[string] interface{}
  Schedules map[string] interface{}
  Scenes  map[string] interface{}
}

type LightState struct {
  On bool
  Bri int
  Hue int
  Sat int
  Xy []float64
  Alert string
  Effect string
  Colormode string
  Reachable bool
}

type LightMessage struct {
  State LightState
  Name string
}

type LightsDict struct {
  dict map[string] interface{}
}

type Lights interface{}
type Groups interface{}
type Config interface{}
type Schedules interface{}

const meethue = "http://www.meethue.com/api/nupnp"
const developer_name = "newdeveloper"

func DoPut(url, data string) {
  req, err := http.NewRequest("PUT", url, data)
  if err != nil {
    fmt.Printf("Error: %v\n", err)
  }
  client := http.Client{}
  resp, err := client.Do(req)
  fmt.Printf("Response: %v\n", resp)

}

func GetStatusMsg() string {
  resp, err := http.Get("http://10.0.1.4/api/newdeveloper")
  if err != nil {
    fmt.Println("Error: %v", err)
  }
  fmt.Print("%v\n", resp)
  dec := json.NewDecoder(resp.Body)
  var m StatusMessage
  err = dec.Decode(&m)
  if err != nil {
    fmt.Printf("Error: %v\n", err)
  }
  fmt.Printf("%v\n", m.Lights["1"])
  //resp, err = DoPut("http://10.0.1.4/api/lights/1/state", "{ \"on\":false }")
  if err != nil {
    fmt.Printf("Error: %v\n", err)
  }
  fmt.Println(resp)
  return "success"
}

func ConvertJsonToMap(b io.Reader) (m map[string] interface{}, ok bool) {
  dec := json.NewDecoder(b)
  var f interface{}
  dec.Decode(&f)
  //Its possible we got an array rather than a simple map
  //In this case, take the first element
  switch v := f.(type) {
    case []interface{}:
      m, ok = v[0].(map[string]interface{})
    default:
      m, ok = v.(map[string] interface{})
  }
  return
}

func GetHueIp() string {
  resp, err := http.Get(meethue)
  if (err != nil) {
    fmt.Println("Error %v", err)
  }
  ms, ok := ConvertJsonToMap(resp.Body)
  if (!ok) {
    fmt.Println("Could not process the response")
  }

  ip := ms["internalipaddress"].(string)
  return ip
}

//func IssueGetRequest(request_uri string) map[string] interface{} {
//}

