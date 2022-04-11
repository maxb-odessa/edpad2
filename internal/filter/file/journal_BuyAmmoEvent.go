package file

import "time"

type BuyAmmoEvent struct {
	Cost      int       `mapstructure:"Cost,omitempty"`
	Event     string    `mapstructure:"event,omitempty"`
	Timestamp time.Time `mapstructure:"timestamp,omitempty"`
}

func (h *handler) evBuyAmmo(ev *BuyAmmoEvent) {
	return
}
