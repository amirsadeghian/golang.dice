package main

import ("fmt"
		"math/rand"
		"time")

func main(){
	//Generate a unix timestamp as the seed for Random function
	var seed int64 = time.Now().UnixNano()
	rand.Seed(seed)
	var randomNumber int = rand.Intn(6)
	drawDice(randomNumber)
}

func drawDice(randomNumber int) {
	if randomNumber == 1 {
		fmt.Println("┌─────────┐")
		fmt.Println("│         │")
		fmt.Println("│    ●    │")
		fmt.Println("│         │")
		fmt.Println("└─────────┘")
	} else if randomNumber == 2 {
		fmt.Println("┌─────────┐")
		fmt.Println("│ ●       │")
		fmt.Println("│         │")
		fmt.Println("│       ● │")
		fmt.Println("└─────────┘")
	} else if randomNumber == 3 {
		fmt.Println("┌─────────┐")
		fmt.Println("│ ●       │")
		fmt.Println("│    ●    │")
		fmt.Println("│       ● │")
		fmt.Println("└─────────┘")
	} else if randomNumber == 4 {
		fmt.Println("┌─────────┐")
		fmt.Println("│ ●     ● │")
		fmt.Println("│         │")
		fmt.Println("│ ●     ● │")
		fmt.Println("└─────────┘")
	} else if randomNumber == 5 {
		fmt.Println("┌─────────┐")
		fmt.Println("│ ●     ● │")
		fmt.Println("│    ●    │")
		fmt.Println("│ ●     ● │")
		fmt.Println("└─────────┘")
	} else if randomNumber == 6 {
		fmt.Println("┌─────────┐")
		fmt.Println("│ ●  ●  ● │")
		fmt.Println("│ ●  ●  ● │")
		fmt.Println("│ ●  ●  ● │")
		fmt.Println("└─────────┘")
	}
}