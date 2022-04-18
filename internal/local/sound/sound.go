package sound

import (
	"edpad2/internal/router"

	"github.com/maxb-odessa/slog"
)

const (
	ALARM int = iota
	BEEP
	CLICK
	DROP
	NOTE
	TING
	TONE
	TWIT
	WARN
	WARP
)

type Track struct {
	Id     int
	Repeat int
	Stop   bool
}

type handler struct {
	// router data
	endpoint  router.Endpoint
	connector *router.Connector

	// priv data
	// snd dev state
}

func Connect(ep router.Endpoint) (router.Endpoint, *router.Connector) {
	h := new(handler)
	h.endpoint = ep
	h.connector = new(router.Connector)
	h.connector.FromRouterCh = make(chan *router.Message, 8) // wait for router messages on this chan
	h.connector.ToRouterCh = make(chan *router.Message, 8)   // send messages to the router into this chan
	h.connector.DoneCh = make(chan bool)                     // termination chan

	if err := h.init(); err != nil {
		slog.Err("endpoint '%s': failed: %s", ep, err)
		return ep, nil
	}

	go h.run()

	// all done, return router connector
	return ep, h.connector
}

func (h *handler) run() {

	// start messages processing
	for {

		select {

		case <-h.connector.DoneCh:
			close(h.connector.ToRouterCh)
			close(h.connector.FromRouterCh)
			return

		case d := <-h.connector.FromRouterCh:
			slog.Debug(9, "local sound got msg! %+v", d)
			snd := d.Data.(*Track)
			h.play(snd)

		} //select...

	} //for...

}
