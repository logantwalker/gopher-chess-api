package engine

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Game represents a gochess game
type Game struct {
	board *Board
}

// NewGame creates a new gochess game and returns a reference
func NewGame() *Game {
	g := new(Game)
	g.board = NewBoard(defaultFEN)

	return g
}

// Run a given game
func (g *Game) Run() {

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		in := scanner.Text()

		if in == "quit" || in == "q" {
			break

		} else if in == "uci" {
			fmt.Println("id name gopher")
			fmt.Println("id author loganwalker")
			fmt.Println("uciok")
		}else if strings.HasPrefix(in, "setoption"){
			words := strings.Fields(in)
			if len(words) > 1 {
				if words[1] == "name" {
					option := strings.Join(words[2:3], " ")
					if option == "Move Overhead"{
						value := words[5]
						fmt.Println("move overhead: ",value)
					}
				} else {
					fmt.Printf("invalid position command\n")
				}
			} else {
				fmt.Printf("invalid uci command\n")
			}

		} else if strings.HasPrefix(in, "position") {
			// Split the input into words
			words := strings.Fields(in)
		
			// Check if the position command is correctly followed by 'fen' or 'startpos'
			if len(words) > 1 {
				if words[1] == "fen" {
					// The position command is followed by a FEN string.
					// Join the rest of the words to form the FEN string.
					fen := strings.Join(words[2:], " ")
					g.board = NewBoard(fen)
				} else if words[1] == "startpos" {
					// The position command is followed by 'startpos', so set the board to the initial position.
					g.board = NewBoard(defaultFEN)
		
					// If there are moves following 'startpos', apply them.
					if len(words) > 2 && words[2] == "moves" {
						g.board = NewBoard(defaultFEN)
						for _, moveStr := range words[3:] {
							if m, err := createMove(moveStr); err == nil {
								gen := NewGenerator(g.board)
								moves := gen.GenerateMoves()
					
								found := Move{From: Invalid}
					
								for _, move := range moves {
									if move.From == m.From && move.To == m.To {
										found = move
										break
									}
								}
								if found.From != Invalid {
									g.board.MakeMove(found)
								} else {
									fmt.Printf("illegal move\n")
								}
					
							}
						}
					}
				} else {
					fmt.Printf("invalid position command\n")
				}
			} else {
				fmt.Printf("invalid position command\n")
			}
		} else if in == "isready" {
			fmt.Println("readyok")
		}else if in == "moves" || in == "m" {
			gen := NewGenerator(g.board)
			printMoves(gen.GenerateMoves())

		} else if in == "turn"{
			fmt.Println(g.board.sideToMove)
		}else if in == "perft" {
			Perft(position1FEN, position1Table)

		} else if in == "perft2" {
			Perft(position2FEN, position2Table)

		} else if in == "ucinewgame" || in == "n" {
			g := new(Game)
			g.board = NewBoard(defaultFEN)

		} else if in == "fen" || in == "f" {
			fmt.Printf("%s\n", generateFEN(g.board))

		} else if in == "undo" || in == "u" {
			g.board.UndoMove()

		} else if strings.HasPrefix(in, "fen ") {
			g.board = NewBoard(in[4:])

		} else if in == "print" || in == "p" {
			fmt.Printf("%s\n", formatBoard(g.board))

		} else if in == "search" || in == "s" {
			Search(g.board)

		} else if strings.HasPrefix(in, "go") || in == "g" {
			move := Search(g.board)
			stringMove := SquareMap[move.From] + SquareMap[move.To]
			fmt.Println("bestmove ", stringMove)
			g.board.MakeMove(move)

		} else if in == "eval" || in == "e" {
			fmt.Printf("Score: %d\n", Evaluate(g.board))

		} else if in == "auto" || in == "a" {
			for g.board.status == statusNormal {
				g.board.MakeMove(Search(g.board))
				fmt.Printf("%s\n", formatBoard(g.board))
			}

		} else if m, err := createMove(in); err == nil {
			fmt.Println("making move")
			gen := NewGenerator(g.board)
			moves := gen.GenerateMoves()

			found := Move{From: Invalid}

			for _, move := range moves {
				if move.From == m.From && move.To == m.To {
					found = move
					break
				}
			}

			if found.From != Invalid {
				g.board.MakeMove(found)
			} else {
				fmt.Printf("illegal move\n")
			}

		}
	}
}
