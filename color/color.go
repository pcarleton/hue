package color

import (
  "math"
  "fmt"
)

type RGB struct {
  R, G, B int
}

type HSB struct {
  H, S, B float64
}


func (rgb RGB) Hex() string {
  return fmt.Sprintf("#%02X%02X%02X", rgb.R, rgb.G, rgb.B)

}

func RGBtoHSB(rgb RGB) HSB {
  r := float64(rgb.R)/255.0
  g := float64(rgb.G)/255.0
  b := float64(rgb.B)/255.0
  var min = math.Min(r, math.Min(g, b))
  var max = math.Max(r, math.Max(g, b))

  var delta = max - min

  hsb := HSB{ B: max }
  if max != 0 {
    hsb.S = delta/max
  }
  if hsb.S <= 0 {
    hsb.H = 0
  } else {
    if r == max {
      hsb.H = (g - b) / delta
      if (g < b) {
        hsb.H += 6
      }
    } else if g == max {
      hsb.H = 2 + (b - r) / delta
    } else if b == max {
      hsb.H = 4 + (r - g) / delta
    }
    hsb.H *= 60

  }

  return hsb
}

func HSBtoRGB(hsb HSB) RGB {
  c := hsb.B * hsb.S
  h := hsb.H/60

  x := c*(1 - math.Abs(math.Abs(math.Remainder(h, 2)) - 1))
  if x < 0 {
    fmt.Println("X is negative!!")
    fmt.Println(x)
    fmt.Println(math.Remainder(h, 2))
    fmt.Println(math.Abs(math.Remainder(h, 2) -1))
  }

  var r, g, b float64
  switch int(h) {
    case 0: r = c; g = x; b = 0
    case 1: r = x; g = c; b = 0
    case 2: r = 0; g = c; b = x
    case 3: r = 0; g = x; b = c
    case 4: r = x; g = 0; b = c
    case 5: r = c; g = 0; b = x
  }
  m := hsb.B - c
  r += m
  g += m
  b += m
  return RGB{ R:int(255*r), G:int(255*g), B:int(255*b)}

//
//  i := int(math.Floor(h *6))
//  f := h*6 - float64(i)
//  p := int(255*hsb.B * (1 - hsb.S))
//  q := int(255*hsb.B * (1 - f*hsb.S))
//  t := int(255*hsb.B * (1 - (1-f)*hsb.S))
//  v := int(255*hsb.B)
//
//  switch i % 6 {
//    case 0: rgb = RGB{ r:v, g:t, b:p}
//    case 1: rgb = RGB{ r:q, g:v, b:p}
//    case 2: rgb = RGB{ r:p, g:v, b:t}
//    case 3: rgb = RGB{ r:p, g:q, b:v}
//    case 4: rgb = RGB{ r:t, g:p, b:v}
//    case 5: rgb = RGB{ r:v, g:p, b:q}
//  }
}
