package simulation

import (
	"time"

	"github.com/gdamore/tcell"
)

type Game struct {
	Running bool
	Cells   Canvas
	Screen  tcell.Screen
}

func NewGame(screen tcell.Screen) *Game {
	w, h := screen.Size()

	var res Game
	res.Cells = newCanvas(h, w)
	res.Screen = screen

	return &res
}

func (g *Game) Run(stop <-chan bool, fps int) {
	g.Running = true
	frameTime := (1000 / fps)
	for {
		g.Screen.Clear()
		g.DrawFrame()
		g.Screen.Show()
		select {
		case <-stop:
			g.Running = false
			return
		default:
			g.Cells = g.Step()
		}
		time.Sleep(time.Duration(frameTime) * time.Millisecond)
	}
}

func (g *Game) SetNonEmptyPoints(points [][]int) {
	g.Cells.setState(points)
}

func (g *Game) Step() Canvas {
	result := conwayStep(g.Cells)
	return result
}

func (g *Game) DrawFrame() {
	h := len(g.Cells)
	cellStyle := tcell.StyleDefault.
		Background(tcell.ColorBlack).
		Foreground(tcell.ColorWhite)

	for rowId := 0; rowId < h; rowId++ {
		for colId := 0; colId < len(g.Cells[0]); colId++ {
			var cellContent rune
			if g.Cells[rowId][colId] == 1 {
				cellContent = '*'
			} else {
				cellContent = ' '
			}
			g.Screen.SetContent(colId, rowId, cellContent, nil, cellStyle)
		}
	}
}

func (g *Game) AddBlock(h int, w int) {
	points := [][]int{
		{h, w},
		{h, w + 1},
		{h + 1, w},
		{h + 1, w + 1},
	}

	g.SetNonEmptyPoints(points)
}

func (g *Game) AddBeehive(h int, w int) {
	points := [][]int{
		{h, w + 1},
		{h, w + 2},
		{h + 1, w},
		{h + 1, w + 3},
		{h + 2, w + 1},
		{h + 2, w + 2},
	}

	g.SetNonEmptyPoints(points)
}

func (g *Game) AddBlinker(h int, w int) {
	points := [][]int{
		{h, w + 1},
		{h + 1, w + 1},
		{h + 2, w + 1},
	}

	g.SetNonEmptyPoints(points)
}

func (g *Game) AddBeacon(h int, w int) {
	g.AddBlock(h, w)
	g.AddBlock(h+2, w+2)
}

func (g *Game) String() string {
	var output string
	for rowId := range g.Cells {
		for colId := range g.Cells[rowId] {
			if g.Cells[rowId][colId] == 1 {
				output += "*"
			} else {
				output += " "
			}
		}
		output += "\n"
	}

	return output
}
