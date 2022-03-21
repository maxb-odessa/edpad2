package file

import "time"

type PayFinesEvent struct {
	AllFines  bool      `json:"AllFines,omitempty"`
	Amount    int       `json:"Amount,omitempty"`
	Faction   string    `json:"Faction,omitempty"`
	ShipID    int       `json:"ShipID,omitempty"`
	Event     string    `json:"event,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evPayFines(ev *PayFinesEvent) {
	return
}
