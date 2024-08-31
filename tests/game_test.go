package tests

import (
	"isopod_game/game"
	"testing"
)

func TestNewGame(t *testing.T) {
	theGame := game.NewGame("../data.toml")

	if theGame.CurrentLocation != "Garden" {
		t.Errorf("expected starting location to be 'Leaf Pile', got '%s'", theGame.CurrentLocation)
	}

	if len(theGame.Inventory) != 0 {
		t.Errorf("expected inventory to be empty at start, got %d items", len(theGame.Inventory))
	}

	if len(theGame.Locations) != 4 {
		t.Errorf("expected 4 locations, got %d", len(theGame.Locations))
	}

	if len(theGame.Items) != 3 {
		t.Errorf("expected 3 items, got %d", len(theGame.Items))
	}
}

func TestLook(t *testing.T) {
	theGame := game.NewGame("../data.toml")

	theGame.Look()
	if len(theGame.Inventory) != 1 {
		t.Errorf("expected inventory to be 1 after looking at 'Leaf Pile', got %d items", len(theGame.Inventory))
	}

}

func TestCheckWinCondition(t *testing.T) {
	theGame := game.NewGame("../data.toml")

	if theGame.CheckWinCondition() {
		t.Error("expected CheckWinCondition to be false initially")
	}

	//"A Place to Hide 🛏️", "Cookie Crumb 🍪", "Isopod Friend 🐾"
	theGame.Inventory["A Place to Hide 🛏️"] = true
	theGame.Inventory["Cookie Crumb 🍪"] = true
	theGame.Inventory["Isopod Friend 🐾"] = true

	if !theGame.CheckWinCondition() {
		t.Error("expected CheckWinCondition to be true after collecting all items")
	}
}
