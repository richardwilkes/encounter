package board

import (
	"fmt"
	"strings"

	"github.com/richardwilkes/encounter/combatant"
	"github.com/richardwilkes/toolbox/collection"
)

type Board struct {
	Round      int
	Current    *combatant.Combatant
	Combatants []*combatant.Combatant `json:"combatants,omitempty"`
}

func (b *Board) NewCombatant(nameHint string) *combatant.Combatant {
	c := combatant.New(b.suggestName(nameHint))
	b.Combatants = append(b.Combatants, c)
	return c
}

func (b *Board) DuplicateCombatant(who *combatant.Combatant) *combatant.Combatant {
	c := who.Clone()
	c.Name = b.suggestName(c.Name)
	b.Combatants = append(b.Combatants, c)
	return c
}

func (b *Board) IsCurrent(c *combatant.Combatant) bool {
	return b.Current == c
}

func (b *Board) CurrentMarker(c *combatant.Combatant) string {
	if b.Current == c {
		return " current"
	}
	return ""
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
