package color

import (
  "testing"
)

type TestCase struct {
  name string
  rgb RGB
  hsb HSB
}

var test_cases = []TestCase{
  TestCase{ name: "yellow",
            rgb: RGB{R:255, G:255, B:0},
            hsb: HSB{H:60, S:1, B:1},
          },
  TestCase{ name: "white",
            rgb: RGB{R:255, G:255, B:255},
            hsb: HSB{H:0, S:0, B:1},
          },
  TestCase{ name: "purple",
            rgb: RGB{R:170, G:0, B:255},
            hsb: HSB{H:280, S:1, B:1},
          },
  TestCase{ name: "navy",
            rgb: RGB{R:25, G:25, B:112},
            hsb: HSB{H:240, S:0.78, B:0.44},
          },
  TestCase{ name: "black",
            rgb: RGB{R:0, G:0, B:0},
            hsb: HSB{H:0, S:0, B:0},
          },
}

func roundHSB(hsb HSB) HSB {
  hsb.H = float64(int(hsb.H + 0.5))
  hsb.S = float64(int(100*hsb.S + 0.5))/100
  hsb.B = float64(int(100*hsb.B + 0.5))/100
  return hsb
}

func difference(a, b int) int {
  if a > b {
    return a - b
  }
  return  b - a

}

func withinOne(a, b RGB) bool {
  return difference(a.R, b.R) <= 1 &&
  difference(a.G, b.G) <= 1 &&
  difference(a.B, b.B) <= 1
}

func GotButExpected(msg string, got,expected interface{}, t *testing.T) {
  t.Errorf("%v Got: %v, but expected %v", msg, got, expected)
}


func Test_RGBtoHex(t *testing.T) {
  foo := RGB{ R:123, G:55, B:255}
  hex := "#7b37ff"
  if foo.Hex() != hex {
    GotButExpected("Hex", foo.Hex(), hex, t)
  }
}

func Test_HSBtoRGB(t *testing.T) {
  for _, test := range test_cases {
    rgb := HSBtoRGB(test.hsb)
    if !withinOne(test.rgb, rgb) {
      GotButExpected(test.name, rgb, test.rgb, t)
    }
  }
}

func Test_RGBtoHSB(t *testing.T) {
  for _, test := range test_cases {
    hsb := roundHSB(RGBtoHSB(test.rgb))
    if test.hsb != hsb {
      GotButExpected(test.name, hsb, test.hsb, t)
    }
  }
}

