package file

import "time"

type SetUserShipNameEvent struct {
	Ship         string    `json:"Ship,omitempty"`
	ShipID       int       `json:"ShipID,omitempty"`
	UserShipID   string    `json:"UserShipId,omitempty"`
	UserShipName string    `json:"UserShipName,omitempty"`
	Event        string    `json:"event,omitempty"`
	Timestamp    time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evSetUserShipName(ev *SetUserShipNameEvent) {
	return
}
