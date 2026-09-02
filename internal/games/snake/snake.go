// This is where the snake game state logic will be written.

package snake

import (
	//"strings"
	//  "fmt"

	tea "charm.land/bubbletea/v2"
  "charm.land/lipgloss/v2"	

)





type point struct {
	x int
	y int
}

type direction int

const (
	up direction = iota
	down
	left
	right
)

type model struct {
	boardWidth int
	boardHeight int

	border lipgloss.Style

	snake []point
	food point

	direction direction

	gameOver bool
	score int
}


func New(boardHeight int, boardWidth int ) model {
	return model{

		boardHeight: boardHeight,
		boardWidth: boardWidth,


		snake: []point{
			{x:12,y:12},
			{x:12,y:13},
		},
		food: point{x:6,y:6},

		gameOver: false,
		score: 0,
	}
}



func (m model) Init() tea.Cmd {
	return nil
}


func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// This will take the commands that are passed through the main.go
	//	Update() function if it's not the "m" button.
 
	return m, nil
}


func (m model) View() tea.View{
	gameBoard := ""

	if m.gameOver == false {
		// This loads the boarder and puts the head of the snake
		// 	into the first spot, but I'm trying to get it to start
		// 	in the middle of the board.
		gameContent := lipgloss.NewStyle().
				Width(m.boardWidth).
				Height(m.boardHeight).
				BorderStyle(lipgloss.ASCIIBorder()).
				BorderForeground(lipgloss.Color("#04B575")).
				Render("O")

		gameBoard = lipgloss.Place(
			m.boardWidth,
			m.boardHeight,
			lipgloss.Center,
			lipgloss.Center,
			gameContent,
		)

		return tea.NewView(gameBoard)

	}

	

	return tea.NewView(gameBoard)
}

