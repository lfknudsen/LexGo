package src

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// Ruleset is a list of Rule structs, each describing a regular expression.
type Ruleset struct {
	Rules []Rule
}

func (rs *Ruleset) Length() int {
	return len(rs.Rules)
}

func (rs *Ruleset) Names() *[]string {
	names := make([]string, len(rs.Rules))
	for i, r := range rs.Rules {
		names[i] = r.Id
	}
	return &names
}

func (rs *Ruleset) String() string {
	sb := strings.Builder{}
	sb.WriteString("-----Rules:-----\n")
	for _, rule := range rs.Rules {
		sb.WriteString(rule.String() + "\n")
	}
	sb.WriteString("----------------\n")
	return sb.String()
}

// CompileToString combines all every Rule in the Ruleset into a single regexp.Regexp in string format.
func (rs *Ruleset) CompileToString() *string {
	sb := strings.Builder{}
	sb.WriteString("(?m)\\s*")
	for i, rule := range rs.Rules {
		sb.WriteString(rule.CompileToString())
		if i != len(rs.Rules)-1 {
			sb.WriteByte('|')
		}
	}
	sb.WriteString("(?:[^\\\\]+)")
	var str = sb.String()
	return &str
}

// Compile combines all every Rule in the Ruleset into a single regexp.Regexp.
func (rs *Ruleset) Compile() *regexp.Regexp {
	regex := regexp.MustCompile(*rs.CompileToString())
	regex.Longest()
	return regex
}

// Add appends the given Rule to the Ruleset, and returns the new length of the Ruleset.
func (rs *Ruleset) Add(rule *Rule) int {
	rs.Rules = append(rs.Rules, *rule)
	return len(rs.Rules)
}

// AddAll appends each given Rule to the Ruleset, and returns the new length of the Ruleset.
func (rs *Ruleset) AddAll(rules ...Rule) int {
	for _, rule := range rules {
		rs.Rules = append(rs.Rules, rule)
	}
	return len(rs.Rules)
}

func (rs *Ruleset) AddArray(rules []Rule) int {
	for _, rule := range rules {
		rs.Rules = append(rs.Rules, rule)
	}
	return len(rs.Rules)
}

func (rs *Ruleset) Remove(id string) int {
	for i, rule := range rs.Rules {
		if rule.Id == id {
			rs.Rules = append((rs.Rules)[:i], (rs.Rules)[i+1:]...)
			return len(rs.Rules)
		}
	}
	return len(rs.Rules)
}

func (rs *Ruleset) RemoveAll(id string) int {
	for i, rule := range rs.Rules {
		if rule.Id == id {
			rs.Rules = append((rs.Rules)[:i], (rs.Rules)[i+1:]...)
		}
	}
	return len(rs.Rules)
}

func Decompile(rx *regexp.Regexp) *Ruleset {
	bs, err := os.ReadFile("Expressions/DecompileRegex.txt")
	if err != nil {
		log.Fatal(err)
	}
	str := string(bs)
	fmt.Println("Compiling the regexp for decompiling regexes...")
	regexOfRegex := regexp.MustCompile(str)
	searchString := rx.String()
	// fmt.Println(searchString)
	matches := regexOfRegex.FindAllStringSubmatchIndex(searchString, -1)
	if matches == nil {
		return nil
	}

	rules := make([]Rule, len(matches))
	rulesCount := 0
	fmt.Printf("Found %d rules\n", len(matches))
	for i, match := range matches {
		if len(match) != 8 {
			log.Println("Match #" + strconv.Itoa(i) + " has wrong length")
			continue
		}
		ID := searchString[match[4]:match[5]]
		RegExp := searchString[match[6]:match[7]]
		fmt.Printf("Compiling an individual regexp:\n%s: %s\n", ID, RegExp)
		rules[rulesCount] = *NewRule(ID, RegExp)
		rulesCount++
	}
	newRuleset := Ruleset{rules}
	return &newRuleset
}
