package game

import (
	"fmt"
	"os"
	"strings"

	"github.com/BurntSushi/toml"
)

// Structs for TOML data
type Config struct {
	Game      GameConfig          `toml:"game"`
	Locations map[string]Location `toml:"locations"`
	Items     map[string]string   `toml:"items"`
}

type GameConfig struct {
	WelcomeText   string   `toml:"welcome_text"`
	WinText       string   `toml:"win_text"`
	StartLocation string   `toml:"start_location"` // New field for starting location
	WinConditions []string `toml:"win_conditions"` // New field for win conditions
}

type Location struct {
	Description string            `toml:"description"`
	Directions  map[string]string `toml:"-"` // Dynamically filled in after parsing
	// Adding fields to parse directions directly from TOML
	North string `toml:"north,omitempty"`
	South string `toml:"south,omitempty"`
	East  string `toml:"east,omitempty"`
	West  string `toml:"west,omitempty"`
}

// Game state
type Game struct {
	CurrentLocation string
	Inventory       map[string]bool
	Locations       map[string]Location
	Items           map[string]string
	Config          GameConfig
}

// Initialize a new game from TOML configuration
func NewGame(configFile string) *Game {
	var config Config
	if _, err := toml.DecodeFile(configFile, &config); err != nil {
		fmt.Println("Error loading configuration:", err)
		os.Exit(1)
	}

	// Set up directions for locations from the TOML data
	for name, location := range config.Locations {
		location.Directions = make(map[string]string)

		// Populate the directions from TOML data
		if location.North != "" {
			location.Directions["north"] = location.North
		}
		if location.South != "" {
			location.Directions["south"] = location.South
		}
		if location.East != "" {
			location.Directions["east"] = location.East
		}
		if location.West != "" {
			location.Directions["west"] = location.West
		}

		config.Locations[name] = location
	}

	// Initialize game with the start location from config
	return &Game{
		CurrentLocation: config.Game.StartLocation,
		Inventory:       map[string]bool{},
		Locations:       config.Locations,
		Items:           config.Items,
		Config:          config.Game,
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
	case command == "quit":
		fmt.Println("Thanks for playing!")
		return false
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
	fmt.Println("  go <direction> - Move to a different location.")
	fmt.Println("  quit       - Quit the game.")
}

// Describe the current location and list possible moves
func (g *Game) Look() {
	loc := g.Locations[g.CurrentLocation]
	fmt.Printf("%s\n", loc.Description)
	if item, found := g.Items[g.CurrentLocation]; found {
		fmt.Printf("You found a %s here!\n", item)
		g.Inventory[item] = true
		delete(g.Items, g.CurrentLocation) // Remove the item from the map
	} else {
		fmt.Println("There's nothing special here.")
	}

	// List available directions and locations
	if len(loc.Directions) > 0 {
		fmt.Println("You can go to:")
		for direction, destination := range loc.Directions {
			fmt.Printf("  %s - %s\n", direction, destination)
		}
	} else {
		fmt.Println("There are no available places to go from here.")
	}
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
func (g *Game) Move(direction string) {
	if location, ok := g.Locations[g.CurrentLocation].Directions[direction]; ok {
		g.CurrentLocation = location
		fmt.Printf("You moved to the %s.\n", location)
		g.Look()
	} else {
		fmt.Println("You can't go there from here.")
	}
}

// Check if the player has won
func (g *Game) CheckWinCondition() bool {
	for _, item := range g.Config.WinConditions {
		if !g.Inventory[item] {
			return false
		}
	}
	return true
}
