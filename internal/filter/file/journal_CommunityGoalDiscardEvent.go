package file

import "time"

type CommunityGoalDiscardEvent struct {
	Cgid      int       `mapstructure:"CGID,omitempty"`
	Name      string    `mapstructure:"Name,omitempty"`
	System    string    `mapstructure:"System,omitempty"`
	Event     string    `mapstructure:"event,omitempty"`
	Timestamp time.Time `mapstructure:"timestamp,omitempty"`
}

func (h *handler) evCommunityGoalDiscard(ev *CommunityGoalDiscardEvent) {
	return
}
