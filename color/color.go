package color

import (
  "math"
  "fmt"
)

type RGB struct {
  r, g, b int
}

type HSB struct {
  h, s, b float64
}


func (rgb *RGB) Hex() string {
  return fmt.Sprintf("#%x%x%x", rgb.r, rgb.g, rgb.b)

}

func RGBtoHSB(rgb RGB) HSB {
  r := float64(rgb.r)/255.0
  g := float64(rgb.g)/255.0
  b := float64(rgb.b)/255.0
  var min = math.Min(r, math.Min(g, b))
  var max = math.Max(r, math.Max(g, b))

  var delta = max - min

  hsb := HSB{ b: max }
  if max != 0 {
    hsb.s = delta/max
  }
  if hsb.s <= 0 {
    hsb.h = 0
  } else {
    if r == max {
      hsb.h = (g - b) / delta
      if (g < b) {
        hsb.h += 6
      }
    } else if g == max {
      hsb.h = 2 + (b - r) / delta
    } else if b == max {
      hsb.h = 4 + (r - g) / delta
    }
    hsb.h *= 60

  }

  return hsb
}

func HSBtoRGB(hsb HSB) RGB {
  c := hsb.b * hsb.s
  h := hsb.h/60



  x := c*(1 - math.Abs(math.Remainder(h, 2) - 1))

  var r, g, b float64
  switch int(h) {
    case 0: r = c; g = x; b = 0
    case 1: r = x; g = c; b = 0
    case 2: r = 0; g = c; b = x
    case 3: r = 0; g = x; b = c
    case 4: r = x; g = 0; b = c
    case 5: r = c; g = 0; b = x
  }
  m := hsb.b - c
  r += m
  g += m
  b += m
  return RGB{ r:int(255*r), g:int(255*g), b:int(255*b)}

//
//  i := int(math.Floor(h *6))
//  f := h*6 - float64(i)
//  p := int(255*hsb.b * (1 - hsb.s))
//  q := int(255*hsb.b * (1 - f*hsb.s))
//  t := int(255*hsb.b * (1 - (1-f)*hsb.s))
//  v := int(255*hsb.b)
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
