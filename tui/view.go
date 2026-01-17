package tui

import (
	"fmt"
	"strings"
)

func (m *Model) View() string {
	switch m.state {
	case SelectingSong:
		return m.viewSongSelection()
	case AcceptingCoins:
		return m.viewCoinInput()
	case PlayingSong:
		return m.viewPlayingSong()
	default:
		return "Something went wrong."
	}
}

func (m *Model) viewSongSelection() string {
	s := "Select a song:\n\n"
	for i, song := range m.config.Songs {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}
		s += fmt.Sprintf("%s %s (%.2f)\n", cursor, song.Name, song.Price)
	}
	s += "\nPress q to quit.\n"
	return s
}

func (m *Model) viewCoinInput() string {
	s := fmt.Sprintf("Selected song: %s\n", m.selectedSong.Name)
	s += fmt.Sprintf("Price: %.2f\n", m.selectedSong.Price)
	s += fmt.Sprintf("Amount inserted: %.2f\n", m.insertedAmount)
	s += fmt.Sprintf("\nEnter coin (accepted: %s): %s", m.acceptedCoinsString(), m.paymentInput)
	return s
}

func (m *Model) viewPlayingSong() string {
	change := m.insertedAmount - m.selectedSong.Price
	s := fmt.Sprintf("Now playing: %s\n", m.selectedSong.Name)
	if change > 0 {
		s += fmt.Sprintf("Your change: %.2f\n", change)
	}
	s += "\nPress Enter to select another song."
	return s
}

func (m *Model) acceptedCoinsString() string {
	var b strings.Builder
	for i, coin := range m.config.AcceptedCoins {
		b.WriteString(fmt.Sprintf("%.2f", coin))
		if i < len(m.config.AcceptedCoins)-1 {
			b.WriteString(", ")
		}
	}
	return b.String()
}
