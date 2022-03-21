package file

import "time"

type MaterialsEvent struct {
	Encoded []struct {
		Count         int    `json:"Count,omitempty"`
		Name          string `json:"Name,omitempty"`
		NameLocalised string `json:"Name_Localised,omitempty"`
	} `json:"Encoded,omitempty"`
	Manufactured []struct {
		Count         int    `json:"Count,omitempty"`
		Name          string `json:"Name,omitempty"`
		NameLocalised string `json:"Name_Localised,omitempty"`
	} `json:"Manufactured,omitempty"`
	Raw []struct {
		Count int    `json:"Count,omitempty"`
		Name  string `json:"Name,omitempty"`
	} `json:"Raw,omitempty"`
	Event     string    `json:"event,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evMaterials(ev *MaterialsEvent) {
	return
}
