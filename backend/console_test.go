package main

import (
	"bytes"
	"strings"
	"testing"
	"time"

	"faceit-discord-rpc/backend/rpc"
)

func TestConsoleDashboardShowsLogoAndMatchData(t *testing.T) {
	var out bytes.Buffer
	mapName := "anubis"
	elo := 1720
	phase := "LIVE"
	teamScore := &rpc.Score{A: 13, B: 9}
	matchStart := time.Now().Add(-time.Minute).Unix()
	ui := newConsoleUI(&out, defaultArt(), false, 90)
	ui.SetDiscordStatus("CONNECTED")
	ui.SetMatchState(rpc.State{
		Status:     "match",
		Map:        &mapName,
		Elo:        &elo,
		Score:      teamScore,
		Phase:      &phase,
		MatchStart: &matchStart,
	}, time.Now())

	for _, expected := range []string{
		"___  ___  ___  _____",
		"LIVE FACEIT DATA",
		"CONNECTED",
		"LIVE MATCH",
		"ANUBIS",
		"1720",
		"13  :  9",
		"LIVE",
		"MATCH TIME",
		"UPDATED",
	} {
		if !strings.Contains(out.String(), expected) {
			t.Errorf("dashboard output does not contain %q", expected)
		}
	}
}
