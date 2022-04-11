package file

import "time"

type SetUserShipNameEvent struct {
	Ship         string    `mapstructure:"Ship,omitempty"`
	ShipID       int       `mapstructure:"ShipID,omitempty"`
	UserShipID   string    `mapstructure:"UserShipId,omitempty"`
	UserShipName string    `mapstructure:"UserShipName,omitempty"`
	Event        string    `mapstructure:"event,omitempty"`
	Timestamp    time.Time `mapstructure:"timestamp,omitempty"`
}

func (h *handler) evSetUserShipName(ev *SetUserShipNameEvent) {
	return
}
