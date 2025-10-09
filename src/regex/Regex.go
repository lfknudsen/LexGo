package regex

import (
	"io"
	. "regexp"
)

type Regex struct {
	src      *Regexp
	SubNames map[string]int
}

// NewRegex initialises a new Regex struct, which involves filling the SubNames
// map for faster name-based matching of capture groups.
func NewRegex(src *Regexp) *Regex {
	subNames := src.SubexpNames()
	subNameMap := make(map[string]int, len(subNames))
	for i, name := range subNames {
		subNameMap[name] = i
	}
	return &Regex{src, subNameMap}
}

func (re *Regex) Src() *Regexp {
	return re.src
}

func (re *Regex) SubexpIndex(name string) int {
	return re.SubNames[name]
}

func (re *Regex) SubExpIndex(name eGroupNames) int {
	return re.SubNames[string(name)]
}

func (re *Regex) Group(name eGroupNames, match []int) (left, right int) {
	index := re.SubexpIndex(string(name))
	if index == -1 {
		return -1, -1
	}
	return match[index*2], match[index*2+1]
}

func (re *Regex) GroupMatched(name eGroupNames, match []int) bool {
	index := re.SubexpIndex(string(name))
	return index != -1 && match[index*2] != match[index*2+1]
}

func MustCompileRegex(s string) *Regex {
	src := MustCompile(s)
	return &Regex{
		src: src,
	}
}

func (re *Regex) FindSubmatchIndex(b []byte) []int {
	return re.src.FindSubmatchIndex(b)
}

func (re *Regex) FindAllSubmatchIndex(b *[]byte) [][]int {
	return re.src.FindAllSubmatchIndex(*b, -1)
}

func (re *Regex) FindReaderSubmatchIndex(r io.RuneReader) []int {
	return re.src.FindReaderSubmatchIndex(r)
}

// eGroupNames is an enum for the named capture groups in the expression which
// reads the rule file.
type eGroupNames string

const (
	ID            eGroupNames = "ID"
	REGEX         eGroupNames = "REGEX"
	COMMENT_BLOCK eGroupNames = "COMMENT_BLOCK"
	COMMENT_LINE  eGroupNames = "COMMENT_LINE"
	MISTAKE       eGroupNames = "MISTAKE"
)
