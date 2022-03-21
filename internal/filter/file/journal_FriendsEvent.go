package file

import "time"

type FriendsEvent struct {
	Name      string    `json:"Name,omitempty"`
	Status    string    `json:"Status,omitempty"`
	Event     string    `json:"event,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evFriends(ev *FriendsEvent) {
	return
}
