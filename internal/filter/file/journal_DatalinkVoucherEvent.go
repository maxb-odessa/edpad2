package file

import "time"

type DatalinkVoucherEvent struct {
	PayeeFaction  string    `json:"PayeeFaction,omitempty"`
	Reward        int       `json:"Reward,omitempty"`
	VictimFaction string    `json:"VictimFaction,omitempty"`
	Event         string    `json:"event,omitempty"`
	Timestamp     time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evDatalinkVoucher(ev *DatalinkVoucherEvent) {
	return
}
