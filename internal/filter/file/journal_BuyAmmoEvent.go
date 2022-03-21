package file

import "time"

type BuyAmmoEvent struct {
	Cost      int       `json:"Cost,omitempty"`
	Event     string    `json:"event,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evBuyAmmo(ev *BuyAmmoEvent) {
	return
}
