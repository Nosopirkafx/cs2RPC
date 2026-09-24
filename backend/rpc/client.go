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
		Name:    "FACEIT",
		Details: details,
		State:   state,
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

func Run(clientID string, in <-chan State, onStatus func(string)) {
	var handler *discordrpc.Client
	var last discordrpc.Activity
	connectionWarningShown := false

	for {
		if handler == nil {
			h, err := discordrpc.New(clientID)
			if err != nil {
				if !connectionWarningShown {
					log.Printf("Discord is not connected yet; start Discord desktop to enable Rich Presence: %v", err)
					if onStatus != nil {
						onStatus("DISCORD NOT RUNNING")
					}
					connectionWarningShown = true
				}
				time.Sleep(5 * time.Second)
				continue
			}
			handler = h
			if connectionWarningShown {
				log.Println("Discord RPC connection restored")
				if onStatus != nil {
					onStatus("CONNECTED")
				}
				connectionWarningShown = false
			} else {
				log.Println("Connected to Discord RPC")
				if onStatus != nil {
					onStatus("CONNECTED")
				}
			}
			if last.Name != "" {
				if err := handler.SetActivity(last); err != nil {
					log.Printf("restore activity failed: %v", err)
					if onStatus != nil {
						onStatus("CONNECTION LOST — RECONNECTING")
					}
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
					if onStatus != nil {
						onStatus("CONNECTION LOST — RECONNECTING")
					}
					handler = nil
					continue
				}
				last = discordrpc.Activity{}
				continue
			}
			a := s.ToActivity()
			if err := handler.SetActivity(a); err != nil {
				log.Printf("set activity failed: %v", err)
				if onStatus != nil {
					onStatus("CONNECTION LOST — RECONNECTING")
				}
				handler = nil
				continue
			}
			last = a
		case <-time.After(30 * time.Second):
			if last.Name != "" {
				if err := handler.SetActivity(last); err != nil {
					log.Printf("discord pipe lost, reconnecting: %v", err)
					if onStatus != nil {
						onStatus("CONNECTION LOST — RECONNECTING")
					}
					handler = nil
				}
			}
		}
	}
}
