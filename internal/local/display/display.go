package display

import (
	"github.com/maxb-odessa/slog"

	"edpad2/internal/router"
)

type handler struct {
	endpoint  router.Endpoint
	connector *router.Connector
}

func Connect(ep router.Endpoint) (router.Endpoint, *router.Connector) {
	h := new(handler)
	h.endpoint = ep
	h.connector = new(router.Connector)
	h.connector.FromRouterCh = make(chan *router.Message) // wait for router messages on this chan
	h.connector.ToRouterCh = make(chan *router.Message)   // send messages to the router into this chan
	h.connector.DoneCh = make(chan bool)                  // termination chan

	if err := h.Init(); err != nil {
		slog.Err("endpoint '%s': init failed: %s", ep, err)
		return ep, nil
	}

	// all done, return router connector
	return ep, h.connector
}

func (h *handler) Init() error {

	// do configs

	// go reader()

	// go sender()

	return nil
}

/*
func (h *handler) runReader() {

	for {

		msg, err := h.reader()

		// remote closed
		if err == io.EOF {
			return
		}

		// cancelled, other error
		if err != nil {
			//slog.Debug(5, "endpoint '%s' failed to receive: %s", h.endpoint, err)
			return
		}

		slog.Debug(9, "endpoint '%s': got msg: '%+v'", h.endpoint, msg)

		h.connector.ToRouterCh <- &router.Message{Dst: router.FilterFile, Data: msg}

	} //for

}

func (h *handler) runSender() {

	for {

		select {
		case data, ok := <-h.connector.FromRouterCh:
			if !ok {
				return // must exit
			}
			msg := data.Data
			if err := h.sender(msg); err != nil {
				slog.Warn("endpoint '%s': failed to send: %s", h.endpoint, err)
				return
			}
		}

	} //for
}
*/
