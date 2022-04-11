package file

import "time"

type FriendsEvent struct {
	Name      string    `mapstructure:"Name,omitempty"`
	Status    string    `mapstructure:"Status,omitempty"`
	Event     string    `mapstructure:"event,omitempty"`
	Timestamp time.Time `mapstructure:"timestamp,omitempty"`
}

func (h *handler) evFriends(ev *FriendsEvent) {
	return
}
