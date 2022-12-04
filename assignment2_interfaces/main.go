package main

import "fmt"

type shape interface {
	getArea() float64
}

type square struct{
	sideLength float64
}

type triangle struct{
	height float64
	base float64
}+


func main() {

	mySquare := square{
		sideLength: 15.0,
	}

	myTriangle := triangle{
		height: 5.0,
		base: 2.0,
	}

	printArea(mySquare)
	printArea(myTriangle)

}
*-/+
func (sq square) getArea() float64 {*-+*
	return sq.sideLength * sq.sideLength
}

func (tr triangle) getArea() float64 {
	return 0.5 * tr.base * tr.height
}

func printArea(sh shape) {
	fmt.Println(sh.getArea())
}