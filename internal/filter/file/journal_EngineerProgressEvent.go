package file

import "time"

type EngineerProgressEvent struct {
	Engineer   string `mapstructure:"Engineer,omitempty"`
	EngineerID int    `mapstructure:"EngineerID,omitempty"`
	Engineers  []struct {
		Engineer     string `mapstructure:"Engineer,omitempty"`
		EngineerID   int    `mapstructure:"EngineerID,omitempty"`
		Progress     string `mapstructure:"Progress,omitempty"`
		Rank         int    `mapstructure:"Rank,omitempty"`
		RankProgress int    `mapstructure:"RankProgress,omitempty"`
	} `mapstructure:"Engineers,omitempty"`
	Progress  string    `mapstructure:"Progress,omitempty"`
	Rank      int       `mapstructure:"Rank,omitempty"`
	Event     string    `mapstructure:"event,omitempty"`
	Timestamp time.Time `mapstructure:"timestamp,omitempty"`
}

func (h *handler) evEngineerProgress(ev *EngineerProgressEvent) {
	return
}
