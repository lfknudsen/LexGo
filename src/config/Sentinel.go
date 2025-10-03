package config

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"log"
)

// Sentinel is a set of letters indicating that the file is a LexGo token file. It follows
// the BOM, and precedes the structs.Version.
type Sentinel [5]byte

func TrueSentinel() Sentinel {
	return [5]byte{SENTINEL_1, SENTINEL_2, SENTINEL_3, SENTINEL_4, SENTINEL_5}
}

type SentinelCharacter = byte

const (
	SENTINEL_1 SentinelCharacter = 'L'
	SENTINEL_2 SentinelCharacter = 'E'
	SENTINEL_3 SentinelCharacter = 'X'
	SENTINEL_4 SentinelCharacter = 'G'
	SENTINEL_5 SentinelCharacter = 'O'
)

func SentinelCorrect(readSentinel []byte) bool {
	canonical := TrueSentinel()
	return bytes.Equal(readSentinel, (&canonical).Bytes())
}

func (s *Sentinel) String() string {
	return fmt.Sprintf("%s", *s)
}

func (s *Sentinel) Bytes() []byte {
	return s[:]
}

func (s *Sentinel) Write(w io.Writer) (totalWritten int) {
	err := binary.Write(w, BYTE_ORDER, *s)
	if err != nil {
		log.Panic(err)
	}
	return binary.Size(*s)
}
