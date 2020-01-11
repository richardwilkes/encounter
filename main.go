// Copyright ©2018-2020 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package main

//go:generate go run board/data/dataconverter/main.go

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
	cmdline.AppIdentifier = "com.trollworks.encounter"
	cmdline.CopyrightYears = "2018-2020"
	cmdline.CopyrightHolder = "Richard A. Wilkes"
	cmdline.License = "Mozilla Public License Version 2.0"

	address := "127.0.0.1:0"

	cl := cmdline.New(true)
	cl.NewStringOption(&address).SetSingle('a').SetName("address").SetUsage(`Network address and port to listen on. Specify just the port (":8001") to listen on all connected networks. Specify a port of 0 to pick a random port.`)
	jotrotate.ParseAndSetup(cl)

	s := server.New(address)
	if err := s.Run(); err != nil {
		jot.Error(errs.NewWithCausef(err, "%s shutdown unexpectedly", s.Protocol()))
	}
	atexit.Exit(0)
}
