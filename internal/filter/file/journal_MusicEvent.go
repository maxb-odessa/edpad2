package file

import "time"

type MusicEvent struct {
	MusicTrack string    `mapstructure:"MusicTrack,omitempty"`
	Event      string    `mapstructure:"event,omitempty"`
	Timestamp  time.Time `mapstructure:"timestamp,omitempty"`
}

func (h *handler) evMusic(ev *MusicEvent) {
	return
}
