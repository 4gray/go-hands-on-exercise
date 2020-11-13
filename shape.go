package main

import "math"

// Shape interface to provide methods we need for our shapes to implement
type Shape interface {
	Area() int
	Scale(x int)
	Perimeter() int
}

type Rectangle struct {
	Width  int `json:"width"`
	Length int `json:"length"`
}

func (r *Rectangle) Area() int {
	return r.Length * r.Width
}

func (r *Rectangle) Perimeter() int {
	return 2 * (r.Length + r.Width)
}

func (r *Rectangle) Scale(x int) {
	r.Width = r.Width * x
	r.Length = r.Length * x
}

type Circle struct {
	X      int `json:"x"`
	Y      int `json:"y"`
	Radius int `json:"radius"`
}

func (c *Circle) Area() int {
	return int(float64(c.Radius*c.Radius) * math.Pi)
}

func (c *Circle) Perimeter() int {
	return int(float64(2*c.Radius) * math.Pi)
}

func (c *Circle) Scale(x int) {
	c.Radius = c.Radius * 2
}
