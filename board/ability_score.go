package board

// AbilityScore holds the ability score's value plus any damage.
type AbilityScore struct {
	Value  int `json:"value"`
	Damage int `json:"damage"`
}

// Current returns the current ability score value.
func (a *AbilityScore) Current() int {
	value := a.Value - a.Damage
	if value < 0 {
		return 0
	}
	return value
}

// Bonus returns the current bonus value.
func (a *AbilityScore) Bonus() int {
	current := a.Current() - 10
	if current > 0 {
		return current / 2
	}
	return (current - 1) / 2
}
