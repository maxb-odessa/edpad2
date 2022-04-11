package file

import "time"

type RenameSuitLoadoutEvent struct {
	LoadoutID         int       `mapstructure:"LoadoutID,omitempty"`
	LoadoutName       string    `mapstructure:"LoadoutName,omitempty"`
	SuitID            int       `mapstructure:"SuitID,omitempty"`
	SuitName          string    `mapstructure:"SuitName,omitempty"`
	SuitNameLocalised string    `mapstructure:"SuitName_Localised,omitempty"`
	Event             string    `mapstructure:"event,omitempty"`
	Timestamp         time.Time `mapstructure:"timestamp,omitempty"`
}

func (h *handler) evRenameSuitLoadout(ev *RenameSuitLoadoutEvent) {
	return
}
