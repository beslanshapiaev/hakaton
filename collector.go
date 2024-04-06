package main

var MyShip *ShipBasket

func xyPieceToRowCol(pieceXY [][]int) [][]int {
	pieceRC := make([][]int, len(pieceXY))
	for i, point := range pieceXY {
		pieceRC[i] = []int{point[1], point[0]}
	}

	return pieceRC
}

type ShipBasket struct {
	Width           int
	Height          int
	Board           [][]int
	EmptyCells      int
	AvalibleGarbage map[string][][]int
}

func NewShipBasket(width, height int) *ShipBasket {
	board := make([][]int, height)
	emptyCells := width * height
	for i := range board {
		board[i] = make([]int, width)
	}
	return &ShipBasket{
		Width:           width,
		Height:          height,
		Board:           board,
		EmptyCells:      emptyCells,
		AvalibleGarbage: make(map[string][][]int, 0),
	}
}

func (ship *ShipBasket) canPlace(piece [][]int, row, col int) bool {
	for _, point := range piece {
		r, c := point[0], point[1]
		if row+r < 0 || row+r >= ship.Height || col+c < 0 || col+c >= ship.Width || ship.Board[row+r][col+c] != 0 {
			return false
		}
	}
	return true
}

func (ship *ShipBasket) placePiece(piece [][]int, row, col, value int) [][]int {
	pieceRCPos := make([][]int, 0)

	for _, point := range piece {
		r, c := point[0], point[1]
		ship.Board[row+r][col+c] = value
		pieceRCPos = append(pieceRCPos, []int{row + r, col + c})
		if value != 0 {
			ship.EmptyCells--
		} else {
			ship.EmptyCells++
		}
	}

	return pieceRCPos
}

func rotatePiece(piece [][]int) [][]int {
	// Определение размера фигуры
	size := len(piece)

	// Создание новой фигуры для хранения повернутой версии
	rotatedPiece := make([][]int, size)
	for i := range rotatedPiece {
		rotatedPiece[i] = make([]int, 2) // каждая точка имеет две координаты
	}

	// Поворот фигуры на 90 градусов против часовой стрелки
	for i := 0; i < size; i++ {
		rotatedPiece[i][0] = piece[size-i-1][1] // переворачиваем x и y координаты
		rotatedPiece[i][1] = piece[size-i-1][0]
	}

	return rotatedPiece
}

func (ship *ShipBasket) findBestPosition(piece [][]int) (int, int, [][]int) {
	for row := 0; row < ship.Height; row++ {
		for col := 0; col < ship.Width; col++ {
			for rotation := 0; rotation < 4; rotation++ {
				if ship.canPlace(piece, row, col) {
					return row, col, piece
				}
				piece = rotatePiece(piece)
			}
		}
	}

	return -1, -1, piece
}

func (ship *ShipBasket) Collect() map[string][][]int {
	result := make(map[string][][]int)
	for name, pieceXY := range MyShip.AvalibleGarbage {
		piece := xyPieceToRowCol(pieceXY)
		row, col, piece := MyShip.findBestPosition(piece)
		if row != -1 && col != -1 {
			pos := xyPieceToRowCol(MyShip.placePiece(piece, row, col, 1))
			result[name] = pos
		}
	}

	return result
}
