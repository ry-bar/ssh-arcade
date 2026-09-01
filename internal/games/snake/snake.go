// This is where the snake game state logic will be written.

package snake

import (
	//"strings"
	//  "fmt"

	tea "charm.land/bubbletea/v2"
  "charm.land/lipgloss/v2"	

)

var style = lipgloss.NewStyle().
		Width(40).
		Height(20).
		BorderStyle(lipgloss.ASCIIBorder()).
		BorderForeground(lipgloss.Color("#04B575"))

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
	width int
	height int

	snake []point
	food point

	direction direction

	gameOver bool
	score int
}


func New() model {
	return model{
		width: 40,
		height: 20,

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
	s := ""

	if m.gameOver == false {
		styleString := style.Render()
		return tea.NewView(styleString)

	}

	

	return tea.NewView(s)
}

