package file

import "time"

type PromotionEvent struct {
	Empire       int       `json:"Empire,omitempty"`
	Exobiologist int       `json:"Exobiologist,omitempty"`
	Explore      int       `json:"Explore,omitempty"`
	Federation   int       `json:"Federation,omitempty"`
	Trade        int       `json:"Trade,omitempty"`
	Event        string    `json:"event,omitempty"`
	Timestamp    time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evPromotion(ev *PromotionEvent) {
	return
}
