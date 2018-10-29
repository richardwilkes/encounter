package main

//go:generate go run board/data/dataconverter/main.go
//go:generate mkembeddedfs --no-modtime --output internal/assets/dynamic_fs_gen.go --pkg assets --name dynamicFS --strip internal/assets/dynamic internal/assets/dynamic
//go:generate mkembeddedfs --no-modtime --output internal/assets/static_fs_gen.go --pkg assets --name staticFS --strip internal/assets/static internal/assets/static

import (
	"runtime"

	"github.com/richardwilkes/encounter/internal/server"
	"github.com/richardwilkes/encounter/internal/ui"
	"github.com/richardwilkes/toolbox/atexit"
	"github.com/richardwilkes/toolbox/cmdline"
	"github.com/richardwilkes/toolbox/errs"
	"github.com/richardwilkes/toolbox/log/jot"
	"github.com/richardwilkes/toolbox/log/jotrotate"
)

func main() {
	cmdline.AppName = "Encounter"
	cmdline.AppCmdName = "encounter"
	cmdline.AppVersion = "0.1"
	cmdline.AppIdentifier = "com.trollworks.encounter"
	cmdline.CopyrightYears = "2018"
	cmdline.CopyrightHolder = "Richard A. Wilkes"
	cmdline.License = "Mozilla Public License Version 2.0"

	address := "127.0.0.1:0"
	serverOnly := false

	cl := cmdline.New(true)
	cl.NewStringOption(&address).SetSingle('a').SetName("address").SetUsage(`Network address and port to listen on. Specify just the port (":8001") to listen on all connected networks. Specify a port of 0 to pick a random port.`)
	cl.NewBoolOption(&serverOnly).SetSingle('s').SetName("server-only").SetUsage(`Do not put up a user interface, just run the server.`)
	jotrotate.ParseAndSetup(cl)

	runtime.LockOSThread() // Ensure the main thread is used when starting the UI
	s := server.New(address)
	go func() {
		if err := s.Run(); err != nil {
			jot.Error(errs.NewfWithCause(err, "%s shutdown unexpectedly", s.Protocol()))
		}
		atexit.Exit(0)
	}()
	if serverOnly {
		select {} // Wait forever
	}
	<-s.StartedChan
	ui.Start(s)
}
