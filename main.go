package main

import "fmt"

// structs
type Naruto struct {
	rating string
}

type DemoSlayer struct {
	rating string
}

// interface
type MovieRating interface {
	movie() string
}

// functions
func (n *Naruto) movie() string {
	return "Boruto: Naruto the movie"
}

func (d *DemoSlayer) movie() string {
	return "DemonSlayer: Hope the movie"
}

// struct example
func callStruct() {
	rating1 := Naruto{rating: "Awesome"}
	rating2 := DemoSlayer{rating: "Fantastic"}

	fmt.Println("-- Struct Example --")
	fmt.Println(rating1.movie())
	fmt.Println(rating2.movie())
}

// interface example
func reuseInterface(moveRating MovieRating) {
	fmt.Println(moveRating.movie())
}
func callInterface() {
	fmt.Println("\n-- Interface Example --")

	//way : 1
	// var rating1 MovieRating = &Naruto{"I love it!"}
	// var rating2 MovieRating = &DemoSlayer{"Great"}

	// fmt.Println(rating1.movie())
	// fmt.Println(rating2.movie())

	//way : 2
	reuseInterface(&Naruto{"I love it!"})
	reuseInterface(&DemoSlayer{"Great"})
}

func main() {
	callStruct()
	callInterface()
}
