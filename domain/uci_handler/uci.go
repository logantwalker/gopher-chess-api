package uci

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/logantwalker/gopher-chess-api/domain/engine"
	model "github.com/logantwalker/gopher-chess-api/models"
)

const defaultFEN = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"


func NewGame(c *gin.Context){
	g := engine.NewGame()
	board := engine.FormatBoard(g.Board)
	c.IndentedJSON(http.StatusOK, board)
}

func Command(c *gin.Context){
	var userCommand model.UciCommand
	if err := c.BindJSON(&userCommand); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	g := engine.NewGame()
	// Split the input into words
	words := strings.Fields(userCommand.UciString)

	// Check if the position command is correctly followed by 'fen' or 'startpos'
	if len(words) > 1 {
		if words[1] == "fen" {
			// The position command is followed by a FEN string.
			// Join the rest of the words to form the FEN string.
			fen := strings.Join(words[2:], " ")
			g.Board = engine.NewBoard(fen)
		} else if words[1] == "startpos" {
			// The position command is followed by 'startpos', so set the board to the initial position.
			g.Board = engine.NewBoard(defaultFEN)

			// If there are moves following 'startpos', apply them.
			if len(words) > 2 && words[2] == "moves" {
				g.Board = engine.NewBoard(defaultFEN)
				for _, moveStr := range userCommand.Moves {
					if m, err := engine.CreateMove(moveStr); err == nil {
						gen := engine.NewGenerator(g.Board)
						moves := gen.GenerateMoves()
			
						found := engine.Move{From: engine.Invalid}
			
						for _, move := range moves {
							if move.From == m.From && move.To == m.To {
								found = move
								break
							}
						}
						if found.From != engine.Invalid {
							g.Board.MakeMove(found)
						} else {
							fmt.Printf("illegal move\n")
						}
			
					}
				}
			}

		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error":"invalid position command"})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error":"invalid position command"})
	}

	boardString := engine.FormatBoard(g.Board)
	c.IndentedJSON(http.StatusOK, boardString)
}