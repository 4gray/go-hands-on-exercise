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

func (r *Circle) Area() int {
	return int(float64(r.Radius * r.Radius) * math.Pi)
}

func (r *Circle) Perimeter() int {
	return int(float64(2 * r.Radius) * math.Pi)
}

func (r *Circle) Scale(x int) int {
	return int(float64(r.Radius * r.Radius) * math.Pi) * 2
}
