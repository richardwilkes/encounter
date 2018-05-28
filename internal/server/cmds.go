package server

import (
	"bytes"
	"math"
	"net/http"
	"sort"
	"strconv"

	"github.com/richardwilkes/encounter/board"
	"github.com/richardwilkes/encounter/internal/assets"
	"github.com/richardwilkes/rpgtools/dice"
	"github.com/richardwilkes/toolbox/errs"
	"github.com/richardwilkes/toolbox/formats/json"
	"github.com/richardwilkes/toolbox/log/jot"
	"github.com/richardwilkes/toolbox/txt"
	"github.com/richardwilkes/toolbox/xio"
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
	case "addNote":
		s.addNote(w, req)
	case "editNote":
		s.editNote(w, req)
	case "adjustHP":
		s.adjustHP(w, req)
	case "editCombatant":
		s.editCombatant(w, req)
	case "rollInitiative":
		s.rollInitiative(w, req)
	case "globalOptions":
		s.globalOptions(w, req)
	default:
		http.Error(w, "unknown command: "+cmd, http.StatusBadRequest)
	}
}

func (s *Server) addNote(w http.ResponseWriter, req *http.Request) {
	// RAW: Implement
}

func (s *Server) editNote(w http.ResponseWriter, req *http.Request) {
	// RAW: Implement
}

func (s *Server) adjustHP(w http.ResponseWriter, req *http.Request) {
	j := json.MustParseStream(req.Body)
	id := int(j.Int64Relaxed("id"))
	panel := j.BoolRelaxed("panel")
	xio.CloseIgnoringErrors(req.Body)
	if c := s.board.Lookup(id); c == nil {
		http.Error(w, "no such combatant", http.StatusBadRequest)
	} else {
		tmpl, err := htmltmpl.Load(nil, assets.DynamicFS, "/", nil)
		if err != nil {
			jot.Error(errs.Wrap(err))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		var buffer bytes.Buffer
		if panel {
			if err := tmpl.ExecuteTemplate(&buffer, "/adjust_hp.html", c); err != nil {
				jot.Error(errs.Wrap(err))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		} else {
			if adjust := int(j.Int64Relaxed("adjust")); adjust != 0 {
				if adjust < 0 {
					c.Harm(-adjust)
				} else {
					c.Heal(adjust)
				}
			}
			if err := tmpl.ExecuteTemplate(&buffer, "/board.html", &s.board); err != nil {
				jot.Error(errs.Wrap(err))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}
		if _, err := buffer.WriteTo(w); err != nil {
			jot.Warn(errs.Wrap(err))
		}
	}
}

func (s *Server) editCombatant(w http.ResponseWriter, req *http.Request) {
	j := json.MustParseStream(req.Body)
	id := int(j.Int64Relaxed("id"))
	panel := j.BoolRelaxed("panel")
	xio.CloseIgnoringErrors(req.Body)
	if c := s.board.Lookup(id); c == nil {
		http.Error(w, "no such combatant", http.StatusBadRequest)
	} else {
		tmpl, err := htmltmpl.Load(nil, assets.DynamicFS, "/", nil)
		if err != nil {
			jot.Error(errs.Wrap(err))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		var buffer bytes.Buffer
		if panel {
			if err := tmpl.ExecuteTemplate(&buffer, "/edit_combatant.html", c); err != nil {
				jot.Error(errs.Wrap(err))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		} else {
			if j.Exists("name") {
				c.Name = j.Str("name")
			}
			if j.Exists("enemy") {
				c.Enemy = j.BoolRelaxed("enemy")
			}
			if j.Exists("init_base") {
				c.InitiativeBase = int(j.Int64Relaxed("init_base"))
			}
			if j.Exists("hp_full") {
				if v := j.Int64Relaxed("hp_full"); v > 0 {
					c.HPFull = int(v)
				}
			}
			if j.Exists("hp_tmp") {
				if v := j.Int64Relaxed("hp_tmp"); v >= 0 {
					c.HPTemporary = int(v)
				}
			}
			if j.Exists("hp_damage") {
				if v := j.Int64Relaxed("hp_damage"); v >= 0 {
					c.HPDamage = int(v)
				}
			}
			if err := tmpl.ExecuteTemplate(&buffer, "/board.html", &s.board); err != nil {
				jot.Error(errs.Wrap(err))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}
		if _, err := buffer.WriteTo(w); err != nil {
			jot.Warn(errs.Wrap(err))
		}
	}
}

func (s *Server) rollInitiative(w http.ResponseWriter, req *http.Request) {
	j := json.MustParseStream(req.Body)
	panel := j.BoolRelaxed("panel")
	xio.CloseIgnoringErrors(req.Body)
	tmpl, err := htmltmpl.Load(nil, assets.DynamicFS, "/", nil)
	if err != nil {
		jot.Error(errs.Wrap(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	var buffer bytes.Buffer
	if panel {
		type roll struct {
			ID         int
			Name       string
			Initiative int
			Enemy      bool
		}
		rolls := make([]roll, len(s.board.Combatants))
		for i := 0; i < len(rolls); i++ {
			rolls[i] = roll{
				ID:         s.board.Combatants[i].ID,
				Name:       s.board.Combatants[i].Name,
				Initiative: s.board.Combatants[i].InitiativeBase + s.board.RollInitiativeDice(),
				Enemy:      s.board.Combatants[i].Enemy,
			}
		}
		sort.Slice(rolls, func(i, j int) bool {
			if rolls[i].Enemy != rolls[j].Enemy {
				return rolls[j].Enemy
			}
			return txt.NaturalLess(rolls[i].Name, rolls[j].Name, true)
		})
		if err := tmpl.ExecuteTemplate(&buffer, "/roll_initiative.html", rolls); err != nil {
			jot.Error(errs.Wrap(err))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	} else {
		rnd := dice.NewCryptoRand()
		for _, c := range s.board.Combatants {
			c.RandomTieBreaker = rnd.Intn(math.MaxInt32)
		}
		if inits := j.Path("inits"); inits.IsArray() {
			count := inits.Size()
			for i := 0; i < count; i++ {
				r := inits.Index(i)
				id := int(r.Int64Relaxed("id"))
				if c := s.board.Lookup(id); c != nil {
					c.Initiative = int(r.Int64Relaxed("init"))
				}
			}
		}
		sort.Slice(s.board.Combatants, func(i, j int) bool {
			if s.board.Combatants[i].Initiative > s.board.Combatants[j].Initiative {
				return true
			}
			if s.board.Combatants[i].Initiative < s.board.Combatants[j].Initiative {
				return false
			}
			if s.board.Combatants[i].InitiativeBase > s.board.Combatants[j].InitiativeBase {
				return true
			}
			if s.board.Combatants[i].InitiativeBase < s.board.Combatants[j].InitiativeBase {
				return false
			}
			if s.board.Combatants[i].RandomTieBreaker > s.board.Combatants[j].RandomTieBreaker {
				return true
			}
			if s.board.Combatants[i].RandomTieBreaker < s.board.Combatants[j].RandomTieBreaker {
				return false
			}
			return s.board.Combatants[i].ID < s.board.Combatants[j].ID
		})
		if len(s.board.Combatants) > 0 {
			s.board.Current = s.board.Combatants[0].ID
		}
		s.board.Round = 1
		w.Header().Set("round", strconv.Itoa(s.board.Round))
		w.Header().Set("new_round", "true")
		if err := tmpl.ExecuteTemplate(&buffer, "/board.html", &s.board); err != nil {
			jot.Error(errs.Wrap(err))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	if _, err := buffer.WriteTo(w); err != nil {
		jot.Warn(errs.Wrap(err))
	}
}

func (s *Server) globalOptions(w http.ResponseWriter, req *http.Request) {
	// RAW: Implement
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
	s.board.Round = 0
	tmpl, err := htmltmpl.Load(nil, assets.DynamicFS, "/", nil)
	if err != nil {
		jot.Error(errs.Wrap(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	var buffer bytes.Buffer
	w.Header().Set("round", strconv.Itoa(s.board.Round))
	if err := tmpl.ExecuteTemplate(&buffer, "/board.html", &s.board); err != nil {
		jot.Error(errs.Wrap(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if _, err := buffer.WriteTo(w); err != nil {
		jot.Warn(errs.Wrap(err))
	}
}

func (s *Server) makeCurrentCombatant(w http.ResponseWriter, req *http.Request) {
	id := int(json.MustParseStream(req.Body).Int64Relaxed("id"))
	xio.CloseIgnoringErrors(req.Body)
	if c := s.board.Lookup(id); c == nil {
		http.Error(w, "no such combatant", http.StatusBadRequest)
	} else {
		s.board.Current = id
		s.refreshBoard(w)
	}
}

func (s *Server) duplicateCombatant(w http.ResponseWriter, req *http.Request) {
	id := int(json.MustParseStream(req.Body).Int64Relaxed("id"))
	xio.CloseIgnoringErrors(req.Body)
	if c := s.board.Lookup(id); c == nil {
		http.Error(w, "no such combatant", http.StatusBadRequest)
	} else {
		s.board.DuplicateCombatant(c)
		s.refreshBoard(w)
	}
}

func (s *Server) deleteCombatant(w http.ResponseWriter, req *http.Request) {
	id := int(json.MustParseStream(req.Body).Int64Relaxed("id"))
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
	id := int(j.Int64Relaxed("id"))
	index := int(j.Int64Relaxed("index"))
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
