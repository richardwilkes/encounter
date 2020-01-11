// Copyright ©2018-2020 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package board

import (
	"fmt"
	"strings"

	"github.com/richardwilkes/encounter/board/data"
)

const hpChangesToKeep = 20

// Combatant holds information for a single entity in combat.
type Combatant struct {
	ID               int      `json:"id"`
	Name             string   `json:"name,omitempty"`
	Initiative       int      `json:"initiative,omitempty"`
	InitiativeBase   int      `json:"initiative_base,omitempty"`
	RandomTieBreaker int      `json:"random_tie_breaker,omitempty"`
	HPFull           int      `json:"hp_full,omitempty"`
	HPTemporary      int      `json:"hp_temp,omitempty"`
	HPDamage         int      `json:"hp_damage,omitempty"`
	HPChanges        []string `json:"hp_changes,omitempty"`
	EntityID         int      `json:"entity_id,omitempty"`
	Notes            []Note   `json:"notes,omitempty"`
	Enemy            bool     `json:"enemy,omitempty"`
}

// PossessiveName returns the possessive form of the combatant's name.
func (c *Combatant) PossessiveName() string {
	if strings.HasSuffix(strings.ToLower(c.Name), "s") {
		return c.Name + "'"
	}
	return c.Name + "'s"
}

// BasedOn returns the name of the entity the combatant was based on.
func (c *Combatant) BasedOn() string {
	for _, entity := range data.Entities {
		if entity.ID == c.EntityID {
			return entity.Name
		}
	}
	return ""
}

// Clone creates a copy of the combatant, but with a new ID.
func (c *Combatant) Clone(newID int) *Combatant {
	clone := *c
	clone.ID = newID
	clone.Notes = make([]Note, len(c.Notes))
	copy(clone.Notes, c.Notes)
	return &clone
}

// Type returns the type of combatant, enemy or ally.
func (c *Combatant) Type() string {
	if c.Enemy {
		return "enemy"
	}
	return "ally"
}

// TypeDescription returns a human-readable type for the combatant.
func (c *Combatant) TypeDescription() string {
	if c.Enemy {
		return "Enemy"
	}
	return "Ally"
}

// Out returns true if the combatant is out of the combat.
func (c *Combatant) Out() bool {
	return c.CurrentHP() < 0
}

// Status returns a description of the combatant's current health.
func (c *Combatant) Status() string {
	if c.HPDamage == 0 {
		return "Healthy"
	}
	current := c.CurrentHP()
	if current > c.HPFull/2 {
		return "Hurt"
	}
	if current > 0 {
		return "Bloody"
	}
	if current == 0 {
		return "Staggered"
	}
	return "Dying"
}

// StatusTag returns the current status tag to use.
func (c *Combatant) StatusTag() string {
	if c.CurrentHP() <= c.HPFull/2 {
		return " danger"
	}
	return ""
}

// CurrentHP returns the current hit point total.
func (c *Combatant) CurrentHP() int {
	return c.HPFull + c.HPTemporary - c.HPDamage
}

// Heal a combatant. 'round' indicates what round the healing occurred.
func (c *Combatant) Heal(round, amount int) {
	if amount < 0 {
		return
	}
	total := amount
	current := c.FormattedHP()
	defer func() {
		now := c.FormattedHP()
		if current != now {
			c.RecordHPChange(round, total)
		}
	}()
	if amount >= c.HPDamage {
		total = c.HPDamage
		c.HPDamage = 0
		return
	}
	c.HPDamage -= amount
}

// Harm a combatant. 'round' indicates what round the harm occurred.
func (c *Combatant) Harm(round, amount int) {
	if amount < 0 {
		return
	}
	total := amount
	current := c.FormattedHP()
	defer func() {
		now := c.FormattedHP()
		if current != now {
			c.RecordHPChange(round, -total)
		}
	}()
	if c.HPTemporary >= amount {
		c.HPTemporary -= amount
		return
	}
	if c.HPTemporary > 0 {
		amount -= c.HPTemporary
		c.HPTemporary = 0
	}
	c.HPDamage += amount
}

// RecordHPChange records a hit point change. 'round' indicates what round the
// change occurred.
func (c *Combatant) RecordHPChange(round, delta int) {
	count := len(c.HPChanges) + 1
	if count > hpChangesToKeep {
		count = hpChangesToKeep
	}
	if len(c.HPChanges) != count {
		revised := make([]string, count)
		copy(revised, c.HPChanges)
		c.HPChanges = revised
	}
	if count > 1 {
		copy(c.HPChanges[1:], c.HPChanges)
	}
	var change strings.Builder
	if round > 0 {
		fmt.Fprintf(&change, "Round %d: ", round)
	}
	if delta != 0 {
		fmt.Fprintf(&change, "%+d (now ", delta)
	}
	change.WriteString(c.FormattedHP())
	if delta != 0 {
		change.WriteString(")")
	}
	c.HPChanges[0] = change.String()
}

// FormattedHP returns the HP formatted for display.
func (c *Combatant) FormattedHP() string {
	if c.HPTemporary == 0 {
		return fmt.Sprintf("%d", c.HPFull-c.HPDamage)
	}
	return fmt.Sprintf("%d+%d", c.HPFull-c.HPDamage, c.HPTemporary)
}
