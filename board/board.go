package board

import (
	"fmt"
	"strings"
	"sync/atomic"

	"github.com/richardwilkes/encounter/board/data"
	"github.com/richardwilkes/rpgtools/dice"
	"github.com/richardwilkes/toolbox/collection"
	"github.com/richardwilkes/toolbox/xio/fs"
)

// Board holds the initiative board data.
type Board struct {
	Round            int
	Current          int
	Combatants       []*Combatant
	InitiativeDice   *dice.Dice
	HPMethod         HPMethod
	LibrarySelection int
	LibraryEntity    *data.Entity `json:"-"`
	LastID           int64        `json:"-"`
}

// Load state from the specified path.
func (b *Board) Load(path string) error {
	if err := fs.LoadJSON(path, b); err != nil {
		return err
	}
	if b.InitiativeDice == nil {
		b.InitiativeDice = dice.New(nil, "1d20")
	}
	var monster *data.Entity
	for _, m := range data.Monsters {
		if m.MonsterID == b.LibrarySelection {
			monster = &m
			break
		}
	}
	if monster == nil {
		monster = &data.Monsters[0]
	}
	b.SetLibrarySelection(monster)
	b.LastID = 0
	for _, c := range b.Combatants {
		if b.LastID < int64(c.ID) {
			b.LastID = int64(c.ID)
		}
	}
	return nil
}

// SetLibrarySelection sets the current library selection.
func (b *Board) SetLibrarySelection(e *data.Entity) {
	b.LibrarySelection = e.MonsterID
	b.LibraryEntity = e
}

// Save state to the specified path.
func (b *Board) Save(path string) error {
	return fs.SaveJSON(path, b, true)
}

// NextID returns the next ID to use for a combatant.
func (b *Board) NextID() int {
	return int(atomic.AddInt64(&b.LastID, 1))
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
