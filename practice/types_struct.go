package main

import "fmt"

func typestruct() {
	poodle := Dog{"Poodle", 20}
	fmt.Println(poodle)
	fmt.Printf("%+v\n", poodle)
	fmt.Printf("Breed: %v\nWeight: %v\n", poodle.Breed, poodle.Weight)
	poodle.Weight = 60
	fmt.Printf("Breed: %v\nWeight: %v\n", poodle.Breed, poodle.Weight)
}

// Dog:  If a struct type initial character start with uppercase it means exported, lower case means private
type Dog struct {
	Breed  string
	Weight int
}
