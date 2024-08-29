/*
Package game provides the core logic for the Isopod Adventure game.

The Isopod Adventure game is a simple text-based adventure where the player controls
an isopod on a mission to find three items: a place to hide, a cookie crumb, and another
isopod friend. The player navigates through different locations, interacts with the environment,
and collects items to win the game.

# Package Structure

  - Game: The Game struct represents the overall game state, including the player's current
    location, inventory, available locations, and items in the game.

  - NewGame: This function initializes and returns a new Game instance, setting up the
    initial game state.

  - HandleCommand: This method processes player input commands, allowing the player to
    move between locations, check their inventory, and interact with the game world.

- Move: This method allows the player to move between adjacent locations.

  - Look: This method describes the current location and reveals any items that can be
    collected.

  - CheckWinCondition: This method checks if the player has collected all required items
    to win the game.

# Usage

To start a new game, create a new Game instance using NewGame:

	game := game.NewGame()

You can then handle player commands in a loop, using the HandleCommand method:

	for {
	    // Read input from the player
	    // Process the input with game.HandleCommand(input)
	    // Check if the player has won with game.CheckWinCondition()
	}
*/
package game
