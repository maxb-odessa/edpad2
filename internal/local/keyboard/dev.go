package keyboard

import (
	"github.com/bendahl/uinput"
)

type dev struct {
	confScope string
	vk        uinput.Keyboard
}

func newDev(confScope string) *dev {
	return &dev{}
}

func (h *dev) run() error {

	vk, err := uinput.CreateKeyboard("/dev/uinput", []byte("EDPad2 Virtual Keyboard"))
	if err != nil {
		return err
	}

	h.vk = vk

	return nil
}

func (h *dev) finish() {
	h.vk.Close()
}

func (h *dev) writeKey(key *Key) error {
	if key.Pressed {
		return h.vk.KeyDown(int(key.Code))
	} else {
		return h.vk.KeyUp(int(key.Code))
	}
}
