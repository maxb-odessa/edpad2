package file

import "time"

type CommunityGoalJoinEvent struct {
	Cgid      int       `json:"CGID,omitempty"`
	Name      string    `json:"Name,omitempty"`
	System    string    `json:"System,omitempty"`
	Event     string    `json:"event,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evCommunityGoalJoin(ev *CommunityGoalJoinEvent) {
	return
}
