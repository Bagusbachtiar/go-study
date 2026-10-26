package main

import "fmt"

func GetName() string {
	name := ""

	fmt.Println("Welcome to the Go programming language tutorial!")
	fmt.Printf("Please enter your name: ")
	_, err := fmt.Scanln(&name)
	if err != nil {
		return ""
	}
	fmt.Printf("Hello, %s! Nice to meet you!\n", name)
	return name
}

func GetBet(balance uint) uint {
	var bet uint
	for {
		fmt.Printf("Enter your bet, or 0 to quit (balance: %d): ", balance)
		fmt.Scan(&bet)

		if bet > balance {
			fmt.Println("You cannot bet more than your balance.")
		} else {
			break
		}
	}
	return bet
}