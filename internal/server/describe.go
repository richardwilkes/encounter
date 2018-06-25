package server

import (
	"bytes"
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
		for _, m := range data.Monsters {
			if m.MonsterID == id {
				tmpl, err := s.loadTemplates()
				if err != nil {
					jot.Error(errs.Wrap(err))
					w.WriteHeader(http.StatusInternalServerError)
					return
				}
				var buffer bytes.Buffer
				if err := tmpl.ExecuteTemplate(&buffer, "/describe.html", &m); err != nil {
					jot.Error(errs.Wrap(err))
					w.WriteHeader(http.StatusInternalServerError)
					return
				}
				if _, err := buffer.WriteTo(w); err != nil {
					jot.Warn(errs.Wrap(err))
				}
				return
			}
		}
	}
	http.Error(w, "invalid monster id", http.StatusNotFound)
}
