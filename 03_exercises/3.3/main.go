// Surface computes an SVG rendering of a 3-D surface function.
package main

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, z := corner(i+1, j)
			isFinite := isCornerFinite(ax, ay)
			if !isFinite {
				continue
			}

			bx, by, _ := corner(i, j)
			isFinite = isCornerFinite(bx, by)
			if !isFinite {
				continue
			}

			cx, cy, _ := corner(i, j+1)
			isFinite = isCornerFinite(cx, cy)
			if !isFinite {
				continue
			}

			dx, dy, _ := corner(i+1, j+1)
			isFinite = isCornerFinite(dx, dy)
			if !isFinite {
				continue
			}

			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g'  stroke='%s' />\n", ax, ay, bx, by, cx, cy, dx, dy, polygonColor(z))
		}
	}

	fmt.Println("</svg>")
}

func corner(i, j int) (float64, float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale

	return sx, sy, z
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
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

func polygonColor(z float64) string {
	if z < 0 {
		return "#0000ff"
	}

	return "#ff0000"
}
