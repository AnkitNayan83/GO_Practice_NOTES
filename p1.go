package main

import (
	"fmt"
	"math"
	"strconv"
)

func add(x int, y int) int {
	return x + y
}

func multipleRet(x float64) (square int, cube int) {
	return int(math.Pow(x, 2)), int(math.Pow(x, 3))
}

type Person struct {
	name string
	age  int
}

func (p Person) intro() {
	fmt.Printf("My name is %s and my age is %d", p.name, p.age)
}

type Alpha struct {
	Person
	power int
}

// ------------------- Interfaces --------------------------

type shape interface {
	area() (area float64)
	perimeter() float64
}

// This struct is of type shape as it impements both the func specified on the shape interface
type rect struct {
	width  float64
	height float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perimeter() float64 {
	return 2 * (r.height + r.width)
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

// now i can create a func with type shape and pass any shape struct which has area and perimeter function

func calc(sh shape) {
	name := ""
	// method 1
	_, ok := sh.(circle)
	if ok {
		name = "circle"
	} else {
		name = "rectangle"
	}
	// method 2
	// switch t := sh.(type) {
	// case circle:
	// 	name = "circle"
	// case rect:
	// 	name = "rectangle"
	// default:
	// 	name = ""
	// }
	fmt.Printf("Area of %s is: %.2f\n", name, sh.area())
	fmt.Printf("Perimeter of %s is: %.2f\n", name, sh.perimeter())
}

// ----------------------- Multiple Interfaces ---------------

type animal interface {
	intro()
}

type landAnimal interface {
	walk()
}

type dog struct {
	breed string
}

func (d dog) intro() {
	fmt.Printf("Hello my breed is %s\n", d.breed)
}

func (d dog) walk() {
	fmt.Println("I can walk")
}

func dogIntro(d dog) {
	d.intro()
	d.walk()
}

func main() {
	fmt.Println("Hello World")

	var name string = "Ankit"

	msg := fmt.Sprintf("Hello %v", name)
	fmt.Println(msg)

	// fuctions
	// a := add(10, 15)
	// fmt.Println(a)

	// multiple retuens and ignoring return
	// sq, _ := multipleRet(7)
	// _, cu2 := multipleRet(9)
	// fmt.Println(sq)
	// fmt.Println(cu2)

	//  structs
	// person1 := Person{}
	// person1.name = "Ankit"
	// person1.age = 22
	// person1.intro()

	// if person1.age >= 18 {
	// 	fmt.Println("You can vote now")
	// } else {
	// 	fmt.Println("You are a kid")
	// }

	// person2 := Alpha{
	// 	power: 100,
	// 	Person: Person{
	// 		name: "Jhon",
	// 		age:  35,
	// 	},
	// }
	// fmt.Println(person2.name)

	// Interfaces
	// circle1 := circle{
	// 	radius: 7,
	// }

	// rectangle1 := rect{
	// 	width:  5,
	// 	height: 10,
	// }

	// calc(circle1)
	// calc(rectangle1)

	// dog1 := dog{
	// 	breed: "Doberman",
	// }
	// dogIntro(dog1)

	// Error handeling

	val, err := strconv.Atoi("43a")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("The int value of this string is %d\n", val)
	}

	// var err error = errors.New("Custom error");

}
