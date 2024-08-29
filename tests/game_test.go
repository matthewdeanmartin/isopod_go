package tests

import (
	"isopod_game/game"
	"testing"
)

func TestNewGame(t *testing.T) {
	theGame := game.NewGame()

	if theGame.CurrentLocation != "Leaf Pile" {
		t.Errorf("expected starting location to be 'Leaf Pile', got '%s'", theGame.CurrentLocation)
	}

	if len(theGame.Inventory) != 0 {
		t.Errorf("expected inventory to be empty at start, got %d items", len(theGame.Inventory))
	}

	if len(theGame.Locations) != 5 {
		t.Errorf("expected 5 locations, got %d", len(theGame.Locations))
	}

	if len(theGame.Items) != 3 {
		t.Errorf("expected 3 items, got %d", len(theGame.Items))
	}
}

func TestMove(t *testing.T) {
	theGame := game.NewGame()

	theGame.Move("Rock")
	if theGame.CurrentLocation != "Rock" {
		t.Errorf("expected current location to be 'Rock', got '%s'", theGame.CurrentLocation)
	}

	theGame.Move("Log")
	if theGame.CurrentLocation != "Log" {
		t.Errorf("expected current location to be 'Log', got '%s'", theGame.CurrentLocation)
	}

	// Test invalid move
	theGame.Move("Pond")
	if theGame.CurrentLocation == "Pond" {
		t.Errorf("expected move to 'Pond' to fail from 'Log', but it succeeded")
	}
}

func TestLook(t *testing.T) {
	theGame := game.NewGame()

	theGame.Look()
	if len(theGame.Inventory) != 0 {
		t.Errorf("expected inventory to be empty after looking at 'Leaf Pile', got %d items", len(theGame.Inventory))
	}

	theGame.Move("Rock")
	theGame.Look()
	if !theGame.Inventory["Place to Hide"] {
		t.Errorf("expected to find 'Place to Hide' in 'Rock', but it wasn't found")
	}
}

func TestCheckWinCondition(t *testing.T) {
	theGame := game.NewGame()

	if theGame.CheckWinCondition() {
		t.Error("expected CheckWinCondition to be false initially")
	}

	theGame.Inventory["Place to Hide"] = true
	theGame.Inventory["Cookie Crumb"] = true
	theGame.Inventory["Another Isopod Friend"] = true

	if !theGame.CheckWinCondition() {
		t.Error("expected CheckWinCondition to be true after collecting all items")
	}
}
