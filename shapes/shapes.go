package shapes

import "math"

// Shape interface
type Shape interface {
	Area() float64
}

// A Rectangle is not a square
type Rectangle struct {
	Width  float64
	Height float64
}

// Area returns the area of a rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Perimeter returns the perimeter of a rectangle
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// A Circle is also not a square
type Circle struct {
	Radius float64
}

// Area returns the area of a circle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// A Triangle is not a circle
type Triangle struct {
	Base   float64
	Height float64
}

// Area returns the area of a triangle
func (t Triangle) Area() float64 {
	return (t.Base * t.Height) * 0.5
}
