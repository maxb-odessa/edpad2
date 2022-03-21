package file

import "time"

type TransferMicroResourcesEvent struct {
	Transfers []struct {
		Category       string `json:"Category,omitempty"`
		Direction      string `json:"Direction,omitempty"`
		LockerNewCount int    `json:"LockerNewCount,omitempty"`
		LockerOldCount int    `json:"LockerOldCount,omitempty"`
		Name           string `json:"Name,omitempty"`
		NameLocalised  string `json:"Name_Localised,omitempty"`
	} `json:"Transfers,omitempty"`
	Event     string    `json:"event,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evTransferMicroResources(ev *TransferMicroResourcesEvent) {
	return
}
