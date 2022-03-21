package file

import "time"

type EngineerCraftEvent struct {
	ApplyExperimentalEffect     string `json:"ApplyExperimentalEffect,omitempty"`
	BlueprintID                 int    `json:"BlueprintID,omitempty"`
	BlueprintName               string `json:"BlueprintName,omitempty"`
	Engineer                    string `json:"Engineer,omitempty"`
	EngineerID                  int    `json:"EngineerID,omitempty"`
	ExperimentalEffect          string `json:"ExperimentalEffect,omitempty"`
	ExperimentalEffectLocalised string `json:"ExperimentalEffect_Localised,omitempty"`
	Ingredients                 []struct {
		Count         int    `json:"Count,omitempty"`
		Name          string `json:"Name,omitempty"`
		NameLocalised string `json:"Name_Localised,omitempty"`
	} `json:"Ingredients,omitempty"`
	Level     int `json:"Level,omitempty"`
	Modifiers []struct {
		Label         string  `json:"Label,omitempty"`
		LessIsGood    int     `json:"LessIsGood,omitempty"`
		OriginalValue float64 `json:"OriginalValue,omitempty"`
		Value         float64 `json:"Value,omitempty"`
	} `json:"Modifiers,omitempty"`
	Module    string    `json:"Module,omitempty"`
	Quality   float64   `json:"Quality,omitempty"`
	Slot      string    `json:"Slot,omitempty"`
	Event     string    `json:"event,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evEngineerCraft(ev *EngineerCraftEvent) {
	return
}
