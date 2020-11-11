package main

// Shape interface to provide methods we need for our shapes to implement
type Shape interface {
	Area() int
	Scale(x int)
	Parameter() int
}

type Rectangle struct {
	Width  int `json:"width"`
	Length int `json:"length"`
}

func (r *Rectangle) Area() int {
	return r.Length * r.Width
}

func (r *Rectangle) Parameter() int {
	return 2 * (r.Length + r.Width)
}

func (r *Rectangle) Scale(x int) {
	r.Width = r.Width * x
	r.Length = r.Length * x
}

type Circle struct {
	// TODO: add json tags for decoding from body to this struct
	X      int
	Y      int
	Radius int
}

//TODO: add implementation of Circle for Shape methods
