package main

import (
	"testing"
)

// provide 4x4 grid with middle 4 cells alive
func gridBlock() (gameGrid, Game) {
	return gameGrid{
		{false, false, false, false},
		{false, true, true, false},
		{false, true, true, false},
		{false, false, false, false},
	}, Game{width: 4, height: 4}
}

// 5x5 grid with middle 3 in middle row alive
func gridBlinker() (gameGrid, Game) {
	return gameGrid{
		{false, false, false, false, false},
		{false, false, false, false, false},
		{false, true, true, true, false},
		{false, false, false, false, false},
		{false, false, false, false, false},
	}, Game{width: 5, height: 5}
}

func gridBlinkerPrime() gameGrid {
	return gameGrid{
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
func gridGlider() (gameGrid, Game) {
	return gameGrid{
		{false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false},
		{false, false, false, true, false, false, false, false},
		{false, false, false, false, true, false, false, false},
		{false, false, true, true, true, false, false, false},
		{false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false},
	}, Game{width: 8, height: 9}
}

func gridGlider1() gameGrid {
	return gameGrid{
		{false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false},
		{false, false, true, false, true, false, false, false},
		{false, false, false, true, true, false, false, false},
		{false, false, false, true, false, false, false, false},
		{false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false},
	}
}

func gridGlider2() gameGrid {
	return gameGrid{
		{false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false},
		{false, false, false, false, true, false, false, false},
		{false, false, true, false, true, false, false, false},
		{false, false, false, true, true, false, false, false},
		{false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false},
	}
}

func gridGlider3() gameGrid {
	return gameGrid{
		{false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false},
		{false, false, false, true, false, false, false, false},
		{false, false, false, false, true, true, false, false},
		{false, false, false, true, true, false, false, false},
		{false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false},
	}
}

// this is back to the original shape, but shifted down and to the right by
// one in each direction
func gridGlider4() gameGrid {
	return gameGrid{
		{false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false},
		{false, false, false, false, true, false, false, false},
		{false, false, false, false, false, true, false, false},
		{false, false, false, true, true, true, false, false},
		{false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false},
	}
}

func TestIsAlive(t *testing.T) {
	grid, game := gridBlock()

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
		if grid.isAlive(&game, c.row, c.col) != c.alive {
			t.Errorf("wrong aliveness for cell at %d,%d", c.row, c.col)
		}
	}
}

func gridsEqual(t *testing.T, g1, g2 gameGrid) bool {
	height := len(g1)
	width := len(g1[0])
	for row := 0; row < height; row++ {
		for col := 0; col < width; col++ {
			if g1[row][col] != g2[row][col] {
				t.Logf("Grids don't match at row,col: %d,%d", row, col)
				return false
			}
		}
	}
	return true
}

func TestAgeBlockGrid(t *testing.T) {
	grid, game := gridBlock()

	if !gridsEqual(t, grid, grid.age(&game)) {
		t.Error("Block grid did not age properly")
	}
}

func TestAgeBlinkerGrid(t *testing.T) {
	grid, game := gridBlinker()

	if !gridsEqual(t, grid.age(&game), gridBlinkerPrime()) {
		t.Error("Blinker grid did not age properly")
	}
}

func TestAgeGliderGrid(t *testing.T) {
	g0, game := gridGlider()

	g1 := gridGlider1()
	if !gridsEqual(t, g0.age(&game), g1) {
		t.Error("Glider grid did not age 1 properly")
	}

	g2 := gridGlider2()
	if !gridsEqual(t, g1.age(&game), g2) {
		t.Error("Glider grid did not age 2 properly")
	}

	g3 := gridGlider3()
	if !gridsEqual(t, g2.age(&game), g3) {
		t.Log("improperly aged glider:")
		t.Log(g2.age(&game))
		t.Error("Glider grid did not age 3 properly")
	}

	g4 := gridGlider4()
	if !gridsEqual(t, g3.age(&game), g4) {
		t.Log("improperly aged glider:")
		t.Log(g3.age(&game))
		t.Error("Glider grid did not age 4 properly")
	}
}
