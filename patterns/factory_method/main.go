package main

import "fmt"

// Vehicle общий интерфейс для всех транспортных средств:
type Vehicle interface {
	Drive() string
}

type Car struct{}

func (c *Car) Drive() string {
	return "Автомобиль едет"
}

type Bicycle struct{}

func (b *Bicycle) Drive() string {
	return "Велосипед едет"
}

func NewVehicle(vehicleType string) Vehicle {
	switch vehicleType {
	case "car":
		return &Car{}
	case "bicycle":
		return &Bicycle{}
	default:
		return nil
	}
}

func main() {
	car := NewVehicle("car")
	fmt.Println(car.Drive())

	bicycle := NewVehicle("bicycle")
	fmt.Println(bicycle.Drive())
}
