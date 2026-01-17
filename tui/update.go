package tui

import (
	"strconv"

	tea "github.com/charmbracelet/bubbletea"
)

func (m *Model) Init() tea.Cmd {
	return nil
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch m.state {
		case SelectingSong:
			return m.handleSongSelection(msg)
		case AcceptingCoins:
			return m.handleCoinInput(msg)
		case PlayingSong:
			return m.handlePlayingSong(msg)
		}
	}
	return m, nil
}

func (m *Model) handleSongSelection(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "ctrl+c", "q":
		return m, tea.Quit
	case "up", "k":
		if m.cursor > 0 {
			m.cursor--
		}
	case "down", "j":
		if m.cursor < len(m.config.Songs)-1 {
			m.cursor++
		}
	case "enter":
		m.selectedSong = &m.config.Songs[m.cursor]
		m.state = AcceptingCoins
	}
	return m, nil
}

func (m *Model) handleCoinInput(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "ctrl+c", "q":
		return m, tea.Quit
	case "enter":
		coin, err := strconv.ParseFloat(m.paymentInput, 64)
		if err != nil {
			m.paymentInput = ""
			return m, nil
		}

		isValidCoin := false
		for _, acceptedCoin := range m.config.AcceptedCoins {
			if coin == acceptedCoin {
				isValidCoin = true
				break
			}
		}

		if isValidCoin {
			m.insertedAmount += coin
		}

		m.paymentInput = ""

		if m.insertedAmount >= m.selectedSong.Price {
			m.state = PlayingSong
		}

	case "backspace":
		if len(m.paymentInput) > 0 {
			m.paymentInput = m.paymentInput[:len(m.paymentInput)-1]
		}
	default:
		if _, err := strconv.ParseFloat(msg.String(), 64); err == nil || msg.String() == "." {
			m.paymentInput += msg.String()
		}
	}
	return m, nil
}

func (m *Model) handlePlayingSong(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "enter":
		*m = *NewModel(m.config)
		return m, nil
	}
	return m, nil
}
