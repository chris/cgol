package main

import (
	"image/color"

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

type gameGrid [][]bool // note: use y,x (row, col) addressing

func (g *gameGrid) height() int {
	return len(*g)
}
func (g *gameGrid) width() int {
	return len((*g)[0])
}

func newGameGrid(rows, cols int) gameGrid {
	grid := make(gameGrid, rows)
	for row := 0; row < rows; row++ {
		grid[row] = make([]bool, cols)
	}
	return grid
}

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

func main() {
	ebiten.SetWindowSize(windowSize, windowSize)
	ebiten.SetWindowTitle("Conway's Game of Life")

	dimension := windowSize / lifeSize
	grid := newGameGrid(dimension, dimension)
	grid.fromCells(gospersGliderGun)
	g := Game{
		width:  dimension,
		height: dimension,
		grid:   &grid,
	}

	if err := ebiten.RunGame(&g); err != nil {
		panic(err)
	}
}

// Update applies game changes for a tick.
func (g *Game) Update(screen *ebiten.Image) error {
	g.age()
	return nil
}

// Draw renders the game for the latest tick.
func (g *Game) Draw(screen *ebiten.Image) {
	// paint the background
	screen.Fill(screenColor)

	for row := 0; row < g.height; row++ {
		for col := 0; col < g.width; col++ {
			if g.isAlive(row, col) {
				op := &ebiten.DrawImageOptions{}
				op.GeoM.Translate(float64(col*lifeSize), float64(row*lifeSize))
				screen.DrawImage(life, op)
			}
		}
	}
}

// Layout sets the scaled game size. We just use same size as screen.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return windowSize, windowSize
}

// isAlive determines if the grid cell at x,y is alive or not.
// It can be passed a coordinate that is one cell off the edge of the grid
// in any direction, and returns the torroidal version (e.g. moving off
// to the left x direction uses the furthest right x cell, and vice versa,
// and similarly for the y direction, as if the edges wrap around and connect
// to each other).
func (g *Game) isAlive(row, col int) bool {
	if col == g.width {
		col = 0
	} else if col < 0 {
		col = g.width - 1
	}

	if row == g.height {
		row = 0
	} else if row < 0 {
		row = g.height - 1
	}

	return (*g.grid)[row][col]
}

// ageCell applies the rules of the game to an individual cell
func (g *Game) ageCell(row, col int) bool {
	// determine number of live neighbors
	liveNeighbors := []bool{
		// row above
		g.isAlive(row-1, col-1), // upper left
		g.isAlive(row-1, col),
		g.isAlive(row-1, col+1),
		// same row as target
		g.isAlive(row, col-1),
		g.isAlive(row, col+1),
		// row below
		g.isAlive(row+1, col-1),
		g.isAlive(row+1, col),
		g.isAlive(row+1, col+1), // lower right
	}

	numAlive := 0
	for _, l := range liveNeighbors {
		if l {
			numAlive++
		}
	}

	if g.isAlive(row, col) {
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
func (g *Game) age() {
	agedGrid := newGameGrid(g.height, g.width)

	for row := 0; row < g.height; row++ {
		for col := 0; col < g.width; col++ {
			agedGrid[row][col] = g.ageCell(row, col)
		}
	}

	g.grid = &agedGrid
}

// Populate this grid with a automaton
func (g *gameGrid) fromCells(cells []cellCoord) {
	for _, cell := range cells {
		(*g)[cell.row][cell.col] = true
	}
}
