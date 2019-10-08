package main

import (
	"flag"
	"io/ioutil"
	"strconv"
	"strings"
	"time"

	"github.com/nsf/termbox-go"
)

var (
	ticks = flag.Int64("ticks", 10, "Number of ticks to simulate. -1 will run to negative overflow.")
	path  = flag.String("input_path", "", "A file path")
	delay = flag.Int64("delay", 1000, "Delay in milliseconds.")
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	flag.Parse()

	if *path == "" {
		panic("input_path must be set")
	}
	cells := initCells(*path)
	check(termbox.Init())
	defer termbox.Close()

	// Support never ending on -1
	for *ticks != 0 {
		drawCells(cells)
		time.Sleep(time.Millisecond * time.Duration(*delay))
		cells = simulate(cells)
		*ticks--
	}
}

func simulate(cells [][]int) [][]int {
	next := make([][]int, len(cells))
	for i := range cells {
		next[i] = make([]int, len(cells[i]))
		copy(next[i], cells[i])
	}

	for r := range next {
		for c := range next[r] {
			switch next[r][c] {
			case 0:
				next[r][c] = 0
			case 1:
				next[r][c] = 2
			case 2:
				next[r][c] = 3
			case 3:
				if updateConductor(r, c, cells) {
					next[r][c] = 1
				} else {
					next[r][c] = 3
				}
			default:
				next[r][c] = 0
			}
		}
	}

	return next
}

func updateConductor(r, c int, cells [][]int) bool {
	headCount := 0
	if r > 0 {
		if cells[r-1][c] == 1 {
			headCount++
		}

		if c > 0 {
			if cells[r-1][c-1] == 1 {
				headCount++
			}
		}

		if c < len(cells[r])-1 {
			if cells[r-1][c+1] == 1 {
				headCount++
			}
		}

	}

	if r < len(cells)-1 {
		if cells[r+1][c] == 1 {
			headCount++
		}
		if c > 0 {
			if cells[r+1][c-1] == 1 {
				headCount++
			}
		}

		if c < len(cells[r])-1 {
			if cells[r+1][c+1] == 1 {
				headCount++
			}
		}
	}

	if c > 0 {
		if cells[r][c-1] == 1 {
			headCount++
		}
	}

	if c < len(cells[r])-1 {
		if cells[r][c+1] == 1 {
			headCount++
		}
	}
	return headCount == 1 || headCount == 2
}

func drawCells(cells [][]int) {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	for r := range cells {
		for c := range cells[r] {
			color := termbox.ColorDefault
			switch cells[r][c] {
			case 1:
				color = termbox.ColorBlue
			case 2:
				color = termbox.ColorRed
			case 3:
				color = termbox.ColorYellow
			default:
				color = termbox.ColorBlack
			}
			termbox.SetCell(c, r, 'X', color, color)
		}
	}
	termbox.Flush()
}

func initCells(path string) [][]int {
	cells := [][]int{}
	// Trust file size fits in memory
	c, err := ioutil.ReadFile(path)
	check(err)
	f := string(c)
	for r, l := range strings.Split(f, "\n") {
		cells = append(cells, []int{})
		for _, v := range strings.Split(l, "") {
			i, err := strconv.Atoi(v)
			check(err)
			cells[r] = append(cells[r], i)
		}
	}
	return cells
}

func draw(cells [][]int) {

}
