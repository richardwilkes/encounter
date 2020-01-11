// Copyright ©2018-2020 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package server

import (
	"bytes"
	"net/http"

	"github.com/richardwilkes/encounter/board"
	"github.com/richardwilkes/encounter/board/data"
	"github.com/richardwilkes/toolbox/cmdline"
	"github.com/richardwilkes/toolbox/errs"
	"github.com/richardwilkes/toolbox/log/jot"
)

type document struct {
	Title     string
	Copyright string
	Board     *board.Board
	Entities  []*data.Entity
}

func (s *Server) handleIndex(w http.ResponseWriter, req *http.Request) {
	tmpl, err := s.loadTemplates()
	if err != nil {
		jot.Error(errs.Wrap(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	var buffer bytes.Buffer
	if err = tmpl.ExecuteTemplate(&buffer, "/index.html", &document{
		Title:     cmdline.AppName,
		Copyright: cmdline.Copyright(),
		Board:     &s.board,
		Entities:  data.Entities,
	}); err != nil {
		jot.Error(errs.Wrap(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if _, err = buffer.WriteTo(w); err != nil {
		jot.Warn(errs.Wrap(err))
	}
}
