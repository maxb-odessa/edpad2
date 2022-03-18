package file

import (
	"fmt"
	"time"

	"github.com/maxb-odessa/slog"

	pb "github.com/maxb-odessa/gamenode/pkg/gamenodepb"

	"edpad2/internal/local/display"
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

	go h.run()

	// all done, return router connector
	return ep, h.connector
}

func (h *handler) run() {

	// TEST
	go func() {
		for {
			time.Sleep(time.Second * 1)
			h.connector.ToRouterCh <- &router.Message{
				Dst: router.LocalDisplay,
				Data: &display.Text{
					ViewPort: display.VP_SYS,
					Text:     fmt.Sprintf("<i>TEST SYSTEMO ITALICO! %d</i>", time.Now()),
				},
			}
		}
	}()

	// TEST
	go func() {
		for {
			time.Sleep(time.Second * 1)
			h.connector.ToRouterCh <- &router.Message{
				Dst: router.LocalDisplay,
				Data: &display.Text{
					ViewPort: display.VP_DSS,
					Text:     fmt.Sprintf("<b>TEST DSS BOLDO! %d</b>", time.Now()),
				},
			}
		}
	}()

	// configure supported processors
	processors := map[string]func(*pb.FileEvent){
		"ed journal": func(m *pb.FileEvent) { h.processJournalMsg(m) },
	}

	// start messages processing
	for {

		select {

		case <-h.connector.DoneCh:

			close(h.connector.ToRouterCh)
			close(h.connector.FromRouterCh)
			return

		case m := <-h.connector.FromRouterCh:

			slog.Debug(9, "file filter got msg! %+v", m)

			msg := m.Data.(*pb.FileMsg)

			name := msg.GetName()
			if proc, ok := processors[name]; !ok {
				slog.Warn("filter file: unconfigured source file '%s'", name)
			} else {
				proc(msg.GetEvent())
			}

		} //select...

	} //for...

	return
}
