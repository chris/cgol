package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten"
)

const (
	lifeSize   = 10
	tps        = 6 // game update rate: ticks per second
	windowSize = 600
)

var life *ebiten.Image
var lifeColor = color.RGBA{0x2c, 0x03, 0xfc, 0xff}
var screenColor = color.RGBA{0xff, 0xff, 0xff, 0xff}
var pos = 0.0

type gameGrid [][]bool // note: use y,x (row, col) addressing

// Define some game starting designs
type cellCoord struct {
	row int
	col int
}

// Game is the interface required by Ebiten
type Game struct {
	width  int
	height int
	grid   *gameGrid
}

func init() {
	ebiten.SetMaxTPS(tps)

	life, _ = ebiten.NewImage(lifeSize, lifeSize, ebiten.FilterDefault)
	life.Fill(lifeColor)
}

func (g *Game) Update(screen *ebiten.Image) error {
	*g.grid = g.grid.age(g)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// paint the background
	screen.Fill(screenColor)

	for row := 0; row < g.height; row++ {
		for col := 0; col < g.width; col++ {
			if g.grid.isAlive(g, row, col) {
				op := &ebiten.DrawImageOptions{}
				op.GeoM.Translate(float64(col*lifeSize), float64(row*lifeSize))
				screen.DrawImage(life, op)
			}
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return windowSize, windowSize
}

func makeGrid(rows, cols int) gameGrid {
	grid := make(gameGrid, rows)
	for row := 0; row < rows; row++ {
		grid[row] = make([]bool, cols)
	}
	return grid
}

func main() {
	ebiten.SetWindowSize(windowSize, windowSize)
	ebiten.SetWindowTitle("Conway's Game of Life")

	dimension := windowSize / lifeSize
	grid := makeGrid(dimension, dimension)
	grid.fromCells(gospersGliderGun)
	g := Game{
		width:  dimension,
		height: dimension,
		grid:   &grid,
	}

	if err := ebiten.RunGame(&g); err != nil {
		log.Fatal(err)
	}
}

// isAlive determines if the grid cell at x,y is alive or not.
// It can be passed a coordinate that is one cell off the edge of the grid
// in any direction, and returns the torroidal version (e.g. moving off
// to the left x direction uses the furthest right x cell, and vice versa,
// and similarly for the y direction, as if the edges wrap around and connect
// to each other).
func (grid *gameGrid) isAlive(game *Game, row, col int) bool {
	if col == game.width {
		col = 0
	} else if col < 0 {
		col = game.width - 1
	}

	if row == game.height {
		row = 0
	} else if row < 0 {
		row = game.height - 1
	}

	return (*grid)[row][col]
}

// ageCell applies the rules of the game to an individual cell
func (grid *gameGrid) ageCell(game *Game, row, col int) bool {
	// determine number of live neighbors
	liveNeighbors := []bool{
		// row above
		grid.isAlive(game, row-1, col-1), // upper left
		grid.isAlive(game, row-1, col),
		grid.isAlive(game, row-1, col+1),
		// same row as target
		grid.isAlive(game, row, col-1),
		grid.isAlive(game, row, col+1),
		// row below
		grid.isAlive(game, row+1, col-1),
		grid.isAlive(game, row+1, col),
		grid.isAlive(game, row+1, col+1), // lower right
	}

	numAlive := 0
	for _, l := range liveNeighbors {
		if l {
			numAlive++
		}
	}

	imAlive := grid.isAlive(game, row, col)

	if imAlive {
		if numAlive < 2 { // rule 1
			return false
		} else if numAlive > 3 { // rule 3
			return false
		}
		return true // rule 2
	}
	// dead
	return numAlive == 3 // rule 4
}

// ageGrid ages all cells in the grid and returns a new updated grid
func (grid *gameGrid) age(game *Game) gameGrid {
	agedGrid := makeGrid(game.height, game.width)

	for row := 0; row < game.height; row++ {
		for col := 0; col < game.width; col++ {
			agedGrid[row][col] = grid.ageCell(game, row, col)
		}
	}

	return agedGrid
}

// Make this grid a  heavyweight spaceship
func (grid *gameGrid) fromCells(cells []cellCoord) {
	for _, cell := range cells {
		(*grid)[cell.row][cell.col] = true
	}
}

var hwss = []cellCoord{
	{4, 4}, {4, 5}, {4, 6}, {4, 7}, {4, 8}, {4, 9},
	{5, 3}, {5, 9},
	{6, 9},
	{7, 3}, {7, 8},
	{8, 5}, {8, 6},
}

var pulsar = []cellCoord{
	{3, 5}, {3, 6}, {3, 7}, {3, 11}, {3, 12}, {3, 13},
	{5, 3}, {5, 8}, {5, 10}, {5, 15},
	{6, 3}, {6, 8}, {6, 10}, {6, 15},
	{7, 3}, {7, 8}, {7, 10}, {7, 15},
	{8, 5}, {8, 6}, {8, 7}, {8, 11}, {8, 12}, {8, 13},

	{10, 5}, {10, 6}, {10, 7}, {10, 11}, {10, 12}, {10, 13},
	{11, 3}, {11, 8}, {11, 10}, {11, 15},
	{12, 3}, {12, 8}, {12, 10}, {12, 15},
	{13, 3}, {13, 8}, {13, 10}, {13, 15},
	{15, 5}, {15, 6}, {15, 7}, {15, 11}, {15, 12}, {15, 13},
}

var gospersGliderGun = []cellCoord{
	{2, 26},
	{3, 24}, {3, 26},
	{4, 14}, {4, 15}, {4, 22}, {4, 23}, {4, 36}, {4, 37},
	{5, 13}, {5, 17}, {5, 22}, {5, 23}, {5, 36}, {5, 37},
	{6, 2}, {6, 3}, {6, 12}, {6, 18}, {6, 22}, {6, 23},
	{7, 2}, {7, 3}, {7, 12}, {7, 16}, {7, 18}, {7, 19}, {7, 24}, {7, 26},
	{8, 12}, {8, 18}, {8, 26},
	{9, 13}, {9, 17},
	{10, 14}, {10, 15},
}
