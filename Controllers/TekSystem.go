package Controllers

import "fmt"

// type Rectangle struct {
// 	Length int
// 	Breadth int
// }
// type Shape interface {
// 	Area() float64
// 	Perimeter() float64
// }

// func (r Rectangle) Ar{
// 	// val:=interface{Area(6,4)}

// 	fmt.Println("Area: ", val) // val:=interface{Area(6,4)}

// }
/*
package Controllers

import "fmt"

func TeckSystem() {
	//make slices
	myMap := make(map[string]int)
	myMap["MyVal1"] = 56
	fmt.Println("myMap: ", myMap["MyVal2"])

	//myMap["MyVal2"]

	return
	s1 := make([]int, 5)
	fmt.Println(s1)
	s1[2] = 5000
	fmt.Println(s1)
	s1 = append(s1, 1000)
	fmt.Println(s1)
}

//[0,0,0,0,0]
//[0,0,5000,0,0]
//[0,0,5000,0,0,1000,0,0,0,0,0]

// To be learned
func(reciver_name Type) method_name(parameter_list)(return_type){
// Code
}
*/

// Define an interface called Shape
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Define a struct for a rectangle
type Rectangle struct {
	Width  float64
	Height float64
}

// Define a struct for a circle
type Circle struct {
	Radius float64
}

// Implement methods for Rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2*r.Width + 2*r.Height
}

// Implement methods for Circle
func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * 3.14 * c.Radius
}

// Function that takes any shape and calculates area and perimeter
func PrintShapeInfo(s Shape) {
	fmt.Printf("Area: %f\n", s.Area())
	fmt.Printf("Perimeter: %f\n", s.Perimeter())
}

func MyFunc1() {
	// Create a rectangle
	rectangle := Rectangle{Width: 5, Height: 3}

	// Create a circle
	// circle := Circle{Radius: 4}

	// Call function PrintShapeInfo with rectangle
	fmt.Println("Rectangle:")
	PrintShapeInfo(rectangle)

	// Call function PrintShapeInfo with circle
	// fmt.Println("Circle:")
	// PrintShapeInfo(circle)
}
func TeckSystem() {
	//make slices
	myMap := make(map[string]int)
	// myMap["MyVal1"] = 56
	fmt.Println(myMap)
	// s1 := make([]int, 4)
	// fmt.Println(s1)
	// s1[1] = 5000
	// fmt.Println(s1)
	// s1 = append(s1, 1000, 8, 3, 5, 8, 6, 8)
	// fmt.Println(s1)
	// fmt.Println(cap(s1))
}
