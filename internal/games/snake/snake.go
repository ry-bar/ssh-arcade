// This is where the snake game state logic will be written.

package snake

import (
	"strings"
//  "fmt"

  tea "charm.land/bubbletea/v2"
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
	width int
	height int

	snake []point
	food point

	direction direction

	gameOver bool
	score int
}


type stopGame string


func stop() tea.Msg {
	return stopGame("stop")
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
	// Maybe this is where I would gather the user data that's needed for the game?
	return nil
}


func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {
		case tea.KeyPressMsg:
			switch msg.String() {
				case "q", "crtl+c":
					return m, tea.Quit
			}

	}
 
	return m, nil
}



//TODO: This function is hella broken
func (m model) View() tea.View{

	s := strings.Repeat("-", m.width) + "\n"


	internalWidth := m.width - 2
	if internalWidth < 0 {
		internalWidth = 0
	}

	internalHeight := m.height - 2
	if internalHeight < 0 {
		internalHeight = 0
	}


	for i := 0; i < internalHeight; i++ {
		s += "|" + strings.Repeat(" ", internalWidth) + "|\n"
	}

	s += strings.Repeat("-", m.width)
	

	return tea.NewView(s)
}

