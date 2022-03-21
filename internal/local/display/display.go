package display

import (
	"fmt"
	"io/ioutil"
	"os"
	"runtime"

	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"

	"github.com/maxb-odessa/sconf"
	"github.com/maxb-odessa/slog"

	"edpad2/internal/router"
)

const (
	VP_SYS int = iota
	VP_SIG
	VP_PLANETS
	VP_SRV
)

// we expect this from file filter
type Text struct {
	ViewPort       int
	Text           string
	UpdateText     bool
	AppendText     bool
	Subtitle       string
	UpdateSubtitle bool
}

type viewPort struct {
	title string
	view  *gtk.TextView
	buff  *gtk.TextBuffer
	sw    *gtk.ScrolledWindow
}

type handler struct {
	endpoint    router.Endpoint
	connector   *router.Connector
	resourceDir string
	gtkBuilder  *gtk.Builder
	viewPorts   map[int]*viewPort
	gtkStack    *gtk.Stack
}

func Connect(ep router.Endpoint) (router.Endpoint, *router.Connector) {
	h := new(handler)
	h.endpoint = ep
	h.connector = new(router.Connector)
	h.connector.FromRouterCh = make(chan *router.Message) // wait for router messages on this chan
	h.connector.ToRouterCh = make(chan *router.Message)   // send messages to the router into this chan
	h.connector.DoneCh = make(chan bool)                  // termination chan

	if err := h.init(); err != nil {
		slog.Err("endpoint '%s': failed: %s", ep, err)
		return ep, nil
	}

	// run!
	go h.run()

	// all done, return router connector
	return ep, h.connector
}

func (h *handler) init() (err error) {

	runtime.LockOSThread()

	h.resourceDir = os.ExpandEnv(sconf.StrDef("local display", "resource dir", "$HOME/.local/share/edpad2"))
	slog.Debug(9, "endpoint '%s': resource dir is '%s'", h.endpoint, h.resourceDir)

	// read ui into mem
	bres, err := ioutil.ReadFile(h.resourceDir + "/edpad2.ui")
	if err != nil {
		return
	}

	gtk.Init(nil)

	h.gtkBuilder, err = gtk.BuilderNew()
	if err != nil {
		return
	}

	if err = h.gtkBuilder.AddFromString(string(bres)); err != nil {
		return
	}

	obj, err := h.gtkBuilder.GetObject("top")
	if err != nil {
		return
	}

	top, ok := obj.(*gtk.Window)
	if !ok {
		err = fmt.Errorf("GTK3 error: obj 'top' is not a window")
		return
	}

	top.Connect("destroy", func() {
		gtk.MainQuit()
		os.Exit(0) // TODO graceful exit
	})

	css, _ := gtk.CssProviderNew()
	css.LoadFromPath(h.resourceDir + "/edpad2.css")
	if err != nil {
		return err
	}

	screen, _ := gdk.ScreenGetDefault()
	gtk.AddProviderForScreen(screen, css, gtk.STYLE_PROVIDER_PRIORITY_USER)

	// init and run keypad processor
	if err = h.keypad(); err != nil {
		return
	}

	// init and setup text view ports

	h.viewPorts = make(map[int]*viewPort)

	st, err := h.gtkBuilder.GetObject("stack")
	if err != nil {
		return err
	}
	h.gtkStack = st.(*gtk.Stack)

	if obj, err = h.gtkBuilder.GetObject("system_tv"); err != nil {
		return
	} else {
		h.viewPorts[VP_SYS] = new(viewPort)
		h.viewPorts[VP_SYS].title = "System"
		h.viewPorts[VP_SYS].view = obj.(*gtk.TextView)
		if h.viewPorts[VP_SYS].buff, err = h.viewPorts[VP_SYS].view.GetBuffer(); err != nil {
			return
		}
		if sw, err := h.gtkBuilder.GetObject("system_sw"); err != nil {
			return err
		} else {
			h.viewPorts[VP_SYS].sw = sw.(*gtk.ScrolledWindow)
		}
	}

	if obj, err = h.gtkBuilder.GetObject("signals_tv"); err != nil {
		return
	} else {
		h.viewPorts[VP_SIG] = new(viewPort)
		h.viewPorts[VP_SIG].title = "Signals"
		h.viewPorts[VP_SIG].view = obj.(*gtk.TextView)
		if h.viewPorts[VP_SIG].buff, err = h.viewPorts[VP_SIG].view.GetBuffer(); err != nil {
			return
		}
		if sw, err := h.gtkBuilder.GetObject("signals_sw"); err != nil {
			return err
		} else {
			h.viewPorts[VP_SIG].sw = sw.(*gtk.ScrolledWindow)
		}
	}

	if obj, err = h.gtkBuilder.GetObject("planets_tv"); err != nil {
		return
	} else {
		h.viewPorts[VP_PLANETS] = new(viewPort)
		h.viewPorts[VP_PLANETS].title = "PLANETS"
		h.viewPorts[VP_PLANETS].view = obj.(*gtk.TextView)
		if h.viewPorts[VP_PLANETS].buff, err = h.viewPorts[VP_PLANETS].view.GetBuffer(); err != nil {
			return
		}
		if sw, err := h.gtkBuilder.GetObject("planets_sw"); err != nil {
			return err
		} else {
			h.viewPorts[VP_PLANETS].sw = sw.(*gtk.ScrolledWindow)
		}
	}

	if obj, err = h.gtkBuilder.GetObject("srv_tv"); err != nil {
		return
	} else {
		h.viewPorts[VP_SRV] = new(viewPort)
		h.viewPorts[VP_SRV].title = "SRV"
		h.viewPorts[VP_SRV].view = obj.(*gtk.TextView)
		if h.viewPorts[VP_SRV].buff, err = h.viewPorts[VP_SRV].view.GetBuffer(); err != nil {
			return
		}
		if sw, err := h.gtkBuilder.GetObject("srv_sw"); err != nil {
			return err
		} else {
			h.viewPorts[VP_SRV].sw = sw.(*gtk.ScrolledWindow)
		}
	}

	// inits done, show the window
	top.ShowAll()

	if ok := sconf.BoolDef("local display", "maximize window", false); ok {
		top.Maximize()
	}

	go gtk.Main()

	return
}

func (h *handler) run() {

	for {

		select {

		case <-h.connector.DoneCh:

			close(h.connector.ToRouterCh)
			close(h.connector.FromRouterCh)
			return

		case d, ok := <-h.connector.FromRouterCh:

			if !ok {
				return // must exit
			}

			// we support file filter only atm
			if d.Src != router.FilterFile {
				continue
			}

			// display data from file filter
			glib.IdleAdd(func() bool { return h.printText(d.Data.(*Text)) })

		} //select

	} //for
}

func (h *handler) printText(t *Text) (ret bool) {

	ret = false

	vp, ok := h.viewPorts[t.ViewPort]
	if !ok {
		slog.Warn("endpoint '%s': undefind viewport %d", t.ViewPort)
		return
	}

	// update title
	if t.UpdateSubtitle {
		h.gtkStack.ChildSetProperty(vp.sw, "title", vp.title+"\n"+t.Subtitle)
	}

	// print the text
	if t.UpdateText {
		// clear viewport if needed
		if !t.AppendText {
			vp.buff.SetText("")
		}
		vp.buff.InsertMarkup(vp.buff.GetEndIter(), t.Text)
	}

	return
}
