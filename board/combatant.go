package board

import (
	"fmt"
)

// Combantant holds information for a single entity in combat.
type Combatant struct {
	ID               int
	Name             string
	Enemy            bool
	Initiative       int
	InitiativeBase   int
	RandomTieBreaker int
	HPFull           int
	HPTemporary      int
	HPDamage         int
	Notes            []Note
}

// New creates a new combatant.
func NewCombatant(name string) *Combatant {
	return &Combatant{
		ID:     NextID(),
		Name:   name,
		Enemy:  true,
		HPFull: 6,
	}
}

// Clone creates a copy of the combatant, but with a new ID.
func (c *Combatant) Clone() *Combatant {
	clone := *c
	clone.ID = NextID()
	clone.Notes = make([]Note, len(c.Notes))
	copy(clone.Notes, c.Notes)
	return &clone
}

func (c *Combatant) Type() string {
	if c.Enemy {
		return "enemy"
	}
	return "ally"
}

func (c *Combatant) TypeDescription() string {
	if c.Enemy {
		return "Enemy"
	}
	return "Ally"
}

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

// Heal a combatant.
func (c *Combatant) Heal(amount int) {
	if amount < 0 {
		return
	}
	if amount >= c.HPDamage {
		c.HPDamage = 0
		return
	}
	c.HPDamage -= amount
}

// Harm a combatant.
func (c *Combatant) Harm(amount int) {
	if amount < 0 {
		return
	}
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

func (c *Combatant) FormattedHP() string {
	if c.HPTemporary == 0 {
		return fmt.Sprintf("%d", c.HPFull-c.HPDamage)
	}
	return fmt.Sprintf("%d+%d", c.HPFull-c.HPDamage, c.HPTemporary)
}
