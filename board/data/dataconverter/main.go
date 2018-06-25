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
	"sort"
	"strconv"
	"strings"
	"unicode"

	"github.com/richardwilkes/encounter/board/data"
	"github.com/richardwilkes/toolbox/atexit"
	"github.com/richardwilkes/toolbox/collection"
	"github.com/richardwilkes/toolbox/log/jot"
	"github.com/richardwilkes/toolbox/txt"
	"github.com/richardwilkes/toolbox/xio"
	"github.com/richardwilkes/toolbox/xio/fs"
)

var pcClasses = []string{
	"alchemist",
	"arcanist",
	"barbarian",
	"bard",
	"bloodrager",
	"brawler",
	"cavalier",
	"cleric",
	"druid",
	"fighter",
	"gunslinger",
	"hunter",
	"inquisitor",
	"investigator",
	"kineticist",
	"magus",
	"medium",
	"mesmerist",
	"monk",
	"ninja",
	"occultist",
	"oracle",
	"paladin",
	"psychic",
	"ranger",
	"rogue",
	"samurai",
	"shaman",
	"skald",
	"slayer",
	"spiritualist",
	"sorcerer",
	"summoner",
	"swashbuckler",
	"vigilante",
	"warpriest",
	"witch",
	"wizard",
}

func main() {
	save(load())
}

func load() []data.Entity {
	f, err := os.Open("board/data/dataconverter/monsters.csv")
	if err != nil {
		fmt.Println(err)
		atexit.Exit(1)
	}
	defer xio.CloseIgnoringErrors(f)
	monsters := make([]data.Entity, 0)
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
		var m data.Entity
		m.Name = record[0]
		m.CR = record[1]
		m.XP = parseInt(record[2], 0, line, "XP")
		m.HasPCClass = hasPCClass(record[4])
		m.Alignment = record[6]
		m.Size = record[7]
		m.Type = record[8]
		m.SubType = record[9]
		m.Init = strings.Replace(record[10], "- ", "-", -1)
		m.Senses = record[11]
		m.Aura = record[12]
		m.AC = record[13]
		m.ACMods = record[14]
		m.HD = cleanupHD(record[16])
		m.HPMods = record[17]
		m.Saves = record[18]
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
		m.BaseAttack = parseInt(record[42], 0, line, "Base Attack")
		m.CMB = record[43]
		m.CMD = record[44]
		m.Feats = record[45]
		m.Skills = record[46]
		m.Languages = record[48]
		m.SQ = record[49]
		m.Environment = record[50]
		m.Organization = record[51]
		m.Treasure = record[52]
		m.DescriptionVisual = record[53]
		m.Source = record[55]
		m.SpecialAbilities = record[57]
		m.Description = record[58]
		m.Gear = record[66]
		m.OtherGear = record[67]
		appendData(record[68], &m.Weaknesses)
		appendData(record[80], &m.SpellsPrepared)
		if record[92] != "" {
			m.Name = record[92]
		}
		if record[90] != "" {
			// Those that have a parenthetical after their name already have the alternate name form embedded.
			if !strings.HasSuffix(m.Name, ")") {
				// Prune out some odd cases
				if !strings.HasPrefix(record[90], "Chaos Lord Of ") {
					m.Name += " (" + record[90] + ")"
				}
			}
		}
		m.Name = renameDragons(m.Name)
		m.MonsterID = parseInt(record[93], -1, line, "Monster ID")
		m.MR = parseInt(record[95], 0, line, "MR")
		monsters = append(monsters, m)
	}
	sort.Slice(monsters, func(i, j int) bool {
		return txt.NaturalLess(monsters[i].Name, monsters[j].Name, true)
	})
	return monsters
}

func appendData(in string, dest *string) {
	if in != "" {
		if *dest == "" {
			*dest = in
		} else {
			*dest += "; " + in
		}
	}
}

func renameDragons(in string) string {
	if !strings.HasSuffix(in, " Dragon") {
		return in
	}
	parts := strings.Split(in, " ")
	switch parts[0] {
	case "Mythic":
		return renameDragons(strings.Join(parts[1:], " ")) + ", Mythic"
	case "Great":
		return "Dragon, " + parts[2] + ", Great Wyrm"
	case "Mature":
		if len(parts) == 3 {
			return "Dragon, " + parts[1] + ", " + parts[0]
		}
		return "Dragon, " + parts[2] + ", Mature Adult"
	case "Young":
		if parts[1] == "Adult" {
			return "Dragon, " + parts[2] + ", Young Adult"
		}
		return "Dragon, " + parts[1] + ", Young"
	case "Very":
		return "Dragon, " + parts[2] + ", " + parts[0] + " " + parts[1]
	case "Adult", "Ancient", "Juvenile", "Old", "Wyrm", "Wyrmling":
		return "Dragon, " + parts[1] + ", " + parts[0]
	case "Elemental":
		return "Dragon, Elemental, " + parts[1]
	default:
		return "Dragon, " + strings.Join(parts[:len(parts)-1], " ")
	}
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

func cleanupHD(in string) string {
	if strings.HasPrefix(in, "(") && strings.HasSuffix(in, ")") {
		in = in[1 : len(in)-1]
	}
	if i := strings.Index(in, "HD; "); i != -1 {
		in = in[i+4:]
	}
	in = strings.Replace(in, " +", "+", -1)
	in = strings.Replace(in, "++", "+", -1)
	return in
}

func hasPCClass(classes string) bool {
	classes = strings.ToLower(classes)
	for _, c := range pcClasses {
		if strings.Contains(classes, c) {
			return true
		}
	}
	return false
}

func save(monsters []data.Entity) {
	var spelling [][]string
	if err := fs.LoadJSON("board/data/dataconverter/spelling.json", &spelling); err != nil {
		fmt.Println(err)
		atexit.Exit(1)
	}

	var buffer bytes.Buffer
	buffer.WriteString("package data\n\n")
	buffer.WriteString("// Monsters holds the monsters obtained from http://www.pathfindercommunity.net/home/databases\n")
	buffer.WriteString("var Monsters = []Entity{\n")
	typ := reflect.TypeOf(data.Entity{})
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
				fmt.Println("Unhandled type within Entity structure: ", f.Kind())
				atexit.Exit(1)
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
