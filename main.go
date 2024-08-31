package main

import (
	"bufio"
	"fmt"
	"isopod_game/game"
	"os"
	"strings"
)

func main() {
	the_game := game.NewGame("data.toml")

	fmt.Println("Welcome to the Isopod Adventure!")
	fmt.Println("Your mission is to find a place to hide, a cookie crumb, and another isopod friend.")
	fmt.Println("Type 'help' for a list of commands.")

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "quit" {
			fmt.Println("Thanks for playing!")
			break
		}

		if the_game.HandleCommand(input) {
			if the_game.CheckWinCondition() {
				fmt.Println("Congratulations! You've found all the items and won the the_game!")
				break
			}
		}
	}
}
