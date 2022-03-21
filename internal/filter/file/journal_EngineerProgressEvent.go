package file

import "time"

type EngineerProgressEvent struct {
	Engineer   string `json:"Engineer,omitempty"`
	EngineerID int    `json:"EngineerID,omitempty"`
	Engineers  []struct {
		Engineer     string `json:"Engineer,omitempty"`
		EngineerID   int    `json:"EngineerID,omitempty"`
		Progress     string `json:"Progress,omitempty"`
		Rank         int    `json:"Rank,omitempty"`
		RankProgress int    `json:"RankProgress,omitempty"`
	} `json:"Engineers,omitempty"`
	Progress  string    `json:"Progress,omitempty"`
	Rank      int       `json:"Rank,omitempty"`
	Event     string    `json:"event,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evEngineerProgress(ev *EngineerProgressEvent) {
	return
}
