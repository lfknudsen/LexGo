package tokens

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"os"

	"LexGo/src/config"
)

type Token struct {
	TotalLength uint16
	ID          byte
	Type        TokenType
	ValueLength uint16
	Value       []byte
	Row         uint32
	Column      uint32
}

func (t *Token) Equals(other Token) bool {
	return t.TotalLength == other.TotalLength &&
		t.ID == other.ID &&
		t.Type == other.Type &&
		t.ValueLength == other.ValueLength &&
		bytes.Equal(t.Value, other.Value) &&
		t.Row == other.Row &&
		t.Column == other.Column
}

func (t *Token) String() string {
	return fmt.Sprintf("[ %d | %b | %d | %d | %v | %d | %d ]",
		t.TotalLength, t.ID, t.Type, t.ValueLength, string(t.Value), t.Row, t.Column)
}

func (t *Token) Write(w io.Writer) (bytesWritten int) {
	totalWritten := 0
	err := binary.Write(w, config.BYTE_ORDER, t.TotalLength)
	if err != nil {
		log.Panic(err)
	}
	err = binary.Write(w, config.BYTE_ORDER, t.ID)
	if err != nil {
		log.Panic(err)
	}
	err = binary.Write(w, config.BYTE_ORDER, t.Type)
	if err != nil {
		log.Panic(err)
	}
	err = binary.Write(w, config.BYTE_ORDER, t.ValueLength)
	if err != nil {
		log.Panic(err)
	}
	for i := 0; i < len(t.Value); i++ {
		err = binary.Write(w, config.BYTE_ORDER, t.Value[i])
		if err != nil {
			log.Panic(err)
		}
		totalWritten += binary.Size(t.Value[i])
	}
	err = binary.Write(w, config.BYTE_ORDER, t.Row)
	if err != nil {
		log.Panic(err)
	}
	err = binary.Write(w, config.BYTE_ORDER, t.Column)
	if err != nil {
		log.Panic(err)
	}

	totalWritten += binary.Size(t.TotalLength)
	totalWritten += binary.Size(t.ID)
	totalWritten += binary.Size(t.Type)
	totalWritten += binary.Size(t.Row)
	totalWritten += binary.Size(t.Column)
	return totalWritten
}

type TokenType uint8

func (t *Token) Print() {
	t.PrintTo(os.Stdout)
}

func (t *Token) PrintTo(out io.Writer) {
	_, _ = fmt.Fprintf(out, "ID: %d; Type: %d; Value length: %d; Value: %s\n",
		t.ID, t.Type, t.ValueLength, t.Value)
}

func DecompileToken(r io.Reader) Token {
	var t Token
	var err error
	err = binary.Read(r, config.BYTE_ORDER, &t.TotalLength)
	if err != nil {
		log.Panic(err)
	}
	err = binary.Read(r, config.BYTE_ORDER, &t.ID)
	if err != nil {
		log.Panic(err)
	}
	err = binary.Read(r, config.BYTE_ORDER, &t.Type)
	if err != nil {
		log.Panic(err)
	}
	err = binary.Read(r, config.BYTE_ORDER, &t.ValueLength)
	if err != nil {
		log.Panic(err)
	}
	t.Value = make([]byte, t.ValueLength)
	for i := 0; i < int(t.ValueLength); i++ {
		err = binary.Read(r, config.BYTE_ORDER, &t.Value[i])
		if err != nil {
			log.Panic(err)
		}
	}
	err = binary.Read(r, config.BYTE_ORDER, &t.Row)
	err = binary.Read(r, config.BYTE_ORDER, &t.Column)
	return t
}
