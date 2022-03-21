package file

import "time"

type LoadoutEvent struct {
	CargoCapacity int `json:"CargoCapacity,omitempty"`
	FuelCapacity  struct {
		Main    float64 `json:"Main,omitempty"`
		Reserve float64 `json:"Reserve,omitempty"`
	} `json:"FuelCapacity,omitempty"`
	HullHealth   float64 `json:"HullHealth,omitempty"`
	HullValue    int     `json:"HullValue,omitempty"`
	MaxJumpRange float64 `json:"MaxJumpRange,omitempty"`
	Modules      []struct {
		AmmoInClip   int `json:"AmmoInClip,omitempty"`
		AmmoInHopper int `json:"AmmoInHopper,omitempty"`
		Engineering  *struct {
			BlueprintID                 int    `json:"BlueprintID,omitempty"`
			BlueprintName               string `json:"BlueprintName,omitempty"`
			Engineer                    string `json:"Engineer,omitempty"`
			EngineerID                  int    `json:"EngineerID,omitempty"`
			ExperimentalEffect          string `json:"ExperimentalEffect,omitempty"`
			ExperimentalEffectLocalised string `json:"ExperimentalEffect_Localised,omitempty"`
			Level                       int    `json:"Level,omitempty"`
			Modifiers                   []struct {
				Label         string  `json:"Label,omitempty"`
				LessIsGood    int     `json:"LessIsGood,omitempty"`
				OriginalValue float64 `json:"OriginalValue,omitempty"`
				Value         float64 `json:"Value,omitempty"`
			} `json:"Modifiers,omitempty"`
			Quality float64 `json:"Quality,omitempty"`
		} `json:"Engineering,omitempty"`
		Health   float64 `json:"Health,omitempty"`
		Item     string  `json:"Item,omitempty"`
		On       bool    `json:"On,omitempty"`
		Priority int     `json:"Priority,omitempty"`
		Slot     string  `json:"Slot,omitempty"`
		Value    int     `json:"Value,omitempty"`
	} `json:"Modules,omitempty"`
	ModulesValue int       `json:"ModulesValue,omitempty"`
	Rebuy        int       `json:"Rebuy,omitempty"`
	Ship         string    `json:"Ship,omitempty"`
	ShipID       int       `json:"ShipID,omitempty"`
	ShipIdent    string    `json:"ShipIdent,omitempty"`
	ShipName     string    `json:"ShipName,omitempty"`
	UnladenMass  float64   `json:"UnladenMass,omitempty"`
	Event        string    `json:"event,omitempty"`
	Timestamp    time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evLoadout(ev *LoadoutEvent) {
	return
}
