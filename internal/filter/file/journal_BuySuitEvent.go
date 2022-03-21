package file

import "time"

type BuySuitEvent struct {
	Name          string    `json:"Name,omitempty"`
	NameLocalised string    `json:"Name_Localised,omitempty"`
	Price         int       `json:"Price,omitempty"`
	SuitID        int       `json:"SuitID,omitempty"`
	SuitMods      []string  `json:"SuitMods,omitempty"`
	Event         string    `json:"event,omitempty"`
	Timestamp     time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evBuySuit(ev *BuySuitEvent) {
	return
}
