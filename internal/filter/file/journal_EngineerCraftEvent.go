package file

import "time"

type EngineerCraftEvent struct {
	ApplyExperimentalEffect     string `mapstructure:"ApplyExperimentalEffect,omitempty"`
	BlueprintID                 int    `mapstructure:"BlueprintID,omitempty"`
	BlueprintName               string `mapstructure:"BlueprintName,omitempty"`
	Engineer                    string `mapstructure:"Engineer,omitempty"`
	EngineerID                  int    `mapstructure:"EngineerID,omitempty"`
	ExperimentalEffect          string `mapstructure:"ExperimentalEffect,omitempty"`
	ExperimentalEffectLocalised string `mapstructure:"ExperimentalEffect_Localised,omitempty"`
	Ingredients                 []struct {
		Count         int    `mapstructure:"Count,omitempty"`
		Name          string `mapstructure:"Name,omitempty"`
		NameLocalised string `mapstructure:"Name_Localised,omitempty"`
	} `mapstructure:"Ingredients,omitempty"`
	Level     int `mapstructure:"Level,omitempty"`
	Modifiers []struct {
		Label         string  `mapstructure:"Label,omitempty"`
		LessIsGood    int     `mapstructure:"LessIsGood,omitempty"`
		OriginalValue float64 `mapstructure:"OriginalValue,omitempty"`
		Value         float64 `mapstructure:"Value,omitempty"`
	} `mapstructure:"Modifiers,omitempty"`
	Module    string    `mapstructure:"Module,omitempty"`
	Quality   float64   `mapstructure:"Quality,omitempty"`
	Slot      string    `mapstructure:"Slot,omitempty"`
	Event     string    `mapstructure:"event,omitempty"`
	Timestamp time.Time `mapstructure:"timestamp,omitempty"`
}

func (h *handler) evEngineerCraft(ev *EngineerCraftEvent) {
	return
}
