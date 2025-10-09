package rules

import (
	"fmt"
	"regexp"
)

// Rule represents a line from the ruleset file.
type Rule struct {
	Id    string
	regex *regexp.Regexp
}

// NewRule creates a new Rule struct with the given values, any tokens.Token
// derived from it being encoded as plaintext.
func NewRule(id string, regex string) *Rule {
	rule := Rule{id, regexp.MustCompile(regex)}
	return &rule
}

// String returns the Rule in human-readable string format. This would not be
// compilable as a regular expression.
func (r *Rule) String() string {
	return r.Id + " -> " + r.regex.String()
}

// MatchString returns whether the given string has any match for the Rule's regular expression.
func (r *Rule) MatchString(str string) bool {
	return r.regex.MatchString(str)
}

// CompileToString writes the Rule's ID and regular expression as a named capture
// group regular expression in string format.
//
// If the Rule's ID is `?:`, it will be a non-capturing group.
func (r *Rule) CompileToString() string {
	if r.Id == "?:" {
		return fmt.Sprintf("(?:%v)", r.regex.String())
	}
	return fmt.Sprintf("(?<%v>%v)", r.Id, r.regex.String())
}

// Compile compiles the Rule to a regular expression after calling CompileToString.
func (r *Rule) Compile() *regexp.Regexp {
	return regexp.MustCompile(r.CompileToString())
}
