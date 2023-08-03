package main

import (
	"fmt"
	"os"
	"os/exec"
)

type Board [3][3]string

type Game struct {
	board       Board
	currentTurn string
	status      string
}

func clearScreen() {
	cmd := exec.Command("clear") // For Windows, use "cls"
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func printBoard(board Board) {
	clearScreen()
	for _, row := range board {
		fmt.Println(row)
	}
}

func initializeGame() *Game {
	return &Game{
		board:       Board{},
		currentTurn: "X",
		status:      "ongoing",
	}
}

func (g *Game) makeMove(row, col int) error {
	if row < 0 || row >= 3 || col < 0 || col >= 3 || g.board[row][col] != "" {
		return fmt.Errorf("invalid move")
	}

	g.board[row][col] = g.currentTurn
	if g.checkWin(row, col) {
		g.status = g.currentTurn + " wins!"
	} else if g.checkDraw() {
		g.status = "It's a draw!"
	}

	if g.currentTurn == "X" {
		g.currentTurn = "O"
	} else {
		g.currentTurn = "X"
	}

	return nil
}

func (g *Game) checkWin(row, col int) bool {
	player := g.board[row][col]

	// Check for a win in the row
	for i := 0; i < 3; i++ {
		if g.board[row][i] != player {
			break
		}
		if i == 2 {
			return true
		}
	}

	// Check for a win in the column
	for i := 0; i < 3; i++ {
		if g.board[i][col] != player {
			break
		}
		if i == 2 {
			return true
		}
	}

	// Check for a win in the diagonal (top-left to bottom-right)
	if row == col {
		for i := 0; i < 3; i++ {
			if g.board[i][i] != player {
				break
			}
			if i == 2 {
				return true
			}
		}
	}

	// Check for a win in the diagonal (top-right to bottom-left)
	if row+col == 2 {
		for i := 0; i < 3; i++ {
			if g.board[i][2-i] != player {
				break
			}
			if i == 2 {
				return true
			}
		}
	}

	return false
}

func (g *Game) checkDraw() bool {
	for _, row := range g.board {
		for _, cell := range row {
			if cell == "" {
				return false
			}
		}
	}
	return true
}

func main() {
	game := initializeGame()
	var row, col int

	for game.status == "ongoing" {
		printBoard(game.board)

		fmt.Printf("Player %s's turn. Enter row (0-2) and column (0-2) separated by a space: ", game.currentTurn)
		fmt.Scan(&row, &col)

		if err := game.makeMove(row, col); err != nil {
			fmt.Println("Invalid move. Try again.")
		}
	}

	printBoard(game.board)
	fmt.Println(game.status)
}
