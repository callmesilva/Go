package main

import (
	"errors"
	"fmt"
	"https://github.com/callmesilva/Go/tree/main/structs"
	"strconv"
)

func main() {
	printBarrier()
	defer printBarrier()
	// game()
	// basicVars()
	listShowcase()

}

func game() {
	status, err := play()
	if status < 1 {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Selected number was %v\n", status)
	}
}

func play() (int, error) {
	fmt.Printf("Please enter a number between 1 and 10: ")
	var input string
	fmt.Scanln(&input)

	num, err := strconv.Atoi(input)
	if err != nil {
		return 0, errors.New("That is not a number")
	}

	if num < 1 || num > 10 {
		return 0, errors.New("Number is not in range")
	}

	return num, nil
}

func printBarrier() {
	const barrier = "==================="
	fmt.Println(barrier)
}

func basicVars() {
	var age int8 = 35
	const PI float32 = 3.14
	greet := "Hello World!"
	var passesValidation bool = true
	var yob int16

	fmt.Printf("User current age %v\n", age)
	fmt.Printf("PI: %v\n", PI)
	fmt.Printf("%v\n", greet)
	fmt.Printf("User allowed status: %v\n", passesValidation)

	year := 2023
	yob = int16(year) - int16(age)

	fmt.Printf("User year of birth: %v\n", yob)
}

func listShowcase() {
	// //array
	// var rouletteFavorites [5]uint8 = [5]uint8{15, 22, 13, 14, 25}
	// for index, value := range rouletteFavorites {
	// 	fmt.Printf("The number %v position for the favorite number is: %v\n", (index + 1), value)
	// }

	// //slice
	// var teamAges []uint8 = make([]uint8, 10)
	// teamAges = append(teamAges, 3, 4, 1)
	// for index, value := range teamAges {
	// 	value = uint8(index) + 28
	// 	fmt.Printf("Current age: %v\n", value)
	// }

	// //map
	// var scores map[string]uint8 = make(map[string]uint8, 4)
	// scores["Dolphins"] = 3
	// scores["Chiefs"] = 5
	// scores["Broncos"] = 0
	// scores["Vikings"] = 1

	// for team, score := range scores {
	// 	fmt.Printf("The %v have a score of %v\n", team, score)
	// }

	c := structs.Customer
	c.First = "Enrique"
	c.Last = "Silva"
	c.Balance = 10_000
	c.Age = 35
	c.Height = 1.76

	fmt.Printf("%v %v (%v) has a balance of %v\n", c.first, c.last, c.age, c.height)
}
