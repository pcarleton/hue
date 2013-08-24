// Library functions for working with HUE
package remote

import  (
  "time"
  "bytes"
  "fmt"
  "net/http"
  "encoding/json"
  "io"
)

const meethue = "http://www.meethue.com/api/nupnp"
const developer_name = "newdeveloper"
const hue_ip = "10.0.1.2"

type StatusMessage struct {
  Config map[string] interface{}
  Lights map[string] LightMessage
  Groups map[string] interface{}
  Schedules map[string] interface{}
  Scenes  map[string] interface{}
}

type LightState struct {
  On bool `json:"on"`
  Bri int `json:"bri,omitempty"`
  Hue int`json:"hue,omitempty"`
  Sat int `json:"sat,omitempty"`
  Xy []float64 `json:"xy,omitempty"`
  Alert string `json:"alert,omitempty"`
  Effect string `json:"effect,omitempty"`
  Colormode string `json:"colormode,omitempty"`
  Reachable bool `json:"reachable,omitempty"`
}

type LightAction struct {
  LightState
  Transition int `json:"transitiontime,omitempty"`
}

var DefaultAction = LightAction {
  LightState: LightState{On:true},
  Transition: 1,
}

var DefaultState = LightState{
  On: true,
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

func buildLightsUrl(lightnum int, postfix string) string {
  return fmt.Sprintf("http://%s/api/%s/lights/%d%s", hue_ip, developer_name, lightnum, postfix)
}

func groupsUrl(groupnum int, postfix string) string {
  return fmt.Sprintf("http://%s/api/%s/groups/%d%s", hue_ip, developer_name, groupnum, postfix)
}

func TurnLightOn() {
  var onmsg = DefaultState
  onmsg.On = true
   b, err := json.Marshal(onmsg)
   if err != nil {
     fmt.Println("error:", err)
   }
   DoPut(buildLightsUrl(1, "/state"), bytes.NewReader(b))
}

type Action func(val int)

func DoStep(min, max, step int, action Action) {
  for i := min; i < max; i = i + step {
    action(i)
  }
}

type SetProperty func(val int) LightAction


func GroupSet(setter SetProperty) Action {
  c := make(chan LightAction)
  var action =  func(val int) {
    fmt.Println(val)
    c <- setter(val)
  }
  go func() { 
    for {
     s := <-c
     SetGroupState(s)
     Delay(5000)
    }
  }()
  return action
}

func SetBri(val int) LightState {
  s := DefaultState
  s.Bri = val
  return s
}

func SetSat(val int)  LightState{
   s := DefaultState
   s.Sat = val
   return s
}

func SetHue(val int) LightAction {
   s := DefaultAction
   s.Hue = val
   return s
}

func Delay(millis int64) {
  <-time.After(time.Duration(millis) * time.Millisecond)
}


func SetGroupState(s interface{}) {
   b, err := json.Marshal(s)
   if err != nil {
     fmt.Println("error:", err)
   }
   DoPut(groupsUrl(0, "/action"), bytes.NewReader(b))
}


func SetLightState(lightnum int, s interface{}) {
   b, err := json.Marshal(s)
   if err != nil {
     fmt.Println("error:", err)
   }
   DoPut(buildLightsUrl(lightnum, "/state"), bytes.NewReader(b))
}

func DoPut(url string, data io.Reader) {
  req, err := http.NewRequest("PUT", url, data)
  if err != nil {
    fmt.Printf("Error: %v\n", err)
  }
  client := http.Client{}
  _, err = client.Do(req)
  if err != nil {
    fmt.Printf("Error: %v\n", err)
  }
  //fmt.Printf("Response: %v\n", resp)
}

func ToggleLight(lightnum int) {
  var state = LightState{}
  state.On = !GetLightStatus(lightnum).On
  SetLightState(lightnum, state)
}

func AllOn() {
   s := DefaultState
   s.Hue = 4500
   s.Sat = 255
   s.Bri = 255
   b, err := json.Marshal(s)
   if err != nil {
     fmt.Println("error:", err)
   }
   DoPut(groupsUrl(0, "/action"), bytes.NewReader(b))
}

func Show() {
  allOff := DefaultAction
  allOff.On = false
  allOff.Hue = 40000
  quickOn := DefaultAction
  quickOn.Bri = 255
  quickOn.Sat = 0

  quickDim := DefaultAction
  quickDim.Bri = 150
  quickDim.Transition = 1

  slowDim := DefaultAction
  slowDim.Bri = 10
  slowDim.Transition = 10

  SetGroupState(allOff)
  Delay(500)
  SetGroupState(quickOn)
  SetGroupState(quickDim)
  Delay(100)
  SetGroupState(slowDim)
  Delay(900)

  for {
  for i := 1; i <= 3; i++ {
    SetLightState(i, quickOn)
    SetLightState(i, quickDim)
    Delay(100)
    SetLightState(i, slowDim)
  }
  Delay(900)
  }
}


func LightLoop(reps int) {
  AllOn()
  <-time.After(time.Second)
  onmsg := DefaultAction
  offmsg := DefaultAction
  offmsg.On = false
  onmsg.Hue = 100
  onmsg.Bri = 255
  //delay := 100.0

  for {
    onmsg.Hue += 5000
    if onmsg.Hue > 65000 {
      onmsg.Hue = 1500
    }
  for i:= 1; i <= 3; i++ {
  SetLightState(i, offmsg)
  Delay(200)
  go SetLightState(i, onmsg)
  }
  }
}

func GetLightStatus(lightnum int) LightState {
  resp, err := http.Get(buildLightsUrl(lightnum, ""))
  if err != nil {
    fmt.Printf("Error: %v\n", err)
  }
    fmt.Printf("Response Body: %v\n, resp.Body")
  dec := json.NewDecoder(resp.Body)
  var m LightMessage
  err = dec.Decode(&m)
  if err != nil {
    fmt.Printf("Error: %v\n", err)
    fmt.Printf("Response Body: %v\n, resp.Body")
  }
  return m.State
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
    fmt.Printf("Response Body: %v\n, resp.Body")
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
  fmt.Println(resp.Body)
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

