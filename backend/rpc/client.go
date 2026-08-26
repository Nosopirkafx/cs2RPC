package rpc

import (
	"log"
	"time"

	discordrpc "github.com/rikkuness/discord-rpc"
)

const (
	largeImageKey  = "cs2"
	largeImageText = "Counter-Strike 2"
	smallImageKey  = "faciet"
	smallImageText = "FACEIT Match"
)

func (s *State) ToActivity() discordrpc.Activity {
	details := "Map: Searching..."
	if s.Map != nil && *s.Map != "" && *s.Map != "null" {
		elo := "?"
		if s.Elo != nil {
			elo = itoa(*s.Elo)
		}
		details = "Map: " + *s.Map + " | ELO: " + elo
	}

	state := "In Menu"
	switch {
	case s.Score != nil:
		state = "Score: " + itoa(s.Score.A) + " : " + itoa(s.Score.B)
	case s.Phase != nil:
		state = *s.Phase
	case s.Status == "queue":
		state = "In Queue"
	}

	act := discordrpc.Activity{
		Name:  "FACEIT",
		Details: details,
		State: state,
		Assets: &discordrpc.Assets{
			LargeImage: largeImageKey,
			LargeText:  largeImageText,
			SmallImage: smallImageKey,
			SmallText:  smallImageText,
		},
	}

	if s.MatchID != nil && *s.MatchID != "" {
		act.Party = &discordrpc.Party{ID: *s.MatchID, Size: []int{5, 5}}
	}
	if s.MatchStart != nil && *s.MatchStart > 0 {
		t := time.Unix(*s.MatchStart, 0)
		act.Timestamps = &discordrpc.Timestamps{Start: &discordrpc.Epoch{Time: t}}
	}
	return act
}

func Run(clientID string, in <-chan State) {
	var handler *discordrpc.Client
	var last discordrpc.Activity

	for {
		if handler == nil {
			h, err := discordrpc.New(clientID)
			if err != nil {
				log.Printf("discord login failed (is discord running?): %v", err)
				time.Sleep(2 * time.Second)
				continue
			}
			handler = h
			if last.Name != "" {
				if err := handler.SetActivity(last); err != nil {
					log.Printf("restore activity failed: %v", err)
					handler = nil
					continue
				}
			}
		}

		select {
		case s := <-in:
			if s.Status == "idle" {
				if err := handler.SetActivity(discordrpc.Activity{Name: "FACEIT"}); err != nil {
					log.Printf("clear activity failed: %v", err)
					handler = nil
					continue
				}
				last = discordrpc.Activity{}
				continue
			}
			a := s.ToActivity()
			if err := handler.SetActivity(a); err != nil {
				log.Printf("set activity failed: %v", err)
				handler = nil
				continue
			}
			last = a
		case <-time.After(30 * time.Second):
			if last.Name != "" {
				if err := handler.SetActivity(last); err != nil {
					log.Printf("discord pipe lost, reconnecting: %v", err)
					handler = nil
				}
			}
		}
	}
}
