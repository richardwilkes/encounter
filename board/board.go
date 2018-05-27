package board

import (
	"fmt"
	"strings"

	"github.com/richardwilkes/toolbox/collection"
)

type Board struct {
	Round      int          `json:"round,omitempty"`
	Current    string       `json:"current,omitempty"`
	Combatants []*Combatant `json:"combatants,omitempty"`
}

func (b *Board) NewCombatant(nameHint string) *Combatant {
	c := NewCombatant(b.suggestName(nameHint))
	b.Combatants = append(b.Combatants, c)
	return c
}

func (b *Board) DuplicateCombatant(who *Combatant) *Combatant {
	c := who.Clone()
	c.Name = b.suggestName(c.Name)
	b.Combatants = append(b.Combatants, c)
	return c
}

func (b *Board) IsCurrent(c *Combatant) bool {
	return b.Current == c.ID
}

func (b *Board) CurrentTag(c *Combatant) string {
	if b.IsCurrent(c) {
		return " current"
	}
	return ""
}

func (b *Board) Lookup(id string) *Combatant {
	for _, c := range b.Combatants {
		if c.ID == id {
			return c
		}
	}
	return nil
}

func (b *Board) suggestName(nameHint string) string {
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
