package main

import (
	"fmt"
	"math"
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

type StructKey struct {
	Country, State string
}

func main() {
	fmt.Println("Hello World")

	var name string = "Ankit"

	msg := fmt.Sprintf("Hello %v", name)
	fmt.Println(msg)

	//! rune for char in go

	// fuctions
	// a := add(10, 15)
	// fmt.Println(a)

	//? multiple returns and ignoring return
	// sq, _ := multipleRet(7)
	// _, cu2 := multipleRet(9)
	// fmt.Println(sq)
	// fmt.Println(cu2)

	//* structs
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

	//* Error handeling

	// val, err := strconv.Atoi("43a")
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Printf("The int value of this string is %d\n", val)
	// }
	// var err error = errors.New("Custom error");

	// Loops
	// for i := 0; i < 10; i++ {
	// 	fmt.Println("Go is best for backend")
	// }

	//! For is Go's "while"
	// i := 10
	// for i > 0 {
	// 	fmt.Println("Go is a compiled language")
	// 	i -= 1
	// }

	//* Arrays

	// arr := [5]int{1, 2, 3, 4, 5}
	// for i := 0; i < len(arr); i++ {
	// 	fmt.Println(arr[i])
	// }

	//* Slice (reference of array)
	//? slice is built over an array
	//? if the size of slice becomes greater than array new array(2*n) is created and the old values are copied to the new location

	// s1 := arr[1:4]

	// for i := 0; i < len(s1); i++ {
	// 	fmt.Printf("%d ", s1[i])
	// }

	//? we can create slice without creating an array using make
	// name   (type)  (size) (max size of underlying arr)
	// mySlice := make([]int,5,10)
	// max size is opt
	// mySlice := make([]int,5)
	// shorthand for slice mySlice := []int{}

	// fmt.Println(test(1, 2, 3, 4, 5))

	//* 2D slice
	// mySlice := [][]int{}
	// mySlice = createMatrix(4, 4)
	// fmt.Print(mySlice)
	// better way of iterating slices
	// for i, arr := range mySlice {
	// 	for j, v := range arr {
	// 		fmt.Printf("mySlice[%d][%d] = %d\n", i, j, v)
	// 	}
	// }

	//* Maps (refrence as slice)
	//! Map := make(map[key_type]value_type)
	// or
	// myMap := map[string]int{
	// 	"Ankit": 14,
	// 	"Rahul": 46,
	// 	"Aryan": 21,
	// 	"Trey":  0,
	// }
	// myMap["KDB"] = 21

	// delete(myMap, "Trey")

	//? To check whether an element exists in map or not
	// val, ok := myMap["JP"]
	// if ok1 {
	// 	fmt.Println("gg")
	// } else {
	// 	fmt.Print("nope")
	// }
	//? or (better way)
	// if _ ok := map["JP"]; ok { //code } else { //code }

	//? Iterating map
	// for key, val := range myMap {
	// 	fmt.Printf("%s : %d\n", key, val)
	// }

	//* Struct as key (netter than using nested maps)

	countryAqiMap := make(map[StructKey]int)
	countryAqiMap[StructKey{"IND", "DEL"}] = 300
	countryAqiMap[StructKey{"IND", "MU"}] = 100
	countryAqiMap[StructKey{"ENG", "LON"}] = 150
	countryAqiMap[StructKey{"USA", "NY"}] = 200

	for key, val := range countryAqiMap {
		fmt.Printf("[%s][%s] : %d\n", key.Country, key.State, val)
	}

	//TODO: Currying: Takes function as an argument and returns another function.

	//! Defer keyword executes a function at the end of the function

	//! Closure: Those functions that uses var outside its body (HOF)

	//? Pointers are same as C++ & -> reference * -> dereference

}

//* Variadic functions
//? fuctions that can take multiple arguments

func test(nums ...int) int {
	n := len(nums)
	ans := 0
	for i := 0; i < n; i++ {
		ans += nums[i]
	}

	return ans
}

func createMatrix(rows, cols int) [][]int {
	matrix := make([][]int, 0)
	for i := 0; i < rows; i++ {
		temp := make([]int, 0)
		for j := 0; j < cols; j++ {
			temp = append(temp, (i+1)*(j+1))
		}
		matrix = append(matrix, temp)
	}
	return matrix
}
