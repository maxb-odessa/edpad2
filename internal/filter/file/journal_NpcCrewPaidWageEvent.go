package file

import "time"

type NpcCrewPaidWageEvent struct {
	Amount      int       `mapstructure:"Amount,omitempty"`
	NpcCrewID   int       `mapstructure:"NpcCrewId,omitempty"`
	NpcCrewName string    `mapstructure:"NpcCrewName,omitempty"`
	Event       string    `mapstructure:"event,omitempty"`
	Timestamp   time.Time `mapstructure:"timestamp,omitempty"`
}

func (h *handler) evNpcCrewPaidWage(ev *NpcCrewPaidWageEvent) {
	return
}
