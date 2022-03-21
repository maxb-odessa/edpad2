package file

import "time"

type FileheaderEvent struct {
	Odyssey     bool      `json:"Odyssey,omitempty"`
	Build       string    `json:"build,omitempty"`
	Event       string    `json:"event,omitempty"`
	Gameversion string    `json:"gameversion,omitempty"`
	Language    string    `json:"language,omitempty"`
	Part        int       `json:"part,omitempty"`
	Timestamp   time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evFileheader(ev *FileheaderEvent) {
	return
}
