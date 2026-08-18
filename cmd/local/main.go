// This will be the entry point for running the app locally.
package main

import (
	"fmt"
	"os"

	tea "charm.land/bubbletea/v2"
)


// Basic implementation of a tea.Model interface
// Init() tea.Cmd
// Update(tea.Msg) (tea.Model, tea.Cmd)
// View() string

type model struct {
	currentScreen screen
	games []string
	cursorLocation int
}

type screen int

const (
	menuScreen screen = iota
	ticTacToeScreen
	snakeScreen
)

func initialModel() model{
	return model{
		currentScreen: menuScreen,
		games: []string{"Tic-Tac-Toe", "Snake"},
		cursorLocation: 0,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type){
		case tea.KeyPressMsg:
			switch(msg.String()) {
				case "w", "up", "k":
					if m.cursorLocation > 0{
						m.cursorLocation--
					}
				case "s", "down", "j":
					if m.cursorLocation < len(m.games)-1 {
						m.cursorLocation++
					}
				case "enter", "space":
					
				case "q", "ctrl+c":
					return m, tea.Quit
			
			}
	}
	return m, nil

}

func (m model) View() tea.View {
	s := "\nSelect a game:\n\n"

	for i, choice := range m.games {
		cursor := " "
		if m.cursorLocation == i {
			cursor = ">"
		}
		s += fmt.Sprintf("%s %s\n", cursor, choice)
	}
		
	s += "\n\nPress ctrl+c or q to quit"

	return tea.NewView(s)
}


func main(){
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("ERROR: %v", err)
		os.Exit(1)
	}
}
