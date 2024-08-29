package game

import (
	"fmt"
	"strings"
)

// Game state
type Game struct {
	CurrentLocation string
	Inventory       map[string]bool
	Locations       map[string][]string
	Items           map[string]string
}

// Initialize a new game
func NewGame() *Game {
	return &Game{
		CurrentLocation: "Leaf Pile",
		Inventory:       map[string]bool{},
		Locations: map[string][]string{
			"Leaf Pile":     {"Rock", "Log", "Pond"},
			"Rock":          {"Leaf Pile", "Log"},
			"Log":           {"Leaf Pile", "Rock"},
			"Pond":          {"Leaf Pile"},
			"Isopod's Home": {"Leaf Pile"},
		},
		Items: map[string]string{
			"Rock":          "Place to Hide",
			"Log":           "Cookie Crumb",
			"Isopod's Home": "Another Isopod Friend",
		},
	}
}

// Handle player commands
func (g *Game) HandleCommand(command string) bool {
	switch {
	case command == "help":
		g.PrintHelp()
	case command == "look":
		g.Look()
	case command == "Inventory":
		g.ShowInventory()
	case strings.HasPrefix(command, "go "):
		direction := strings.TrimPrefix(command, "go ")
		g.Move(direction)
	default:
		fmt.Println("Unknown command. Type 'help' for a list of commands.")
	}
	return true
}

// Print the help message
func (g *Game) PrintHelp() {
	fmt.Println("Commands:")
	fmt.Println("  look       - Look around the current location.")
	fmt.Println("  Inventory  - Check your Inventory.")
	fmt.Println("  go <place> - Move to a different location.")
	fmt.Println("  quit       - Quit the game.")
}

// Describe the current location
func (g *Game) Look() {
	fmt.Printf("You are at the %s.\n", g.CurrentLocation)
	if item, found := g.Items[g.CurrentLocation]; found {
		fmt.Printf("You found a %s here!\n", item)
		g.Inventory[item] = true
		delete(g.Items, g.CurrentLocation) // Remove the item from the map
	} else {
		fmt.Println("There's nothing special here.")
	}
	fmt.Printf("You can go to: %v\n", g.Locations[g.CurrentLocation])
}

// Show the player's Inventory
func (g *Game) ShowInventory() {
	if len(g.Inventory) == 0 {
		fmt.Println("Your Inventory is empty.")
		return
	}
	fmt.Println("Your Inventory contains:")
	for item := range g.Inventory {
		fmt.Printf("  - %s\n", item)
	}
}

// Move to a different location
func (g *Game) Move(location string) {
	if destinations, ok := g.Locations[g.CurrentLocation]; ok {
		for _, dest := range destinations {
			if strings.EqualFold(dest, location) {
				g.CurrentLocation = dest
				fmt.Printf("You moved to the %s.\n", dest)
				g.Look()
				return
			}
		}
	}
	fmt.Println("You can't go there from here.")
}

// Check if the player has won
func (g *Game) CheckWinCondition() bool {
	return g.Inventory["Place to Hide"] && g.Inventory["Cookie Crumb"] && g.Inventory["Another Isopod Friend"]
}
