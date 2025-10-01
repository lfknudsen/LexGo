package src

import (
	"fmt"
	"regexp"
)

type Rule struct {
	id    string
	regex *regexp.Regexp
}

func NewRule(id string, regex string) *Rule {
	rule := Rule{id, regexp.MustCompile(regex)}
	return &rule
}

func (pair *Rule) String() string {
	return pair.id + " -> " + pair.regex.String()
}

func (pair *Rule) MatchString(str string) bool {
	return pair.regex.MatchString(str)
}

func (pair *Rule) CompileToString() string {
	if pair.id == "?:" {
		return fmt.Sprintf("(?:%v)", pair.regex.String())
	}
	return fmt.Sprintf("(?<%v>%v)", pair.id, pair.regex.String())
}

func (pair *Rule) Compile() *regexp.Regexp {
	return regexp.MustCompile(pair.CompileToString())
}
