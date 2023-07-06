package main

import (
	"fmt"
	"math"
	"week1/linkedlist"
)

type Cube struct {
	Length float32
}

type Cylinder struct {
	Radius float32
	Height float32
}

type GeometryCalc interface {
	volume() float32
	setValue(value float32, ptr *float32)
	whatType()
}

func main() {
	var cube GeometryCalc = &Cube{Length: 32.0}
	var cylinder GeometryCalc = &Cylinder{Radius: 2.0, Height: 5.0}

	fmt.Printf("Volume of cube: %v \n", cube.volume())
	fmt.Printf("Volume of cylinder: %v \n", cylinder.volume())

	cube.setValue(10.0, &cube.(*Cube).Length)
	fmt.Printf("Volume of new length cube: %v \n", cube.volume())

	
	fmt.Printf("Volume of cylinder: %v \n", cylinder.volume())

	cube.whatType()
	cylinder.whatType()

	var linkedList linkedlist.Collections = &linkedlist.LinkedList{}

	linkedList.Add(10)
	linkedList.Add(20)
	linkedList.Add(30)

	linkedList.Display()
	fmt.Println(linkedList.Length())
}

func (cube *Cube) volume() float32 {
	return float32(math.Pow(float64(cube.Length), 3))
}

func (cube *Cube) setValue(value float32, ptr *float32) {
	*ptr = value
}

func (cylinder *Cylinder) volume() float32 {
	return float32(math.Pi * math.Pow(float64(cylinder.Radius), 2) * float64(cylinder.Height))
}

func (cube *Cylinder) setValue(value float32, ptr *float32) {
	*ptr = value
}

func (cube Cube) whatType() {
	fmt.Printf("Type %T\n", cube)
}

func (cylinder Cylinder) whatType() {
	fmt.Printf("Type %T\n", cylinder)
}