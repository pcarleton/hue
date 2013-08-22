// Library functions for working with HUE
package remote

import  (
  "fmt"
  "strings"
  "net/http"
  "encoding/json"
  "io"
)

const meethue = "http://www.meethue.com/api/nupnp"
const developer_name = "newdeveloper"
const hue_ip = "http://10.0.1.3"

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

func DoPut(url, data string) {
  req, err := http.NewRequest("PUT", url, strings.NewReader(data))
  if err != nil {
    fmt.Printf("Error: %v\n", err)
  }
  client := http.Client{}
  resp, err := client.Do(req)
  fmt.Printf("Response: %v\n", resp)
}

func GetStatusMsg() string {
  resp, err := http.Get(hue_ip + "/api/" + developer_name)
  if err != nil {
    fmt.Printf("Error: %v\n", err)
  }
  fmt.Printf("%v\n", resp)
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

func GetStatus() (string, error){
  resp, err := http.Get(hue_ip + "newdeveloper/api")
  if err != nil {
    return "", err
  }
  ms, ok := ConvertJsonToMap(resp.Body)
  if (!ok) {
    fmt.Println("Could not process the response")
  }
 return fmt.Sprint("%v", ms), nil
}

func GetHueIp() string {
  resp, err := http.Get(meethue)
  if (err != nil) {
    fmt.Printf("Error %v\n", err)
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

