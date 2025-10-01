package src

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"reflect"
	"strconv"
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

func NewToken(id byte, typ TokenType, value any, filename string, row int, col int) Token {
	buffer := make([]byte, binary.Size(value))
	n, err := binary.Encode(buffer, BYTE_ORDER, value)
	if err != nil {
		panic(err)
	}
	if n != binary.Size(value) {
		log.Panic("Exception while creating a token with value ", value,
			":\nExpected ", n, " to be equal to the size in bytes of the value ",
			binary.Size(value), ".")
	}
	t := Token{
		ID:          id,
		Type:        typ,
		ValueLength: uint16(n),
		Value:       buffer,
		Row:         uint32(row),
		Column:      uint32(col),
	}

	var fields []reflect.StructField = reflect.VisibleFields(reflect.TypeOf(t))
	for _, field := range fields {
		fmt.Println(field.Name + ": " + field.Type.String() + ". Size:" +
			strconv.FormatUint(uint64(field.Type.Size()), 10))
		sizeOfField := field.Type.Size()
		t.TotalLength += uint16(sizeOfField)
	}
	fmt.Println(binary.Size(&t))
	return t
}

func (t *Token) String() string {
	return fmt.Sprintf("[ %d | %b | %d | %d | %v | %d | %d ]",
		t.TotalLength, t.ID, t.Type, t.ValueLength, string(t.Value), t.Row, t.Column)
}

func (t *Token) Marshall() []byte {
	bs := make([]byte, t.TotalLength)
	var buffer = bytes.NewBuffer(bs)
	_ = binary.Write(buffer, BYTE_ORDER, t.TotalLength)
	_ = binary.Write(buffer, BYTE_ORDER, t.ID)
	_ = binary.Write(buffer, BYTE_ORDER, t.Type)
	_ = binary.Write(buffer, BYTE_ORDER, t.ValueLength)
	for i := 0; i < len(t.Value); i++ {
		_ = binary.Write(buffer, BYTE_ORDER, t.Value[i])
	}
	_ = binary.Write(buffer, BYTE_ORDER, t.Row)
	_ = binary.Write(buffer, BYTE_ORDER, t.Column)
	return buffer.Bytes()
}

func (t *Token) Write(w io.Writer) (bytesWritten int) {
	totalWritten := 0
	err := binary.Write(w, BYTE_ORDER, t.TotalLength)
	if err != nil {
		log.Panic(err)
	}
	err = binary.Write(w, BYTE_ORDER, t.ID)
	if err != nil {
		log.Panic(err)
	}
	err = binary.Write(w, BYTE_ORDER, t.Type)
	if err != nil {
		log.Panic(err)
	}
	err = binary.Write(w, BYTE_ORDER, t.ValueLength)
	if err != nil {
		log.Panic(err)
	}
	for i := 0; i < len(t.Value); i++ {
		err = binary.Write(w, BYTE_ORDER, t.Value[i])
		if err != nil {
			log.Panic(err)
		}
		totalWritten += binary.Size(t.Value[i])
	}
	err = binary.Write(w, BYTE_ORDER, t.Row)
	if err != nil {
		log.Panic(err)
	}
	err = binary.Write(w, BYTE_ORDER, t.Column)
	if err != nil {
		log.Panic(err)
	}

	totalWritten += binary.Size(t.TotalLength)
	totalWritten += binary.Size(t.ID)
	totalWritten += binary.Size(t.Type)
	totalWritten += binary.Size(t.Row)
	totalWritten += binary.Size(t.Column)
	log.Printf("Wrote token to disk; %d bytes", totalWritten)
	return totalWritten
}

func UnmarshallToken(data []byte) (Token, error) {
	var t Token
	var err error

	n, err := binary.Decode(data, BYTE_ORDER, &t.TotalLength)
	data = data[n:]
	fmt.Printf("Length: %d\n", t.TotalLength)
	if err != nil {
		log.Fatal(err)
	}

	n, err = binary.Decode(data, BYTE_ORDER, &t.ID)
	data = data[n:]
	fmt.Printf("ID: %d\n", t.ID)
	if err != nil {
		log.Fatal(err)
	}

	n, err = binary.Decode(data, BYTE_ORDER, &t.Type)
	data = data[n:]
	fmt.Printf("Type: %d\n", t.Type)
	if err != nil {
		log.Fatal(err)
	}

	n, err = binary.Decode(data, BYTE_ORDER, &t.ValueLength)
	data = data[n:]
	fmt.Printf("Value length: %d\n", t.ValueLength)
	if err != nil {
		log.Fatal(err)
	}

	t.Value = data[:t.ValueLength]
	data = data[t.ValueLength:]
	fmt.Printf("Value: %v\n", string(t.Value))

	n, err = binary.Decode(data, BYTE_ORDER, &t.Row)
	data = data[n:]
	fmt.Printf("Row: %d\n", t.Row)
	if err != nil {
		log.Fatal(err)
	}

	n, err = binary.Decode(data, BYTE_ORDER, &t.Column)
	fmt.Printf("Column: %d\n", t.Column)
	if err != nil {
		log.Fatal(err)
	}

	return t, nil
}

func ReadToken(data []byte) (Token, error) {
	var t Token
	var err error
	b := new(BuffStruct)
	b.b = data
	b = b.DecodeU16(&t.TotalLength)
	b = b.DecodeU8(&t.ID)
	b = b.DecodeTokenType(&t.Type)
	b = b.DecodeU16(&t.ValueLength)
	t.Value = make([]byte, t.ValueLength)
	b = b.DecodeByteArray(t.Value)
	b = b.DecodeU32(&t.Row)
	b = b.DecodeU32(&t.Column)
	return t, err
}

type TokenType uint8
