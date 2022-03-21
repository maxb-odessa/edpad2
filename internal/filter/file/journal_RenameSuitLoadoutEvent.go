package file

import "time"

type RenameSuitLoadoutEvent struct {
	LoadoutID         int       `json:"LoadoutID,omitempty"`
	LoadoutName       string    `json:"LoadoutName,omitempty"`
	SuitID            int       `json:"SuitID,omitempty"`
	SuitName          string    `json:"SuitName,omitempty"`
	SuitNameLocalised string    `json:"SuitName_Localised,omitempty"`
	Event             string    `json:"event,omitempty"`
	Timestamp         time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evRenameSuitLoadout(ev *RenameSuitLoadoutEvent) {
	return
}
