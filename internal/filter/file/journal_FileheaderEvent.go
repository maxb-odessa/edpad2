package file

import "time"

type FileheaderEvent struct {
	Odyssey     bool      `mapstructure:"Odyssey,omitempty"`
	Build       string    `mapstructure:"build,omitempty"`
	Event       string    `mapstructure:"event,omitempty"`
	Gameversion string    `mapstructure:"gameversion,omitempty"`
	Language    string    `mapstructure:"language,omitempty"`
	Part        int       `mapstructure:"part,omitempty"`
	Timestamp   time.Time `mapstructure:"timestamp,omitempty"`
}

func (h *handler) evFileheader(ev *FileheaderEvent) {
	return
}
