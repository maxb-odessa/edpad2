package file

import "time"

type NpcCrewPaidWageEvent struct {
	Amount      int       `json:"Amount,omitempty"`
	NpcCrewID   int       `json:"NpcCrewId,omitempty"`
	NpcCrewName string    `json:"NpcCrewName,omitempty"`
	Event       string    `json:"event,omitempty"`
	Timestamp   time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evNpcCrewPaidWage(ev *NpcCrewPaidWageEvent) {
	return
}
