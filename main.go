package main

import (
	"fmt"
	"strconv"
)

func main() {

	printPlayerMessage("asmongold", "belatra")
	printPlayerMessage("kiko")
	printPlayerMessage("pduts")
	printPlayerNameAndAge("indaysara", "25")
	printPlayerNameAndAge("bastiduts", "6")

}

func printPlayerNameAndAge(s1, s2 string) {
	age, err := strconv.Atoi(s2)
	if err != nil {
		fmt.Printf("s2 must be convertible to int, got %s\n", s2)
		return
	}

	fmt.Printf("Player name: %s Age: %d\n", s1, age)

}
func printPlayerMessage(player ...string) {
	fmt.Printf("player %s is on the starting line \n", player)
}
