package file

import "time"

type LoadoutEvent struct {
	CargoCapacity int `mapstructure:"CargoCapacity,omitempty"`
	FuelCapacity  struct {
		Main    float64 `mapstructure:"Main,omitempty"`
		Reserve float64 `mapstructure:"Reserve,omitempty"`
	} `mapstructure:"FuelCapacity,omitempty"`
	HullHealth   float64 `mapstructure:"HullHealth,omitempty"`
	HullValue    int     `mapstructure:"HullValue,omitempty"`
	MaxJumpRange float64 `mapstructure:"MaxJumpRange,omitempty"`
	Modules      []struct {
		AmmoInClip   int `mapstructure:"AmmoInClip,omitempty"`
		AmmoInHopper int `mapstructure:"AmmoInHopper,omitempty"`
		Engineering  *struct {
			BlueprintID                 int    `mapstructure:"BlueprintID,omitempty"`
			BlueprintName               string `mapstructure:"BlueprintName,omitempty"`
			Engineer                    string `mapstructure:"Engineer,omitempty"`
			EngineerID                  int    `mapstructure:"EngineerID,omitempty"`
			ExperimentalEffect          string `mapstructure:"ExperimentalEffect,omitempty"`
			ExperimentalEffectLocalised string `mapstructure:"ExperimentalEffect_Localised,omitempty"`
			Level                       int    `mapstructure:"Level,omitempty"`
			Modifiers                   []struct {
				Label         string  `mapstructure:"Label,omitempty"`
				LessIsGood    int     `mapstructure:"LessIsGood,omitempty"`
				OriginalValue float64 `mapstructure:"OriginalValue,omitempty"`
				Value         float64 `mapstructure:"Value,omitempty"`
			} `mapstructure:"Modifiers,omitempty"`
			Quality float64 `mapstructure:"Quality,omitempty"`
		} `mapstructure:"Engineering,omitempty"`
		Health   float64 `mapstructure:"Health,omitempty"`
		Item     string  `mapstructure:"Item,omitempty"`
		On       bool    `mapstructure:"On,omitempty"`
		Priority int     `mapstructure:"Priority,omitempty"`
		Slot     string  `mapstructure:"Slot,omitempty"`
		Value    int     `mapstructure:"Value,omitempty"`
	} `mapstructure:"Modules,omitempty"`
	ModulesValue int       `mapstructure:"ModulesValue,omitempty"`
	Rebuy        int       `mapstructure:"Rebuy,omitempty"`
	Ship         string    `mapstructure:"Ship,omitempty"`
	ShipID       int       `mapstructure:"ShipID,omitempty"`
	ShipIdent    string    `mapstructure:"ShipIdent,omitempty"`
	ShipName     string    `mapstructure:"ShipName,omitempty"`
	UnladenMass  float64   `mapstructure:"UnladenMass,omitempty"`
	Event        string    `mapstructure:"event,omitempty"`
	Timestamp    time.Time `mapstructure:"timestamp,omitempty"`
}

func (h *handler) evLoadout(ev *LoadoutEvent) {
	return
}
