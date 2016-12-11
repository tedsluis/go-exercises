// Surface computes an SVG rendering of a 3-D surface function.
package main

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 200                 // number of grid cells
	xyrange       = 90.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)
func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, z1 := corner(i+1, j)
			bx, by, z2 := corner(i, j)
			cx, cy, z3 := corner(i, j+1)
			dx, dy, z4 := corner(i+1, j+1)
			if math.IsInf(ax, 0) || math.IsInf(ay, 0) || math.IsInf(bx, 0) || math.IsInf(by, 0) || math.IsInf(cx, 0) || math.IsInf(cy, 0) || math.IsInf(dx, 0) || math.IsInf(dy, 0) {
				continue
			}
			if z1 < z2 {
				z1 = z2
			}
			if z1 < z3 {
				z1 = z3
			}
			if z1 < z4 {
				z1 = z4
			}
			z := int((math.Sin(z1*0.06)*8388352)+8388352)/4 + 30000
			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g' fill='#%x'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy, z)
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
	z := 0.0
	if x > -30 && x < 30 && y > -30 && y < 30 {
		z = -0.4
	}
	return (math.Sin(r) / r) + (math.Sin(x)-math.Cos(y))/100 + z
}
