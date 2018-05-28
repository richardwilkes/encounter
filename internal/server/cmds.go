package server

import (
	"bytes"
	"net/http"
	"strconv"

	"github.com/richardwilkes/toolbox/xio"

	"github.com/richardwilkes/encounter/board"
	"github.com/richardwilkes/encounter/internal/assets"
	"github.com/richardwilkes/toolbox/errs"
	"github.com/richardwilkes/toolbox/formats/json"
	"github.com/richardwilkes/toolbox/log/jot"
	"github.com/richardwilkes/toolbox/xio/fs/embedded/htmltmpl"
	"github.com/richardwilkes/toolbox/xio/network/xhttp/web"
)

func (s *Server) handleCmds(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Error(w, "only POST is permitted", http.StatusMethodNotAllowed)
		return
	}
	cmd := web.PathHeadThenShift(req)
	switch cmd {
	case "newCombatant":
		s.newCombatant(w, req)
	case "deleteAllEnemies":
		s.deleteAllEnemies(w, req)
	case "makeCurrentCombatant":
		s.makeCurrentCombatant(w, req)
	case "duplicateCombatant":
		s.duplicateCombatant(w, req)
	case "deleteCombatant":
		s.deleteCombatant(w, req)
	case "deleteNote":
		s.deleteNote(w, req)
	case "nextTurn":
		s.nextTurn(w, req)
	default:
		http.Error(w, "unknown command: "+cmd, http.StatusBadRequest)
	}
}

func (s *Server) newCombatant(w http.ResponseWriter, req *http.Request) {
	s.board.NewCombatant("#1")
	s.refreshBoard(w)
}

func (s *Server) deleteAllEnemies(w http.ResponseWriter, req *http.Request) {
	remaining := make([]*board.Combatant, 0)
	for _, c := range s.board.Combatants {
		if !c.Enemy {
			remaining = append(remaining, c)
		}
	}
	s.board.Combatants = remaining
	s.refreshBoard(w)
}

func (s *Server) makeCurrentCombatant(w http.ResponseWriter, req *http.Request) {
	id := json.MustParseStream(req.Body).Str("id")
	xio.CloseIgnoringErrors(req.Body)
	if c := s.board.Lookup(id); c == nil {
		http.Error(w, "no such combatant", http.StatusBadRequest)
	} else {
		s.board.Current = id
		s.refreshBoard(w)
	}
}

func (s *Server) duplicateCombatant(w http.ResponseWriter, req *http.Request) {
	id := json.MustParseStream(req.Body).Str("id")
	xio.CloseIgnoringErrors(req.Body)
	if c := s.board.Lookup(id); c == nil {
		http.Error(w, "no such combatant", http.StatusBadRequest)
	} else {
		s.board.DuplicateCombatant(c)
		s.refreshBoard(w)
	}
}

func (s *Server) deleteCombatant(w http.ResponseWriter, req *http.Request) {
	id := json.MustParseStream(req.Body).Str("id")
	xio.CloseIgnoringErrors(req.Body)
	found := false
	for i, c := range s.board.Combatants {
		if c.ID == id {
			copy(s.board.Combatants[i:], s.board.Combatants[i+1:])
			s.board.Combatants[len(s.board.Combatants)-1] = nil
			s.board.Combatants = s.board.Combatants[:len(s.board.Combatants)-1]
			found = true
			break
		}
	}
	if found {
		s.refreshBoard(w)
	} else {
		http.Error(w, "no such combatant", http.StatusBadRequest)
	}
}

func (s *Server) deleteNote(w http.ResponseWriter, req *http.Request) {
	j := json.MustParseStream(req.Body)
	xio.CloseIgnoringErrors(req.Body)
	id := j.Str("id")
	index := int(j.Int64("index"))
	if c := s.board.Lookup(id); c == nil {
		http.Error(w, "no such combatant", http.StatusBadRequest)
	} else {
		if index < 0 || index >= len(c.Notes) {
			http.Error(w, "no such note", http.StatusBadRequest)
		} else {
			c.Notes = append(c.Notes[:index], c.Notes[index+1:]...)
			s.refreshBoard(w)
		}
	}
}

func (s *Server) nextTurn(w http.ResponseWriter, req *http.Request) {
	if s.board.Round > 0 {
		found := -1
		for i, c := range s.board.Combatants {
			if s.board.Current == c.ID {
				found = i
				break
			}
		}
		if found == -1 {
			if len(s.board.Combatants) > 0 {
				s.board.Current = s.board.Combatants[0].ID
			}
		} else {
			if found < len(s.board.Combatants)-1 {
				s.board.Current = s.board.Combatants[found+1].ID
			} else {
				s.board.Current = s.board.Combatants[0].ID
				s.board.Round++
				w.Header().Set("new_round", "true")
			}
		}
	}
	w.Header().Set("round", strconv.Itoa(s.board.Round))
	s.refreshBoard(w)
}

func (s *Server) refreshBoard(w http.ResponseWriter) {
	tmpl, err := htmltmpl.Load(nil, assets.DynamicFS, "/", nil)
	if err != nil {
		jot.Error(errs.Wrap(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	var buffer bytes.Buffer
	if err := tmpl.ExecuteTemplate(&buffer, "/board.html", &s.board); err != nil {
		jot.Error(errs.Wrap(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if _, err := buffer.WriteTo(w); err != nil {
		jot.Warn(errs.Wrap(err))
	}
}
