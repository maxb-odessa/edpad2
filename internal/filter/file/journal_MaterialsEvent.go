package file

import "time"

type MaterialsEvent struct {
	Encoded []struct {
		Count         int    `mapstructure:"Count,omitempty"`
		Name          string `mapstructure:"Name,omitempty"`
		NameLocalised string `mapstructure:"Name_Localised,omitempty"`
	} `mapstructure:"Encoded,omitempty"`
	Manufactured []struct {
		Count         int    `mapstructure:"Count,omitempty"`
		Name          string `mapstructure:"Name,omitempty"`
		NameLocalised string `mapstructure:"Name_Localised,omitempty"`
	} `mapstructure:"Manufactured,omitempty"`
	Raw []struct {
		Count int    `mapstructure:"Count,omitempty"`
		Name  string `mapstructure:"Name,omitempty"`
	} `mapstructure:"Raw,omitempty"`
	Event     string    `mapstructure:"event,omitempty"`
	Timestamp time.Time `mapstructure:"timestamp,omitempty"`
}

func (h *handler) evMaterials(ev *MaterialsEvent) {
	return
}
