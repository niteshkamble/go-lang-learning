package main

import "fmt"

func main() {

	//print & scan
	fmt.Println("Welcome to the quiz game!")

	fmt.Printf("Enter Your Name: ")
	var name string
	fmt.Scan(&name)

	fmt.Printf("Hello %v, Please Enter Your Age : ", name)
	var age uint
	fmt.Scan(&age)

	//Age Validation
	if age >= 12 {
		fmt.Println("Yo! You can play!")
	} else {
		fmt.Println("You can't play.....")
		return
	}

	result := 0
	questioin := 2

	//Q1
	fmt.Printf("Name of 1st electronic computer is ENIAC or ENIAC-I? ")
	var ans string
	fmt.Scan(&ans)
	if ans == "ENIAC" {
		result += 1
		fmt.Println("Correct Ans!")
	} else {
		fmt.Println("Wrong Ans!")
	}

	//Q2
	fmt.Printf("When was React Native Published to the public? ")
	var year uint
	fmt.Scan(&year)
	if year == 2015 {
		result += 1
		fmt.Println("Correct Ans!")
	} else {
		fmt.Println("Wrong Ans!")
	}

	fmt.Printf("You have scored %v out of %v.\n", result, questioin)

}
