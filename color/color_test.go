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
            rgb: RGB{r:255, g:255, b:0},
            hsb: HSB{h:60, s:1, b:1},
          },
  TestCase{ name: "white",
            rgb: RGB{r:255, g:255, b:255},
            hsb: HSB{h:0, s:0, b:1},
          },
  TestCase{ name: "purple",
            rgb: RGB{r:170, g:0, b:255},
            hsb: HSB{h:280, s:1, b:1},
          },
  TestCase{ name: "navy",
            rgb: RGB{r:25, g:25, b:112},
            hsb: HSB{h:240, s:0.78, b:0.44},
          },
  TestCase{ name: "black",
            rgb: RGB{r:0, g:0, b:0},
            hsb: HSB{h:0, s:0, b:0},
          },
}

func roundHSB(hsb HSB) HSB {
  hsb.h = float64(int(hsb.h + 0.5))
  hsb.s = float64(int(100*hsb.s + 0.5))/100
  hsb.b = float64(int(100*hsb.b + 0.5))/100
  return hsb
}

func difference(a, b int) int {
  if a > b {
    return a - b
  }
  return  b - a

}

func withinOne(a, b RGB) bool {
  return difference(a.r, b.r) <= 1 &&
  difference(a.g, b.g) <= 1 &&
  difference(a.b, b.b) <= 1
}

func GotButExpected(msg string, got,expected interface{}, t *testing.T) {
  t.Errorf("%v Got: %v, but expected %v", msg, got, expected)
}


func Test_RGBtoHex(t *testing.T) {
  foo := RGB{ r:123, g:55, b:255}
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

