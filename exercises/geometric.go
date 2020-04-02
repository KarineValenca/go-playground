package main

import "fmt"

type Square struct {
	sideSize float64
}

type Circle struct {
	radius float64
}

type Shape interface {
	calculateArea() float64
}

func (s Square) calculateArea(sideSize float64) float64{
	squareArea := sideSize*2
	return squareArea
}

func (c Circle) calculateArea(radius float64) float64 {
	circleArea := 3.14*radius*radius
	return circleArea
}

func info(shape Shape) {
	fmt.Print(shape.calculateArea())
}

func main () {
	circle1 := Circle{
		12,
	}

	
	square1 := Square{
		10,
	}

	info(circle1)
	info(square1)
}