package board

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
	"sync/atomic"

	"github.com/richardwilkes/rpgtools/dice"
	"github.com/richardwilkes/toolbox/collection"
	"github.com/richardwilkes/toolbox/errs"
	"github.com/richardwilkes/toolbox/xio"
	"github.com/richardwilkes/toolbox/xio/fs/safe"
)

// Board holds the initiative board data.
type Board struct {
	Round          int
	Current        int
	Combatants     []*Combatant
	InitiativeDice *dice.Dice
	lastID         int64
}

type boardForJSON struct {
	Round          int
	Current        int
	Combatants     []*Combatant
	InitiativeDice string
	LastID         int64
}

// Load state from the specified path.
func (b *Board) Load(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return errs.Wrap(err)
	}
	defer xio.CloseIgnoringErrors(f)
	var data boardForJSON
	if err = json.NewDecoder(bufio.NewReader(f)).Decode(&data); err != nil {
		return errs.Wrap(err)
	}
	b.Round = data.Round
	b.Current = data.Current
	b.Combatants = data.Combatants
	b.InitiativeDice = dice.New(nil, data.InitiativeDice)
	b.lastID = data.LastID
	return nil
}

// Save state to the specified path.
func (b *Board) Save(path string) error {
	return safe.WriteFile(path, func(w io.Writer) error {
		encoder := json.NewEncoder(w)
		encoder.SetIndent("", "  ")
		return errs.Wrap(encoder.Encode(&boardForJSON{
			Round:          b.Round,
			Current:        b.Current,
			Combatants:     b.Combatants,
			InitiativeDice: b.InitiativeDice.String(),
			LastID:         b.lastID,
		}))
	})
}

// NextID returns the next ID to use for a combatant.
func (b *Board) NextID() int {
	return int(atomic.AddInt64(&b.lastID, 1))
}

// NewCombatant creates a new combatant and adds them to the board.
func (b *Board) NewCombatant(nameHint string) *Combatant {
	c := NewCombatant(b.NextID(), b.SuggestName(nameHint))
	b.Combatants = append(b.Combatants, c)
	return c
}

// DuplicateCombatant duplicates an existing combatant and add them to the
// board.
func (b *Board) DuplicateCombatant(who *Combatant) *Combatant {
	c := who.Clone(b.NextID())
	c.Name = b.SuggestName(c.Name)
	b.Combatants = append(b.Combatants, c)
	return c
}

// IsCurrent returns true if the specified combatant is the current one.
func (b *Board) IsCurrent(c *Combatant) bool {
	return b.Round > 0 && b.Current == c.ID
}

// CurrentTag returns the current tag to use for the combatant, if any.
func (b *Board) CurrentTag(c *Combatant) string {
	if b.IsCurrent(c) {
		return " current"
	}
	return ""
}

// Lookup a combatant by ID.
func (b *Board) Lookup(id int) *Combatant {
	for _, c := range b.Combatants {
		if c.ID == id {
			return c
		}
	}
	return nil
}

// SuggestName suggests a name for a combatant.
func (b *Board) SuggestName(nameHint string) string {
	names := collection.NewStringSet()
	for _, c := range b.Combatants {
		names.Add(strings.ToLower(c.Name))
	}
	if !names.Contains(strings.ToLower(nameHint)) {
		return nameHint
	}
	nameHint = strings.TrimRight(nameHint, "0123456789 ")
	if !strings.HasSuffix(nameHint, "#") {
		nameHint = nameHint + " "
	}
	counter := 2
	for {
		name := fmt.Sprintf("%s%d", nameHint, counter)
		if !names.Contains(strings.ToLower(name)) {
			return name
		}
		counter++
	}
}
