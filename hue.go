// Library functions for working with HUE
package hue

import  (
  "fmt"
  "net/http"
  "encoding/json"
  "io"
)

const meethue = "http://www.meethue.com/api/nupnp"
const developer_name = "newdeveloper"

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

