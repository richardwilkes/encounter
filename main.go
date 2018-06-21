package main

//go:generate go run internal/assets/original/main.go
//go:generate mkembeddedfs --no-modtime --output internal/assets/dynamic_fs_gen.go --pkg assets --name dynamicFS --strip internal/assets/dynamic internal/assets/dynamic
//go:generate mkembeddedfs --no-modtime --output internal/assets/static_fs_gen.go --pkg assets --name staticFS --strip internal/assets/static internal/assets/static

import (
	"github.com/richardwilkes/encounter/internal/server"
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
	cmdline.CopyrightYears = "2018"
	cmdline.CopyrightHolder = "Richard A. Wilkes"
	cmdline.License = "Mozilla Public License Version 2.0"

	address := "127.0.0.1:8001"

	cl := cmdline.New(true)
	cl.NewStringOption(&address).SetSingle('a').SetName("address").SetUsage(`Network address and port to listen on. Specify just the port (":8001") to listen on all connected networks`)
	jotrotate.ParseAndSetup(cl)

	s := server.New(address)
	if err := s.Run(); err != nil {
		jot.Error(errs.NewfWithCause(err, "%s shutdown unexpectedly", s.Protocol()))
	}
	atexit.Exit(0)
}
