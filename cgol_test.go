package main

import (
	"testing"
)

// provide 4x4 grid with middle 4 cells alive
func gridBlock() Game {
	return Game{width: 4, height: 4, grid: &gameGrid{
		{false, false, false, false},
		{false, true, true, false},
		{false, true, true, false},
		{false, false, false, false},
	}}
}

// 5x5 grid with middle 3 in middle row alive
func gridBlinker() Game {
	return Game{width: 5, height: 5, grid: &gameGrid{
		{false, false, false, false, false},
		{false, false, false, false, false},
		{false, true, true, true, false},
		{false, false, false, false, false},
		{false, false, false, false, false},
	}}
}

func gridBlinkerPrime() *gameGrid {
	return &gameGrid{
		{false, false, false, false, false},
		{false, false, true, false, false},
		{false, false, true, false, false},
		{false, false, true, false, false},
		{false, false, false, false, false},
	}
}

// glider using a 7x7 grid to start (2 empty cells buffer on each side), but we
// add extra rows & columns to allow to test the initial pattern not being
// centered and non-square grids
func gridGlider() Game {
	return Game{width: 8, height: 9, grid: &gameGrid{
		{false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false},
		{false, false, false, true, false, false, false, false},
		{false, false, false, false, true, false, false, false},
		{false, false, true, true, true, false, false, false},
		{false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false},
	}}
}

func gridGlider1() Game {
	return Game{width: 8, height: 9, grid: &gameGrid{
		{false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false},
		{false, false, true, false, true, false, false, false},
		{false, false, false, true, true, false, false, false},
		{false, false, false, true, false, false, false, false},
		{false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false},
	}}
}

func gridGlider2() Game {
	return Game{width: 8, height: 9, grid: &gameGrid{
		{false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false},
		{false, false, false, false, true, false, false, false},
		{false, false, true, false, true, false, false, false},
		{false, false, false, true, true, false, false, false},
		{false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false},
	}}
}

func gridGlider3() Game {
	return Game{width: 8, height: 9, grid: &gameGrid{
		{false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false},
		{false, false, false, true, false, false, false, false},
		{false, false, false, false, true, true, false, false},
		{false, false, false, true, true, false, false, false},
		{false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false},
	}}
}

// this is back to the original shape, but shifted down and to the right by
// one in each direction
func gridGlider4() Game {
	return Game{width: 8, height: 9, grid: &gameGrid{
		{false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false},
		{false, false, false, false, true, false, false, false},
		{false, false, false, false, false, true, false, false},
		{false, false, false, true, true, true, false, false},
		{false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false},
	}}
}

func TestIsAlive(t *testing.T) {
	game := gridBlock()

	type cellTest struct {
		cellCoord
		alive bool
	}

	var cases = []cellTest{
		{cellCoord{0, 0}, false},
		{cellCoord{1, 1}, true},
		{cellCoord{2, 2}, true},
		{cellCoord{3, 3}, false},
		{cellCoord{-1, -1}, false}, // equates to 3,3
		{cellCoord{-1, 2}, false},  // equates to 3, 2
		{cellCoord{2, 4}, false},   // equates to 2, 0
	}

	for _, c := range cases {
		if game.isAlive(c.row, c.col) != c.alive {
			t.Errorf("wrong aliveness for cell at %d,%d", c.row, c.col)
		}
	}
}

func gridsEqual(t *testing.T, g1, g2 *gameGrid) bool {
	height := g1.height()
	width := g1.width()
	for row := 0; row < height; row++ {
		for col := 0; col < width; col++ {
			if (*g1)[row][col] != (*g2)[row][col] {
				t.Logf("Grids don't match at row,col: %d,%d", row, col)
				return false
			}
		}
	}
	return true
}

func TestAgeBlockGrid(t *testing.T) {
	game := gridBlock()
	origGrid := game.grid

	game.age()

	if !gridsEqual(t, origGrid, game.grid) {
		t.Error("Block grid did not age properly")
	}
}

func TestAgeBlinkerGrid(t *testing.T) {
	game := gridBlinker()
	game.age()

	if !gridsEqual(t, game.grid, gridBlinkerPrime()) {
		t.Error("Blinker grid did not age properly")
	}
}

func TestAgeGliderGrid(t *testing.T) {
	game := gridGlider()
	game.age()

	g1 := gridGlider1()
	if !gridsEqual(t, game.grid, g1.grid) {
		t.Error("Glider grid did not age 1 properly")
	}

	g1.age()
	g2 := gridGlider2()
	if !gridsEqual(t, g1.grid, g2.grid) {
		t.Error("Glider grid did not age 2 properly")
	}

	g2.age()
	g3 := gridGlider3()
	if !gridsEqual(t, g2.grid, g3.grid) {
		t.Log("improperly aged glider:")
		t.Error("Glider grid did not age 3 properly")
	}

	g3.age()
	g4 := gridGlider4()
	if !gridsEqual(t, g3.grid, g4.grid) {
		t.Log("improperly aged glider:")
		t.Error("Glider grid did not age 4 properly")
	}
}
