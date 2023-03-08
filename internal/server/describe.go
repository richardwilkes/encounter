// Copyright ©2018-2023 by Richard A. Wilkes. All rights reserved.
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
	"html/template"
	"net/http"
	"strconv"

	"github.com/richardwilkes/encounter/board/data"
	"github.com/richardwilkes/toolbox/errs"
	"github.com/richardwilkes/toolbox/log/jot"
	"github.com/richardwilkes/toolbox/xio/network/xhttp/web"
)

func (s *Server) handleDescribe(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		http.Error(w, "only GET is permitted", http.StatusMethodNotAllowed)
		return
	}
	if id, err := strconv.Atoi(web.PathHeadThenShift(req)); err == nil {
		for _, e := range data.Entities {
			if e.ID == id { //nolint:gocritic
				var tmpl *template.Template
				tmpl, err = s.loadTemplates()
				if err != nil {
					jot.Error(errs.Wrap(err))
					w.WriteHeader(http.StatusInternalServerError)
					return
				}
				var buffer bytes.Buffer
				if err = tmpl.ExecuteTemplate(&buffer, "describe.html", e); err != nil {
					jot.Error(errs.Wrap(err))
					w.WriteHeader(http.StatusInternalServerError)
					return
				}
				if _, err = buffer.WriteTo(w); err != nil {
					jot.Warn(errs.Wrap(err))
				}
				return
			}
		}
	}
	http.Error(w, "invalid entity id", http.StatusNotFound)
}
