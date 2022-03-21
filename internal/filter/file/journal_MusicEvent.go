package file

import "time"

type MusicEvent struct {
	MusicTrack string    `json:"MusicTrack,omitempty"`
	Event      string    `json:"event,omitempty"`
	Timestamp  time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evMusic(ev *MusicEvent) {
	return
}
