package main

import (
	"fmt"
)

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func getArea(shape Shape) {

	fmt.Println(shape.Area())
}

func main() {

	r := Rectangle{Width: 5, Height: 10}
	getArea(r)

}
