package main

import (
    "fmt"
    "math"
)

const (
    width, height = 600, 320
    cells         = 100
    xyrange       = 30.0
    xyscale       = width / 2 / xyrange
    zscale        = height * 0.4
    angle         = math.Pi / 6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main() {
    fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
        "style='stroke: grey; fill: white; stroke-width: 0.7' "+
        "width='%d' height='%d'>", width, height)
    for i := 0; i < cells; i++ {
        for j := 0; j < cells; j++ {
            /*_, _, az := corner(i+1, j)
            _, _, bz := corner(i, j)
            _, _, cz := corner(i, j+1)
            _, _, dz := corner(i+1, j+1)
            c := getColor((az + bz + cz + dz) / 4)
            fmt.Println(c)*/
            
            ax, ay, az := corner(i+1, j)
            bx, by, bz := corner(i, j)
            cx, cy, cz := corner(i, j+1)
            dx, dy, dz := corner(i+1, j+1)
            fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g' fill='%s'/>\n",
                ax, ay, bx, by, cx, cy, dx, dy, getColor((az + bz + cz + dz) / 4))
        }
    }
    fmt.Println("</svg>")
}

func getColor(z float64) string {
    color := int(0xff * (math.Abs(z)))
    if color > 0xff {
        color = 0xff
    }
    strColor := fmt.Sprintf("%02x", color)

    if (z > 0) {
       return ("#" + strColor + "0000")
    } else {
       return ("#0000" + strColor)
    }
}

func corner(i, j int) (float64, float64, float64) {
    x := xyrange * (float64(i)/cells - 0.5)
    y := xyrange * (float64(j)/cells - 0.5)

    z := f(x, y)

    sx := width/2 + (x-y)*cos30*xyscale
    sy := height/2 + (x+y)*sin30*xyscale - z*zscale
    return sx, sy, z
}

func f(x, y float64) float64 {
    r := math.Hypot(x, y)
    return math.Sin(r) / r
}