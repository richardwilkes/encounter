package main

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"fmt"
	"go/format"
	"io"
	"io/ioutil"
	"os"
	"reflect"
	"strconv"
	"strings"
	"unicode"

	"github.com/richardwilkes/encounter/board"
	"github.com/richardwilkes/rpgtools/dice"
	"github.com/richardwilkes/toolbox/atexit"
	"github.com/richardwilkes/toolbox/collection"
	"github.com/richardwilkes/toolbox/log/jot"
	"github.com/richardwilkes/toolbox/xio"
	"github.com/richardwilkes/toolbox/xio/fs"
)

func main() {
	save(load())
}

func load() []board.Entity {
	f, err := os.Open("internal/assets/original/monsters.csv")
	if err != nil {
		fmt.Println(err)
		atexit.Exit(1)
	}
	defer xio.CloseIgnoringErrors(f)
	monsters := make([]board.Entity, 0)
	r := csv.NewReader(bufio.NewReader(f))
	line := 0
	for {
		record, err := r.Read()
		line++
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Printf("Line %d: %v\n", line, err)
			atexit.Exit(1)
		}
		if line == 1 {
			continue
		}
		var m board.Entity
		m.Name = record[0]
		m.CR = record[1]
		m.XP = parseInt(record[2], 0, line, "XP")
		m.Race = record[3]
		m.Class = record[4]
		//m.MonsterSource = record[5]
		m.Alignment = record[6]
		m.Size = record[7]
		m.Type = record[8]
		m.SubType = record[9]
		m.Init = strings.Replace(record[10], "- ", "-", -1)
		m.Senses = record[11]
		m.Aura = record[12]
		m.AC = record[13]
		m.ACMods = record[14]
		m.HP = parseInt(record[15], 0, line, "HP")
		m.HD = processHD(record[16])
		m.HPMods = record[17]
		m.Saves = record[18]
		// m.Fort = record[19]
		// m.Ref = record[20]
		// m.Will = record[21]
		// m.SaveMods = record[22]
		m.DefensiveAbilities = record[23]
		m.DR = record[24]
		m.Immune = record[25]
		m.Resist = record[26]
		m.SR = record[27]
		m.Weaknesses = record[28]
		m.Speed = record[29]
		m.SpeedMod = record[30]
		m.Melee = record[31]
		m.Ranged = record[32]
		m.Space = record[33]
		m.Reach = record[34]
		m.SpecialAttacks = record[35]
		m.SpellLikeAbilities = record[36]
		m.SpellsKnown = record[37]
		m.SpellsPrepared = fixSpellsPrepared(record[38])
		m.SpellDomains = record[39]
		m.AbilityScores = record[40]
		// m.AbilityScoreMods = record[41]
		m.BaseAttack = parseInt(record[42], 0, line, "Base Attack")
		m.CMB = record[43]
		m.CMD = record[44]
		m.Feats = record[45]
		m.Skills = record[46]
		m.RacialMods = record[47]
		m.Languages = record[48]
		m.SQ = record[49]
		m.Environment = record[50]
		m.Organization = record[51]
		m.Treasure = record[52]
		m.DescriptionVisual = record[53]
		m.Group = record[54]
		m.Source = record[55]
		m.IsTemplate = parseFlag(record[56], line, "Is Template")
		m.SpecialAbilities = record[57]
		m.Description = record[58]
		// m.FullHTMLText = record[59]
		m.Gender = record[60]
		m.Bloodline = record[61]
		m.ProhibitedSchools = record[62]
		m.BeforeCombat = record[63]
		m.DuringCombat = record[64]
		m.Morale = record[65]
		m.Gear = record[66]
		m.OtherGear = record[67]
		m.Vulnerability = record[68]
		m.Note = record[69]
		m.CharacterFlag = parseFlag(record[70], line, "Is Character")
		m.CompanionFlag = parseFlag(record[71], line, "Is Companion")
		m.Fly = parseFlag(record[72], line, "Has Fly Speed")
		m.Climb = parseFlag(record[73], line, "Has Climb Speed")
		m.Burrow = parseFlag(record[74], line, "Has Burrow Speed")
		m.Swim = parseFlag(record[75], line, "Has Swim Speed")
		m.Land = parseFlag(record[76], line, "Has Land Speed")
		m.TemplatesApplied = record[77]
		m.OffenseNote = record[78]
		m.BaseStatistics = record[79]
		m.ExtractsPrepared = record[80]
		m.AgeCategory = record[81]
		m.DontUseRacialHD = parseFlag(record[82], line, "Don't Use Racial HD")
		m.VariantParent = record[83]
		m.Mystery = record[84]
		m.ClassArchetypes = record[85]
		m.Patron = record[86]
		m.CompanionFamiliarLink = record[87]
		m.FocusedSchool = record[88]
		m.Traits = record[89]
		m.AlternateNameForm = record[90]
		m.StatisticsNote = record[91]
		m.LinkText = record[92]
		m.MonsterID = parseInt(record[93], -1, line, "Monster ID")
		m.UniqueMonster = parseFlag(record[94], line, "Is Unique")
		m.MR = parseInt(record[95], 0, line, "MR")
		m.Mythic = parseFlag(record[96], line, "Is Mythic")
		m.MT = parseFlag(record[97], line, "MT")
		monsters = append(monsters, m)
	}
	return monsters
}

func fixSpellsPrepared(in string) string {
	var buffer bytes.Buffer
	scanner := bufio.NewScanner(strings.NewReader(in))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		if buffer.Len() > 0 {
			buffer.WriteByte(' ')
		}
		txt := scanner.Text()
		if len(txt) > 1 && strings.HasSuffix(txt, "D") {
			if unicode.IsLower(rune(txt[len(txt)-2])) {
				txt = txt[:len(txt)-1] + " (D)"
			}
		} else if len(txt) > 2 && strings.HasSuffix(txt, "D,") {
			if unicode.IsLower(rune(txt[len(txt)-3])) {
				txt = txt[:len(txt)-2] + " (D),"
			}
		}
		buffer.WriteString(txt)
	}
	return buffer.String()
}

func processHD(in string) []*dice.Dice {
	original := in
	if strings.HasPrefix(in, "(") && strings.HasSuffix(in, ")") {
		in = in[1 : len(in)-1]
	}
	if i := strings.Index(in, "HD; "); i != -1 {
		in = in[i+4:]
	}
	in = strings.Replace(in, " +", "+", -1)
	in = strings.Replace(in, "++", "+", -1)
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
				jot.Warnf("Unable to parse HD: %s", original)
			}
			if len(result) > 0 {
				result[len(result)-1].Modifier += v
			} else {
				jot.Warnf("No dice prior to modifier: %s", original)
			}
		} else {
			result = append(result, d)
		}
	}
	return result
}

func save(monsters []board.Entity) {
	var spelling [][]string
	if err := fs.LoadJSON("internal/assets/original/spelling.json", &spelling); err != nil {
		fmt.Println(err)
		atexit.Exit(1)
	}

	var buffer bytes.Buffer
	buffer.WriteString("package data\n\n")
	buffer.WriteString("import (\n")
	buffer.WriteString("\t\"github.com/richardwilkes/encounter/board\"\n")
	buffer.WriteString("\t\"github.com/richardwilkes/rpgtools/dice\"\n")
	buffer.WriteString(")\n")
	buffer.WriteString("// Monsters holds the monsters obtained from http://www.pathfindercommunity.net/home/databases\n")
	buffer.WriteString("var Monsters = []board.Entity{\n")
	typ := reflect.TypeOf(board.Entity{})
	ids := collection.NewIntSet()
	fieldCount := typ.NumField()
	for _, e := range monsters {
		if ids.Contains(e.MonsterID) {
			fmt.Printf("Monster ID %d is duplicated\n", e.MonsterID)
			atexit.Exit(1)
		}
		ids.Add(e.MonsterID)
		buffer.WriteString("{\n")
		v := reflect.ValueOf(e)
		for i := 0; i < fieldCount; i++ {
			fs := typ.Field(i)
			f := v.Field(i)
			switch f.Kind() {
			case reflect.Bool:
				if f.Bool() {
					buffer.WriteString(fs.Name)
					buffer.WriteString(": true,\n")
				}
			case reflect.Int:
				val := f.Int()
				if val != 0 {
					fmt.Fprintf(&buffer, "%s: %d,\n", fs.Name, val)
				}
			case reflect.String:
				val := strings.TrimSpace(f.String())
				if val != "" && strings.ToLower(val) != "null" {
					// Perform some cleanup to the data
					for _, one := range spelling {
						val = strings.Replace(val, one[0], one[1], -1)
					}
					fmt.Fprintf(&buffer, "%s: %q,\n", fs.Name, val)
				}
			default:
				if fs.Name == "HD" {
					buffer.WriteString("HD: []*dice.Dice{ ")
					for j, d := range e.HD {
						if j != 0 {
							buffer.WriteString(", ")
						}
						fmt.Fprintf(&buffer, "dice.New(nil, %q)", d.String())
					}
					buffer.WriteString(" },\n")
				} else {
					fmt.Println("Unhandled type within Entity structure: ", f.Kind())
					atexit.Exit(1)
				}
			}
		}
		buffer.WriteString("},\n")
	}
	buffer.WriteString("}\n")

	data, err := format.Source(buffer.Bytes())
	if err != nil {
		fmt.Println(err)
		atexit.Exit(1)
	}
	if err = os.MkdirAll("board/data", 0755); err != nil {
		fmt.Println(err)
		atexit.Exit(1)
	}
	if err = ioutil.WriteFile("board/data/monsters_gen.go", data, 0644); err != nil {
		fmt.Println(err)
		atexit.Exit(1)
	}
}

func parseInt(in string, def, line int, what string) int {
	if in == "" {
		return def
	}
	v, err := strconv.Atoi(strings.Replace(in, ",", "", -1))
	if err != nil {
		jot.Warnf("Line %d (%s): %v", line, what, err)
		return def
	}
	return v
}

func parseFlag(in string, line int, what string) bool {
	switch in {
	case "1":
		return true
	case "0":
		return false
	default:
		jot.Warnf("Line %d (%s): invalid flag (%s)", line, what, in)
		return false
	}
}
