package tui

import "jukebox/cli/config"

type JukeboxState int

const (
	SelectingSong JukeboxState = iota
	AcceptingCoins
	PlayingSong
)

type Model struct {
	state          JukeboxState
	config         *config.Config
	cursor         int
	selectedSong   *config.Song
	insertedAmount float64
	paymentInput   string
}

func NewModel(cfg *config.Config) *Model {
	return &Model{
		state:  SelectingSong,
		config: cfg,
	}
}
