// This will be the entry point for running the app locally.
package main

import (
	"fmt"
	"os"

	tea "charm.land/bubbletea/v2"

	"github.com/ry-bar/ascii-run/internal/games/snake"
)


// Basic implementation of a tea.Model interface
// Init() tea.Cmd
// Update(tea.Msg) (tea.Model, tea.Cmd)
// View() string

func returnMainMenu(m model) tea.Msg {

	m.activeGame = nil
	m.currentScreen = menuScreen

	return m
}



type model struct {
	currentScreen screen
	games []string
	cursorLocation int

	activeGame tea.Model
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

		activeGame: nil,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	if m.currentScreen == snakeScreen {
		updatedGame, cmd := m.activeGame.Update(msg)
		m.activeGame = updatedGame

		return m, cmd	
	}

	if m.currentScreen == menuScreen {
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
						if m.cursorLocation == 1 {
							m.currentScreen = snakeScreen
							m.activeGame = snake.New()

							return m, m.activeGame.Init()
						}
					case "q", "ctrl+c":
						return m, tea.Quit
				
				}
		}
	}

	return m, nil
}

func (m model) View() tea.View {
	s := ""

	if m.currentScreen == snakeScreen {
		return m.activeGame.View()
	}


	if m.currentScreen == menuScreen {
		s = "\nSelect a game:\n\n"

		for i, choice := range m.games {
			cursor := " "
			if m.cursorLocation == i {
				cursor = ">"
			}
			s += fmt.Sprintf("%s %s\n", cursor, choice)
		}

		s += "\n\nPress ctrl+c or q to quit"
	}


	return tea.NewView(s)
}


func main(){
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("ERROR: %v", err)
		os.Exit(1)
	}
}
