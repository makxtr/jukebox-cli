package tui

import (
	"strings"
	"testing"

	"jukebox/cli/config"

	tea "github.com/charmbracelet/bubbletea"
)

func TestCoinValidation(t *testing.T) {
	cfg := &config.Config{
		AcceptedCoins: []float64{0.01, 0.05, 0.10, 0.25, 0.50, 1.00},
	}

	m := NewModel(cfg)
	m.selectedSong = &config.Song{Price: 1.50}
	m.state = AcceptingCoins

	// Test valid coin
	m.paymentInput = "0.25"
	m.Update(tea.KeyMsg{Type: tea.KeyEnter})

	if m.insertedAmount != 0.25 {
		t.Errorf("Expected inserted amount to be 0.25, but got %.2f", m.insertedAmount)
	}

	// Test invalid coin
	m.paymentInput = "0.03"
	m.Update(tea.KeyMsg{Type: tea.KeyEnter})

	if m.insertedAmount != 0.25 {
		t.Errorf("Expected inserted amount to remain 0.25, but got %.2f", m.insertedAmount)
	}
}

func TestChangeCalculation(t *testing.T) {
	cfg := &config.Config{
		Songs: []config.Song{
			{ID: 1, Name: "Test Song", Price: 1.25},
		},
	}

	m := NewModel(cfg)
	m.selectedSong = &cfg.Songs[0]
	m.insertedAmount = 1.50
	m.state = PlayingSong

	view := m.viewPlayingSong()
	expectedChange := "Your change: 0.25"
	if !strings.Contains(view, expectedChange) {
		t.Errorf("Expected view to contain '%s', but it didn't. View: %s", expectedChange, view)
	}
}

func TestSongSelection(t *testing.T) {
	cfg := &config.Config{
		Songs: []config.Song{
			{ID: 1, Name: "Song 1", Price: 1.00},
			{ID: 2, Name: "Song 2", Price: 1.50},
		},
	}

	m := NewModel(cfg)

	// Move down
	m.Update(tea.KeyMsg{Type: tea.KeyDown})
	if m.cursor != 1 {
		t.Errorf("Expected cursor to be at 1, but got %d", m.cursor)
	}

	// Move up
	m.Update(tea.KeyMsg{Type: tea.KeyUp})
	if m.cursor != 0 {
		t.Errorf("Expected cursor to be at 0, but got %d", m.cursor)
	}

	// Select song
	m.Update(tea.KeyMsg{Type: tea.KeyEnter})
	if m.state != AcceptingCoins {
		t.Errorf("Expected state to be AcceptingCoins, but got %d", m.state)
	}
	if m.selectedSong.Name != "Song 1" {
		t.Errorf("Expected selected song to be 'Song 1', but got '%s'", m.selectedSong.Name)
	}
}
