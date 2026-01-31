package main

import (
	"fmt"
	"log"
	"os"

	"jukebox/cli/config"
	"jukebox/cli/tui"

	tea "github.com/charmbracelet/bubbletea"
)

const analyticsAddr = "localhost:50051"

func main() {
	cfg, err := config.Load("./config/config.json")
	if err != nil {
		fmt.Printf("failed to load config: %v", err)
		os.Exit(1)
	}

	analyticsClient, err := NewAnalyticsClient(analyticsAddr)
	if err != nil {
		log.Printf("warning: could not connect to analytics service: %v", err)
	} else {
		defer analyticsClient.Close()
		log.Println("connected to analytics service")
	}

	model := tui.NewModel(cfg, analyticsClient)

	p := tea.NewProgram(model)
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
