package main

import (
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/pborman/getopt/v2"

	"github.com/maxb-odessa/sconf"
	"github.com/maxb-odessa/slog"

	filterFile "edpad2/internal/filter/file"
	localDisplay "edpad2/internal/local/display"
	localSound "edpad2/internal/local/sound"
	network "edpad2/internal/net"
	"edpad2/internal/router"
)

func main() {

	// get cmdline args and parse them
	help := false
	debug := 0
	configFile := os.ExpandEnv("$HOME/.local/etc/edpad2.conf")
	getopt.HelpColumn = 0
	getopt.FlagLong(&help, "help", 'h', "Show this help")
	getopt.FlagLong(&debug, "debug", 'd', "Set debug log level")
	getopt.FlagLong(&configFile, "config", 'c', "Path to config file")
	getopt.Parse()

	// help-only requested
	if help {
		getopt.Usage()
		return
	}

	// setup logger
	slog.Init("edpad2", debug, "2006-01-02 15:04:05")

	slog.Info("started")

	// read config file
	if err := sconf.Read(configFile); err != nil {
		slog.Fatal("config failed: %s", err)
	}

	// init the router
	if err := router.Init(); err != nil {
		slog.Fatal("router, failed to init: %s", err)
	}

	// register everything in router
	// no special error check, router.Register() will take care of them
	router.Register(network.Connect(router.NetFile))
	//router.Register(network.Connect(router.NetJoystick))
	//router.Register(network.Connect(router.NetKeyboard))
	//router.Register(network.Connect(router.NetSound))
	router.Register(filterFile.Connect(router.FilterFile))
	router.Register(localDisplay.Connect(router.LocalDisplay))
	router.Register(localSound.Connect(router.LocalSound))

	// set proggie termination signal handler(s)
	done := make(chan bool)
	go func() {
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
		for sig := range sigChan {
			slog.Info("got signal '%s'", sig)
			done <- true
		}
	}()

	// run the router
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		router.DoRouting()
	}()

	// wait for term signal
	<-done

	// unregister all endpoints
	router.Unregister(router.NetFile)
	//router.Unregister(router.NetJoystick)
	//router.Unregister(router.NetKeyboard)
	//router.Unregister(router.NetSound)
	router.Unregister(router.FilterFile)
	router.Unregister(router.LocalDisplay)
	router.Unregister(router.LocalSound)

	// wait for the router to finish
	wg.Wait()

	slog.Info("stopped")

	// all done
	return
}
