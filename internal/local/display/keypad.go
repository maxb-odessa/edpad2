package display

import (
	"edpad2/internal/router"
	"fmt"

	"github.com/gotk3/gotk3/gtk"

	pb "github.com/maxb-odessa/gamenode/pkg/gamenodepb"
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

	onButtonDownFunc := func(s interface{}) {
		if name, err := s.(*gtk.Button).GetName(); err == nil {
			h.sendKeyEvent(name, true)
		}
	}

	onButtonUpFunc := func(s interface{}) {
		if name, err := s.(*gtk.Button).GetName(); err == nil {
			h.sendKeyEvent(name, false)
		}
	}

	onButtonToggleFunc := func(s interface{}) {
		if name, err := s.(*gtk.ToggleButton).GetName(); err == nil {
			h.sendKeyEvent(name, true)
			h.sendKeyEvent(name, false)
		}
	}

	signals := map[string]interface{}{
		"toggle_keypad": toggleKeypad,
		"onPress":       onButtonDownFunc,
		"onRelease":     onButtonUpFunc,
		"onToggle":      onButtonToggleFunc,
	}

	h.gtkBuilder.ConnectSignals(signals)

	return nil
}

func (h *handler) sendKeyEvent(name string, pressed bool) {

	msg := &pb.KbdMsg{
		Name: "keyboard",
		Msg: &pb.KbdMsg_Event{
			Event: &pb.KbdEvent{
				Obj: &pb.KbdEvent_Key_{
					Key: &pb.KbdEvent_Key{
						Id:      name,
						Pressed: pressed,
					},
				},
			},
		},
	}

	h.connector.ToRouterCh <- &router.Message{
		Dst:  router.NetKeyboard,
		Data: msg,
	}
}
