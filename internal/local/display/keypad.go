package display

import (
	"edpad2/internal/router"
	"fmt"

	uin "github.com/bendahl/uinput"
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

func nameToCode(name string) uint32 {
	if kc, ok := keyCodes[name]; ok {
		return kc
	}
	return uin.KeySpace // TODO default?
}

func (h *handler) sendKeyEvent(name string, pressed bool) {

	msg := &pb.KbdMsg{
		Name: "keyboard",
		Msg: &pb.KbdMsg_Event{
			Event: &pb.KbdEvent{
				Obj: &pb.KbdEvent_Key_{
					Key: &pb.KbdEvent_Key{
						Code:    nameToCode(name),
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

var keyCodes = map[string]uint32{
	"esc":       uin.KeyEsc,
	"escape":    uin.KeyEsc,
	" ":         uin.KeySpace,
	"space":     uin.KeySpace,
	"enter":     uin.KeyEnter,
	"tab":       uin.KeyTab,
	"bs":        uin.KeyBackspace,
	"backspace": uin.KeyBackspace,

	"0": uin.Key0,
	"1": uin.Key1,
	"2": uin.Key2,
	"3": uin.Key3,
	"4": uin.Key4,
	"5": uin.Key5,
	"6": uin.Key6,
	"7": uin.Key7,
	"8": uin.Key8,
	"9": uin.Key9,

	"left":  uin.KeyLeft,
	"right": uin.KeyRight,
	"up":    uin.KeyUp,
	"down":  uin.KeyDown,

	"kp0": uin.KeyKp1,
	"kp1": uin.KeyKp1,
	"kp2": uin.KeyKp1,
	"kp3": uin.KeyKp1,
	"kp4": uin.KeyKp1,
	"kp5": uin.KeyKp1,
	"kp6": uin.KeyKp1,
	"kp7": uin.KeyKp1,
	"kp8": uin.KeyKp1,
	"kp9": uin.KeyKp1,
	"kp-": uin.KeyKpminus,
	"kp+": uin.KeyKpplus,
	"kp*": uin.KeyKpasterisk,
	"kp/": uin.KeyKpslash,

	"q": uin.KeyQ,
	"w": uin.KeyW,
	"e": uin.KeyE,
	"r": uin.KeyR,
	"t": uin.KeyT,
	"y": uin.KeyY,
	"u": uin.KeyU,
	"i": uin.KeyI,
	"o": uin.KeyO,
	"p": uin.KeyP,
	"[": uin.KeyLeftbrace,
	"]": uin.KeyRightbrace,
	"a": uin.KeyA,
	"s": uin.KeyS,
	"d": uin.KeyD,
	"f": uin.KeyF,
	"g": uin.KeyG,
	"h": uin.KeyH,
	"j": uin.KeyJ,
	"k": uin.KeyK,
	"l": uin.KeyL,
	";": uin.KeySemicolon,
	"'": uin.KeyApostrophe,
	"z": uin.KeyZ,
	"x": uin.KeyX,
	"c": uin.KeyC,
	"v": uin.KeyV,
	"b": uin.KeyB,
	"n": uin.KeyN,
	"m": uin.KeyM,
	",": uin.KeyComma,
	".": uin.KeyDot,
	"/": uin.KeySlash,
}
