package file

import "time"

type RankEvent struct {
	Cqc          int       `mapstructure:"CQC,omitempty"`
	Combat       int       `mapstructure:"Combat,omitempty"`
	Empire       int       `mapstructure:"Empire,omitempty"`
	Exobiologist int       `mapstructure:"Exobiologist,omitempty"`
	Explore      int       `mapstructure:"Explore,omitempty"`
	Federation   int       `mapstructure:"Federation,omitempty"`
	Soldier      int       `mapstructure:"Soldier,omitempty"`
	Trade        int       `mapstructure:"Trade,omitempty"`
	Event        string    `mapstructure:"event,omitempty"`
	Timestamp    time.Time `mapstructure:"timestamp,omitempty"`
}

func (h *handler) evRank(ev *RankEvent) {
	return
}
