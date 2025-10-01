package src

import (
	"fmt"
	"reflect"
	"regexp"
)

type Rule struct {
	Id       string
	regex    *regexp.Regexp
	Encoding reflect.Kind
}

func NewEncodedRule(id string, regex string, encoding string) *Rule {
	rule := Rule{id, regexp.MustCompile(regex), Encoding[encoding]}
	return &rule
}

func NewRule(id string, regex string) *Rule {
	rule := Rule{id, regexp.MustCompile(regex), Encoding["string"]}
	return &rule
}

func (pair *Rule) String() string {
	return pair.Id + " -> " + pair.regex.String()
}

func (pair *Rule) MatchString(str string) bool {
	return pair.regex.MatchString(str)
}

func (pair *Rule) CompileToString() string {
	if pair.Id == "?:" {
		return fmt.Sprintf("(?:%v)", pair.regex.String())
	}
	return fmt.Sprintf("(?<%v>%v)", pair.Id, pair.regex.String())
}

func (pair *Rule) Compile() *regexp.Regexp {
	return regexp.MustCompile(pair.CompileToString())
}

var Encoding map[string]reflect.Kind = map[string]reflect.Kind{
	"char":       reflect.Uint8,
	"string":     reflect.String,
	"bool":       reflect.Bool,
	"boolean":    reflect.Bool,
	"float":      reflect.Float32,
	"float32":    reflect.Float32,
	"float64":    reflect.Float64,
	"double":     reflect.Float64,
	"int":        reflect.Int,
	"int8":       reflect.Int8,
	"int16":      reflect.Int16,
	"int32":      reflect.Int32,
	"int64":      reflect.Int64,
	"uint":       reflect.Uint,
	"uint8":      reflect.Uint8,
	"uint16":     reflect.Uint16,
	"uint32":     reflect.Uint32,
	"uintptr":    reflect.Uintptr,
	"byte":       reflect.Uint8,
	"complex":    reflect.Complex64,
	"complex64":  reflect.Complex64,
	"complex128": reflect.Complex128,
	"array":      reflect.Array,
	"map":        reflect.Map,
	"struct":     reflect.Struct,
	"pointer":    reflect.Ptr,
	"func":       reflect.Func,
	"interface":  reflect.Interface,
	"slice":      reflect.Slice,
	"unsafe_ptr": reflect.UnsafePointer,
}
