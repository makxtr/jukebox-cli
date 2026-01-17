package tui

import "jukebox/cli/config"

// JukeboxState represents the current state of the UI.
type JukeboxState int

const (
	SelectingSong JukeboxState = iota
	AcceptingCoins
	PlayingSong
)

// Model is the core struct for our UI.
type Model struct {
	state         JukeboxState
	config        *config.Config
	cursor        int
	selectedSong  *config.Song
	insertedAmount float64
	paymentInput  string
}

// NewModel creates a new model with its initial state.
func NewModel(cfg *config.Config) *Model {
	return &Model{
		state:  SelectingSong,
		config: cfg,
	}
}
