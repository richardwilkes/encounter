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
	"strconv"
	"strings"

	"github.com/richardwilkes/encounter/board/data"
	"github.com/richardwilkes/rpgtools/dice"
)

// HPMethod is an enum of possible ways to determine HP for an entity.
type HPMethod byte

// Possible HPMethods.
const (
	RolledHPMethod HPMethod = iota
	MinimumHPMethod
	AverageHPMethod
	MaximumHPMethod
	ThreeQuartersHPMethod
)

var (
	// AllHPMethods holds all possible HPMethods.
	AllHPMethods   []HPMethod
	hpMethodsTable = []struct {
		key    string
		desc   string
		method func(*data.Entity) int
	}{
		{
			key:    "rolled",
			desc:   "Rolled",
			method: rolledHP,
		},
		{
			key:    "minimum",
			desc:   "Minimum",
			method: minimumHP,
		},
		{
			key:    "average",
			desc:   "Average",
			method: averageHP,
		},
		{
			key:    "maximum",
			desc:   "Maximum",
			method: maximumHP,
		},
		{
			key:    "three-quarters",
			desc:   "Three-quarters",
			method: threeQuartersHP,
		},
	}
)

func init() {
	AllHPMethods = make([]HPMethod, len(hpMethodsTable))
	for i := 0; i < len(hpMethodsTable); i++ {
		AllHPMethods[i] = HPMethod(i)
	}
}

// AllMethods returns all possible methods.
func (hpm HPMethod) AllMethods() []HPMethod {
	return AllHPMethods
}

// Lookup returns the HPMethod by keyed name.
func (hpm HPMethod) Lookup(key string) HPMethod {
	for i, one := range hpMethodsTable {
		if one.key == key {
			return HPMethod(i)
		}
	}
	return AverageHPMethod
}

func (hpm HPMethod) index() int {
	i := int(hpm)
	if i < 0 || i >= len(hpMethodsTable) {
		return 0
	}
	return i
}

// HP determines the HP for an entity based on the method chosen.
func (hpm HPMethod) HP(e *data.Entity) int {
	return hpMethodsTable[hpm.index()].method(e)
}

func (hpm HPMethod) String() string {
	return hpMethodsTable[hpm.index()].desc
}

// Key returns the key used for encoding.
func (hpm HPMethod) Key() string {
	return hpMethodsTable[hpm.index()].key
}

// MarshalText implements the encoding.TextMarshaler interface.
func (hpm HPMethod) MarshalText() (text []byte, err error) {
	return []byte(hpm.Key()), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (hpm *HPMethod) UnmarshalText(text []byte) error {
	key := string(text)
	for i, one := range hpMethodsTable {
		if one.key == key {
			*hpm = HPMethod(i)
			return nil
		}
	}
	*hpm = AverageHPMethod
	return nil
}

func rolledHP(e *data.Entity) int {
	total := 0
	hd := extractDice(e.HD)
	for _, d := range hd {
		total += d.Roll()
	}
	return total
}

func minimumHP(e *data.Entity) int {
	total := 0
	hd := extractDice(e.HD)
	for i, d := range hd {
		if e.HasPCClass && i == len(hd)-1 {
			d.Count--
			total += d.Sides
		}
		if d.Count > 0 {
			total += d.Minimum()
		} else {
			total += d.Modifier
		}
	}
	return total
}

func averageHP(e *data.Entity) int {
	total := 0
	hd := extractDice(e.HD)
	for i, d := range hd {
		if e.HasPCClass && i == len(hd)-1 {
			d.Count--
			total += d.Sides
		}
		if d.Count > 0 {
			total += d.Average()
		} else {
			total += d.Modifier
		}
	}
	return total
}

func maximumHP(e *data.Entity) int {
	total := 0
	hd := extractDice(e.HD)
	for _, d := range hd {
		total += d.Maximum()
	}
	return total
}

func threeQuartersHP(e *data.Entity) int {
	total := 0
	hd := extractDice(e.HD)
	for i, d := range hd {
		if e.HasPCClass && i == len(hd)-1 {
			d.Count--
			total += d.Sides
		}
		d.Sides = 3 * d.Sides / 4
		if d.Count > 0 {
			total += d.Maximum()
		} else {
			total += d.Modifier
		}
	}
	return total
}

func extractDice(in string) []*dice.Dice {
	in = strings.Replace(in, " plus ", "+", -1)
	d := dice.New(nil, in)
	if d.String() == in || "1"+d.String() == in {
		return []*dice.Dice{d}
	}
	result := make([]*dice.Dice, 0)
	for _, p := range strings.Split(in, "+") {
		d = dice.New(nil, p)
		if d.String() != p && "1"+d.String() != p {
			var buffer strings.Builder
			for _, c := range p {
				if c >= '0' && c <= '9' {
					buffer.WriteRune(c)
				} else {
					break
				}
			}
			v, err := strconv.Atoi(buffer.String())
			if err != nil {
				v = 0
			}
			if len(result) > 0 {
				result[len(result)-1].Modifier += v
			}
		} else {
			result = append(result, d)
		}
	}
	return result
}
