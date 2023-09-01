package simulation

type Canvas [][]int

func newCanvas(h int, w int) Canvas {
	canvas := make(Canvas, h)
	for rowId := range canvas {
		canvas[rowId] = make([]int, w)
	}

	return canvas
}

func (c Canvas) setState(points [][]int) {
	for _, point := range points {
		row, col := point[0], point[1]
		c[row][col] = 1
	}
}

func conwayStep(currentState Canvas) [][]int {
	h := len(currentState)
	w := len(currentState[0])
	resultCanvas := make(Canvas, h)
	for id := range currentState {
		resultCanvas[id] = make([]int, w)
	}

	for rowId := range currentState {
		for colId := range currentState[rowId] {
			resultCanvas[rowId][colId] = deadOrAlive(
				currentState,
				h,
				w,
				rowId,
				colId,
			)
		}
	}

	return resultCanvas
}

func deadOrAlive(canvas Canvas, h int, w int, rowId int, colId int) int {
	cellAlive := canvas[rowId][colId]
	res := cellAlive

	aliveNeighbours := 0

	ids := [][]int{
		{rowId - 1, colId},
		{rowId - 1, colId + 1},
		{rowId, colId + 1},
		{rowId + 1, colId + 1},
		{rowId + 1, colId},
		{rowId + 1, colId - 1},
		{rowId, colId - 1},
		{rowId - 1, colId - 1},
	}

	for _, pos := range ids {
		if 0 <= pos[0] && pos[0] < h && 0 <= pos[1] && pos[1] < w {
			aliveNeighbours += canvas[pos[0]][pos[1]]
		}
	}

	if cellAlive == 1 && (aliveNeighbours < 2 || aliveNeighbours > 3) {
		res = 0
	}

	if cellAlive == 0 && aliveNeighbours == 3 {
		res = 1
	}

	return res
}
