// Surface computes an SVG rendering of a 3-D surface function.
package main

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 120            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.2        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			isFinite := isCornerFinite(ax, ay)
			if !isFinite {
				continue
			}

			bx, by := corner(i, j)
			isFinite = isCornerFinite(bx, by)
			if !isFinite {
				continue
			}

			cx, cy := corner(i, j+1)
			isFinite = isCornerFinite(cx, cy)
			if !isFinite {
				continue
			}

			dx, dy := corner(i+1, j+1)
			isFinite = isCornerFinite(dx, dy)
			if !isFinite {
				continue
			}

			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n", ax, ay, bx, by, cx, cy, dx, dy)
		}
	}

	fmt.Println("</svg>")
}

func corner(i, j int) (float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	return math.Abs(math.Sin(x) + math.Sin(y))
}

func isCornerFinite(x float64, y float64) bool {
	if math.IsNaN(x) || math.IsInf(x, 0) {
		return false
	}
	if math.IsNaN(y) || math.IsInf(y, 0) {
		return false
	}

	return true
}