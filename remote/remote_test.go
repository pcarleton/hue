package remote

import (
  "testing"
  "strings"
)

const hue_input =
"[{'id':'001788fffe0a4be2', 'internalipaddress':'10.0.1.2','macaddress':'00:17:88:0a:4b:e2'},{'id':'001788fffe0aec34','internalipaddress':'10.0.1.3'}]"

var json_dict = map[string] string {
  "id" : "001788fffe0a4be2",
  "internalipaddress" : "10.0.1.2",
  "macaddress" : "00:17:88:0a:4b:e2",
}

func GotButExpected(msg string, got,expected interface{}, t *testing.T) {
  t.Errorf("%v Got: %v, but expected %v", msg, got, expected)
}


func Test_GetStatus(t *testing.T) {
  GetStatusMsg()
  

}

//func Test_GetHueIp(t *testing.T) {
//  expectedIp := "10.0.1.2"
//  ip := GetHueIp()
//  if (ip != expectedIp) {
//    GotButExpected("Local hue IP", ip, expectedIp, t)
//  }
//}
//
func Test_ConvertJsonToMap_GoodInput(t *testing.T) {
  jsonMap, ok := ConvertJsonToMap(strings.NewReader(hue_input))
  if !ok {
    t.Errorf("Error converting json to map %v", jsonMap)
  }
  for k, v := range jsonMap {
    if json_dict[k] != v {
      GotButExpected("ConvertJsonToMap", v, json_dict[k], t)
    }
  }
}
