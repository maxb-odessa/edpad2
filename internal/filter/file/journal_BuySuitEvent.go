package file

import "time"

type BuySuitEvent struct {
	Name          string    `mapstructure:"Name,omitempty"`
	NameLocalised string    `mapstructure:"Name_Localised,omitempty"`
	Price         int       `mapstructure:"Price,omitempty"`
	SuitID        int       `mapstructure:"SuitID,omitempty"`
	SuitMods      []string  `mapstructure:"SuitMods,omitempty"`
	Event         string    `mapstructure:"event,omitempty"`
	Timestamp     time.Time `mapstructure:"timestamp,omitempty"`
}

func (h *handler) evBuySuit(ev *BuySuitEvent) {
	return
}
