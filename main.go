package main

import (
	"fmt"
	"strconv"
)

func main() {

	printPlayerMessage("asmongold", "belatra")
	printPlayerMessage("kiko")

	printPlayerNameAndAge("indaysara", "25")
	printPlayerNameAndAge("bastiduts", 6)

}

func printPlayerNameAndAgeStr(s1, s2 string) {
	age, err := strconv.Atoi(s2)
	if err != nil {
		fmt.Printf("s2 must be convertible to int, got %s\n", s2)
		return
	}

	fmt.Printf("Player name: %s Age: %d\n", s1, age)

}

func printPlayerNameAndAgeInt(s1 string, s2 int) {
	fmt.Printf("player name: %s, and Age: %d\n", s1, s2)
}

func printPlayerNameAndAge(s1 string, s2 interface{}) {
	switch v := s2.(type) {
	case int:
		printPlayerNameAndAgeInt(s1, v)
	case string:
		printPlayerNameAndAgeStr(s1, v)
	default:
		fmt.Printf("s2 must be a string or an int,got %T\n", v)
	}

}

func printPlayerMessage(player ...string) {
	fmt.Printf("player %s is on the starting line \n", player)
}
