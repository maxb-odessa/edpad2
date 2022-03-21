package file

import "time"

type RankEvent struct {
	Cqc          int       `json:"CQC,omitempty"`
	Combat       int       `json:"Combat,omitempty"`
	Empire       int       `json:"Empire,omitempty"`
	Exobiologist int       `json:"Exobiologist,omitempty"`
	Explore      int       `json:"Explore,omitempty"`
	Federation   int       `json:"Federation,omitempty"`
	Soldier      int       `json:"Soldier,omitempty"`
	Trade        int       `json:"Trade,omitempty"`
	Event        string    `json:"event,omitempty"`
	Timestamp    time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evRank(ev *RankEvent) {
	return
}
