package board

import "fmt"

// HitPoints holds the hits points of a combatant.
type HitPoints struct {
	Full      int `json:"full"`
	Temporary int `json:"tmp"`
	Damage    int `json:"damage"`
}

// Current returns the current hit point total.
func (hp *HitPoints) Current() int {
	return hp.Full + hp.Temporary - hp.Damage
}

// Heal a combatant.
func (hp *HitPoints) Heal(amount int) {
	if amount < 0 {
		return
	}
	if amount >= hp.Damage {
		hp.Damage = 0
		return
	}
	hp.Damage -= amount
}

// Harm a combatant.
func (hp *HitPoints) Harm(amount int) {
	if amount < 0 {
		return
	}
	if hp.Temporary >= amount {
		hp.Temporary -= amount
		return
	}
	if hp.Temporary > 0 {
		amount -= hp.Temporary
		hp.Temporary = 0
	}
	hp.Damage += amount
}

func (hp *HitPoints) String() string {
	if hp.Temporary == 0 {
		return fmt.Sprintf("%d", hp.Full-hp.Damage)
	}
	return fmt.Sprintf("%d+%d", hp.Full-hp.Damage, hp.Temporary)
}
