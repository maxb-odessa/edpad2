package display

import (
	"fmt"

	"github.com/gotk3/gotk3/gtk"
)

func (h *handler) keypad() error {

	obj, err := h.gtkBuilder.GetObject("keypad")
	if err != nil {
		return err
	}

	kp, ok := obj.(*gtk.Popover)
	if !ok {
		return fmt.Errorf("endpoint '%s': 'keypad' is not a *gtk.Popover obj", h.endpoint)
	}

	kp.Popdown()

	toggleKeypad := func(s interface{}, e interface{}) {
		kp.Popup()
	}

	signals := map[string]interface{}{
		"toggle_keypad": toggleKeypad,
		"onPress":       dummy,
		"onRelease":     dummy,
		"onToggle":      dummy,
	}

	h.gtkBuilder.ConnectSignals(signals)

	return nil
}

func dummy() {}
