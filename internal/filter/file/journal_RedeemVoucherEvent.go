package file

import "time"

type RedeemVoucherEvent struct {
	Amount           int     `json:"Amount,omitempty"`
	BrokerPercentage float64 `json:"BrokerPercentage,omitempty"`
	Faction          string  `json:"Faction,omitempty"`
	Factions         []struct {
		Amount  int    `json:"Amount,omitempty"`
		Faction string `json:"Faction,omitempty"`
	} `json:"Factions,omitempty"`
	Type      string    `json:"Type,omitempty"`
	Event     string    `json:"event,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evRedeemVoucher(ev *RedeemVoucherEvent) {
	return
}
