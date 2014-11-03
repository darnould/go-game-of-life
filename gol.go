package main

import (
	"fmt"
	"time"
)

type Grid struct {
	Width  int
	Height int
	Cells  [][]*Cell
}

type Cell struct {
	Alive bool
}

func (c *Cell) String() string {
	if c.Alive {
		return "*"
	} else {
		return "&"
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

func (grid *Grid) tick() *Grid {
	next_grid := grid.copy()

	for y := 0; y < grid.Height; y++ {
		for x := 0; x < grid.Width; x++ {
			live_neighbours := 0

			if x > 0 && grid.Cells[y][x-1].Alive {
				live_neighbours += 1
			}

			if x > 0 && y < grid.Height-1 && grid.Cells[y+1][x-1].Alive {
				live_neighbours += 1
			}

			if y < grid.Height-1 && grid.Cells[y+1][x].Alive {
				live_neighbours += 1
			}

			if x < grid.Width-1 && y < grid.Height-1 && grid.Cells[y+1][x+1].Alive {
				live_neighbours += 1
			}

			if x < grid.Width-1 && grid.Cells[y][x+1].Alive {
				live_neighbours += 1
			}

			if x < grid.Width-1 && y > 0 && grid.Cells[y-1][x+1].Alive {
				live_neighbours += 1
			}
			if y > 0 && grid.Cells[y-1][x].Alive {
				live_neighbours += 1
			}
			if y > 0 && x > 0 && grid.Cells[y-1][x-1].Alive {
				live_neighbours += 1
			}

			if grid.Cells[y][x].Alive {
				if live_neighbours < 2 {
					next_grid.Cells[y][x] = &Cell{Alive: false}
				} else if live_neighbours > 3 {
					next_grid.Cells[y][x] = &Cell{Alive: false}
				} else {
					next_grid.Cells[y][x] = &Cell{Alive: true}
				}
			} else {
				if live_neighbours == 3 {
					next_grid.Cells[y][x] = &Cell{Alive: true}
				} else {
					next_grid.Cells[y][x] = &Cell{Alive: false}
				}
			}
		}
	}

	return next_grid
}

func main() {
	width := 30
	height := 30

	grid := newGrid(width, height)

	grid.Cells[15][15] = &Cell{true}
	grid.Cells[14][16] = &Cell{true}
	grid.Cells[13][14] = &Cell{true}
	grid.Cells[13][15] = &Cell{true}
	grid.Cells[13][16] = &Cell{true}

	fmt.Println("\033c")
	fmt.Print(grid)

	for i := 0; i <= 100; i++ {
		time.Sleep(500 * time.Millisecond)

		fmt.Println("\033c")
		grid = grid.tick()
		fmt.Print(grid)
	}

}
