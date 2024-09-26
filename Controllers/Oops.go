package Controllers

import "fmt"

type Car struct{}

func (c Car) Start() {
	fmt.Println("Car started")
}

type Truck struct{}

func (t Truck) Start() {
	fmt.Println("Truck started")
}

// func main() {
// 	vehicles := []Vehicle{Car{}, Truck{}}
// 	for _, vehicle := range vehicles {
// 		vehicle.Start()
// 	}
// }
