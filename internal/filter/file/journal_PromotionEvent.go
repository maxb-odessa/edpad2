package file

import "time"

type PromotionEvent struct {
	Empire       int       `mapstructure:"Empire,omitempty"`
	Exobiologist int       `mapstructure:"Exobiologist,omitempty"`
	Explore      int       `mapstructure:"Explore,omitempty"`
	Federation   int       `mapstructure:"Federation,omitempty"`
	Trade        int       `mapstructure:"Trade,omitempty"`
	Event        string    `mapstructure:"event,omitempty"`
	Timestamp    time.Time `mapstructure:"timestamp,omitempty"`
}

func (h *handler) evPromotion(ev *PromotionEvent) {
	return
}
