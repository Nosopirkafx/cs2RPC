package rpc

import (
	"errors"
	"strconv"
)

type Score struct {
	A int `json:"team_a"`
	B int `json:"team_b"`
}

type State struct {
	Status     string  `json:"status"`
	Map        *string `json:"map"`
	Elo        *int    `json:"elo"`
	Score      *Score  `json:"score"`
	Phase      *string `json:"phase"`
	MatchStart *int64  `json:"match_start"`
	MatchID    *string `json:"match_id"`
}

func (s *State) Validate() error {
	switch s.Status {
	case "match", "queue", "idle":
	default:
		return errors.New("invalid status")
	}
	if s.Status == "idle" {
		return nil
	}
	if s.Score != nil && (s.Score.A < 0 || s.Score.B < 0) {
		return errors.New("invalid score")
	}
	return nil
}

func itoa(v int) string { return strconv.Itoa(v) }
