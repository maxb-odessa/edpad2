package file

import "time"

type TransferMicroResourcesEvent struct {
	Transfers []struct {
		Category       string `mapstructure:"Category,omitempty"`
		Direction      string `mapstructure:"Direction,omitempty"`
		LockerNewCount int    `mapstructure:"LockerNewCount,omitempty"`
		LockerOldCount int    `mapstructure:"LockerOldCount,omitempty"`
		Name           string `mapstructure:"Name,omitempty"`
		NameLocalised  string `mapstructure:"Name_Localised,omitempty"`
	} `mapstructure:"Transfers,omitempty"`
	Event     string    `mapstructure:"event,omitempty"`
	Timestamp time.Time `mapstructure:"timestamp,omitempty"`
}

func (h *handler) evTransferMicroResources(ev *TransferMicroResourcesEvent) {
	return
}
