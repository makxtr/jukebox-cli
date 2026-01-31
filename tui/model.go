package tui

import "jukebox/cli/config"

type JukeboxState int

const (
	SelectingSong JukeboxState = iota
	AcceptingCoins
	PlayingSong
)

type AnalyticsLogger interface {
	LogPlayback(trackID int, amountPaid float64) error
}

type Model struct {
	state          JukeboxState
	config         *config.Config
	cursor         int
	selectedSong   *config.Song
	insertedAmount float64
	paymentInput   string
	analytics      AnalyticsLogger
}

func NewModel(cfg *config.Config, analytics AnalyticsLogger) *Model {
	return &Model{
		state:     SelectingSong,
		config:    cfg,
		analytics: analytics,
	}
}
