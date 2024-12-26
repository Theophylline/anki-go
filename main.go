package main

import (
	"fmt"
    "os"
	tea "github.com/charmbracelet/bubbletea"
)

/**
- create deck
    - if a new deck is created, automatically set to default
- delete deck
- set deck
    - the deck that add, remove, review will operate on
    - show some stats of the deck, like the number of cards, description, last time reviewed, etc
- add card
- edit card
    - shows a list of cards, then give option for user to edit
- remove card
- review
    - select deck to review
    - loop through cards using anki-like algorithm... until user asks to stop
**/

type model struct {
	menu []string
    cursor int
    activeWindow int
}

func initialModel() model {
    return model{
        menu: []string{"View Deck List", "View Cards in Deck", "Review"}
    }

} 

func (m model) Init() tea.Cmd {
    return nil
}

func (m model) View() string {
    return "work in progress"
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    return m, nil
}

func main() {
	init_state := model{
	}
    p := tea.NewProgram(init_state)
    if _, err := p.Run(); err != nil {
        fmt.Printf("Error initiating program: %v", err)
        os.Exit(1)
    }
}
