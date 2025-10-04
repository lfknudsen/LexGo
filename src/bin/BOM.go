package bin

import (
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	"LexGo/src/config"
)

// BOM represents the Byte Order Mark (Unicode code point FEFF).
// It is placed as the very first character(s) in a file, and indicates the
// correct endianness (i.e. byte order). If the receiver reads them to be FFFE,
// a meaningless Unicode code point, then it knows to switch. Because FEFF
// represents a zero-width blank space, it has the added benefit of being invisible
// in plaintext.
//
// It is present in the output of this programme because the big endian byte
// order is much easier to read for a human, and is therefore what is being used.
// However, the little endian format is much more commonly used on machines, and
// the default for any programme reading file contents.
// The inclusion of the BOM makes it possible for the receiver - in whatever
// programme they use - to adjust and understand the contents correctly.
type BOM uint16

// NewBOM returns a new Byte Order Mark with the correct code point.
func NewBOM() BOM {
	return correct
}

// correct is the true version.
const correct = 0xFEFF

// reversed is the incorrect version. If this is read at the start of a file,
// the endianness should be switched.
const reversed = 0xFFFE

func (b BOM) IsCorrect() bool {
	return b == correct
}

func (b BOM) IsReversed() bool {
	return b == reversed
}

// CheckBOM returns 1 if the BOM indicates correct endianness, 0 if reversed,
// and -1 if the characters were not a BOM.
func (b BOM) CheckBOM() int {
	if b.IsCorrect() {
		return 1
	} else if b.IsReversed() {
		return 0
	} else {
		return -1
	}
}

func (b BOM) String() string {
	str := strconv.FormatUint(uint64(b), 16)
	return fmt.Sprintf("BOM: %x", strings.ToUpper(str))
}

// Print is a convenience function for calling PrintTo on the standard output.
func (b BOM) Print() {
	b.PrintTo(os.Stdout)
}

// PrintTo prints text explaining what this BOM indicates regarding the reading order
// of the text it was read from.
func (b BOM) PrintTo(out io.Writer) {
	if b.IsCorrect() {
		_, _ = fmt.Fprintf(out, "BOM indicates the correct reading order.")
	} else if b.IsReversed() {
		_, _ = fmt.Fprintf(out, "BOM indicates the reversed reading order.")
	} else {
		_, _ = fmt.Fprintf(out, "BOM not found.")
	}
}

// Write encodes and writes the BOM to a binary format.
func (b BOM) Write(w io.Writer) (totalWritten int) {
	err := binary.Write(w, config.BYTE_ORDER, b)
	if err != nil {
		log.Panic(err)
	}
	return binary.Size(b)
}

// DecodeBOM reads two bytes, and toggles the endianness going forward accordingly.
// If a BOM could not be found, it panics.
func DecodeBOM(r io.Reader) (fileBOM BOM) {
	var bom BOM
	err := binary.Read(r, config.BYTE_ORDER, &bom)
	if err != nil {
		log.Panic(err)
	}
	if bom.IsReversed() {
		config.ToggleByteOrder()
	} else if bom.IsCorrect() {
	} else {
		log.Panicf("Did not find bom at start of file.\n")
	}
	return bom
}
