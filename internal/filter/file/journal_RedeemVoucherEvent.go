package file

import "time"

type RedeemVoucherEvent struct {
	Amount           int     `mapstructure:"Amount,omitempty"`
	BrokerPercentage float64 `mapstructure:"BrokerPercentage,omitempty"`
	Faction          string  `mapstructure:"Faction,omitempty"`
	Factions         []struct {
		Amount  int    `mapstructure:"Amount,omitempty"`
		Faction string `mapstructure:"Faction,omitempty"`
	} `mapstructure:"Factions,omitempty"`
	Type      string    `mapstructure:"Type,omitempty"`
	Event     string    `mapstructure:"event,omitempty"`
	Timestamp time.Time `mapstructure:"timestamp,omitempty"`
}

func (h *handler) evRedeemVoucher(ev *RedeemVoucherEvent) {
	return
}
