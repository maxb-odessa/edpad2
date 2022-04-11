package file

import "time"

type PayFinesEvent struct {
	AllFines  bool      `mapstructure:"AllFines,omitempty"`
	Amount    int       `mapstructure:"Amount,omitempty"`
	Faction   string    `mapstructure:"Faction,omitempty"`
	ShipID    int       `mapstructure:"ShipID,omitempty"`
	Event     string    `mapstructure:"event,omitempty"`
	Timestamp time.Time `mapstructure:"timestamp,omitempty"`
}

func (h *handler) evPayFines(ev *PayFinesEvent) {
	return
}
