package main

import (
	"fmt"
	"time"
)

type Game struct {
  Previous *Grid
  Current *Grid
  CellTicked chan bool
}

type Grid struct {
	Width      int
	Height     int
	Cells      [][]*Cell
}

type Cell struct {
	Alive bool
}

func (c *Cell) String() string {
	if c.Alive {
		return "o"
	} else {
		return "#"
	}
}

func newGrid(width, height int) *Grid {
	grid := &Grid{Width: width, Height: height}

	grid.Cells = make([][]*Cell, height, height)

	for y := 0; y < height; y++ {
		grid.Cells[y] = make([]*Cell, height, height)
		for x := 0; x < width; x++ {
			grid.Cells[y][x] = &Cell{Alive: false}
		}
	}

	return grid
}

func (grid *Grid) String() string {
	grid_string := ""

	for y := 0; y < grid.Height; y++ {
		for x := 0; x < grid.Width; x++ {
			grid_string += grid.Cells[y][x].String()
		}
		grid_string += "\n"
	}

	return grid_string
}

func (grid *Grid) copy() *Grid {
	grid_copy := &Grid{Width: grid.Width, Height: grid.Height}

	grid_copy.Cells = make([][]*Cell, grid.Height, grid.Height)

	for y := 0; y < grid.Height; y++ {
		grid_copy.Cells[y] = make([]*Cell, grid.Height, grid.Height)

		for x := 0; x < grid.Width; x++ {
			grid_copy.Cells[y][x] = grid.Cells[y][x]
		}
	}

	return grid_copy
}

func (game *Game) live_neighbours(x, y int) int {
	live_neighbours := 0

	if x > 0 && game.Previous.Cells[y][x-1].Alive {
		live_neighbours += 1
	}

	if x > 0 && y < game.Previous.Height-1 && game.Previous.Cells[y+1][x-1].Alive {
		live_neighbours += 1
	}

	if y < game.Previous.Height-1 && game.Previous.Cells[y+1][x].Alive {
		live_neighbours += 1
	}

	if x < game.Previous.Width-1 && y < game.Previous.Height-1 && game.Previous.Cells[y+1][x+1].Alive {
		live_neighbours += 1
	}

	if x < game.Previous.Width-1 && game.Previous.Cells[y][x+1].Alive {
		live_neighbours += 1
	}

	if x < game.Previous.Width-1 && y > 0 && game.Previous.Cells[y-1][x+1].Alive {
		live_neighbours += 1
	}

	if y > 0 && game.Previous.Cells[y-1][x].Alive {
		live_neighbours += 1
	}

	if y > 0 && x > 0 && game.Previous.Cells[y-1][x-1].Alive {
		live_neighbours += 1
	}

	return live_neighbours
}

func (game *Game) tick_cell(x, y int) {
	live_neighbours := game.live_neighbours(x, y)

	if game.Previous.Cells[y][x].Alive {
		if live_neighbours < 2 {
			game.Current.Cells[y][x] = &Cell{Alive: false}
		} else if live_neighbours > 3 {
			game.Current.Cells[y][x] = &Cell{Alive: false}
		} else {
			game.Current.Cells[y][x] = &Cell{Alive: true}
		}
	} else {
		if live_neighbours == 3 {
			game.Current.Cells[y][x] = &Cell{Alive: true}
		} else {
			game.Current.Cells[y][x] = &Cell{Alive: false}
		}
	}

	game.CellTicked <- true
}

func (game *Game) tick() {
	game.Previous = game.Current.copy()

	for y := 0; y < game.Current.Height; y++ {
		for x := 0; x < game.Current.Width; x++ {
			go game.tick_cell(x, y)
		}
	}
}


func main() {
	width := 30
	height := 30

	grid := newGrid(width, height)

	game := &Game{Current: grid, CellTicked: make(chan bool, width * height)}

	game.Current.Cells[15][15] = &Cell{true}
	game.Current.Cells[14][16] = &Cell{true}
	game.Current.Cells[13][14] = &Cell{true}
	game.Current.Cells[13][15] = &Cell{true}
	game.Current.Cells[13][16] = &Cell{true}

	fmt.Println("\033c")
	fmt.Print(game.Current)

	for i := 0; i <= 100000; i++ {
		game.tick()
		for i := 1; i <= game.Current.Width*game.Current.Height; i++ {
			<-game.CellTicked
		}
		time.Sleep(100 * time.Millisecond)

		fmt.Println("\033c")
		fmt.Print(game.Current)
	}
}
