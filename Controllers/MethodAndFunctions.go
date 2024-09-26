package Controllers

import "fmt"

func FunT() {
	u := User{Name: "Avanish", Email: "avanish@xyz.com", Mobile: 7999042555}
	u.Show()
}

type User struct {
	Name   string
	Mobile int
	Email  string
}

// type Admin struct {
// 	Employee()
// 	Students()
// }

func (u User) Show() {
	fmt.Println("Name: ", u.Name)
	fmt.Println("Mobile: ", u.Mobile)
	fmt.Println("Email: ", u.Email)
}

type Shapes interface {
	Area() float64
	Perimeter() float64
}
type Rectangles struct {
	Length  int
	Breadth int
}
type Circles struct {
	Radius int
}

func (r Rectangles) Area() float64 {
	return float64(r.Length) * float64(r.Breadth)
}
func (r Circles) Area() float64 {
	return 2 * 2.33 * float64(r.Radius) * float64(r.Radius)
}
func (r Rectangles) Perimeter() float64 {
	return 2 * 2.33 * float64(r.Length)
}
func Run() {
	r := Rectangles{Length: 4, Breadth: 5}
	fmt.Println("Area of Rect : ", r.Area())
	fmt.Println("Perimeter is: ", r.Perimeter())

	c := Circles{Radius: 5}
	fmt.Println("Area of Circle: ", c.Area())
}
