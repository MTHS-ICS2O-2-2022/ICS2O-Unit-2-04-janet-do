// Created by: Mr Coxall
// Created on: Sep 2020
//
// This program finds the area of a triangle

package main

import "fmt"

func main() {
	// This function finds the area of a triangle

	var height int
	var base int
	var area int

	// input
	fmt.Println("This program finds the area of a triangle.")
	fmt.Println()
	fmt.Print("Enter the height (cm): ")
	fmt.Scanln(&height)
	fmt.Print("Enter the base (cm): ")
	fmt.Scanln(&base)

	// process
	area = base * height / 2

	// output
	fmt.Println()
	fmt.Println("The area is:", area, "cm².")
	fmt.Println("\nDone.")
}
