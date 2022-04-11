package file

import "time"

type DatalinkVoucherEvent struct {
	PayeeFaction  string    `mapstructure:"PayeeFaction,omitempty"`
	Reward        int       `mapstructure:"Reward,omitempty"`
	VictimFaction string    `mapstructure:"VictimFaction,omitempty"`
	Event         string    `mapstructure:"event,omitempty"`
	Timestamp     time.Time `mapstructure:"timestamp,omitempty"`
}

func (h *handler) evDatalinkVoucher(ev *DatalinkVoucherEvent) {
	return
}
