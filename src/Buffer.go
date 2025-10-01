package src

import (
	"encoding/binary"
	"fmt"
	"log"
	"reflect"
	"strconv"
)

type Buffer []byte

type BuffStruct struct {
	b []byte
}

func (b *BuffStruct) DecodeU8(dst *uint8) *BuffStruct {
	n, err := binary.Decode(b.b, binary.BigEndian, dst)
	fmt.Println("Read " + strconv.Itoa(n) + " bytes: " +
		strconv.FormatUint(uint64(*dst), 10))
	if err != nil {
		log.Fatal(err)
	}
	b.b = b.b[n:]
	return b
}

func (b *BuffStruct) DecodeU16(dst *uint16) *BuffStruct {
	n, err := binary.Decode(b.b, binary.BigEndian, dst)
	fmt.Println("Read " + strconv.Itoa(n) + " bytes: " +
		strconv.FormatUint(uint64(*dst), 10))
	if err != nil {
		log.Fatal(err)
	}
	b.b = b.b[n:]
	return b
}

func (b *BuffStruct) DecodeU32(dst *uint32) *BuffStruct {
	n, err := binary.Decode(b.b, binary.BigEndian, dst)
	fmt.Println("Read " + strconv.Itoa(n) + " bytes: " +
		strconv.FormatUint(uint64(*dst), 10))
	if err != nil {
		log.Fatal(err)
	}
	b.b = b.b[n:]
	return b
}

func (b *BuffStruct) DecodeByte(dst *byte) *BuffStruct {
	n, err := binary.Decode(b.b, binary.BigEndian, dst)
	fmt.Println("Read " + strconv.Itoa(n) + " bytes: " +
		strconv.FormatUint(uint64(*dst), 10))
	if err != nil {
		log.Fatal(err)
	}
	b.b = b.b[n:]
	return b
}

func (b *BuffStruct) DecodeRune(dst *rune) *BuffStruct {
	n, err := binary.Decode(b.b, binary.BigEndian, dst)
	fmt.Println("Read " + strconv.Itoa(n) + " bytes: " +
		strconv.FormatUint(uint64(*dst), 10))
	if err != nil {
		log.Fatal(err)
	}
	b.b = b.b[n:]
	return b
}

func (b *BuffStruct) DecodeTo(dst any) *BuffStruct {
	if reflect.TypeOf(dst).Kind() == reflect.Array {
		return b.DecodeByteArray(dst.([]byte))
	}
	n, err := binary.Decode(b.b, binary.BigEndian, &dst)
	if err != nil {
		log.Fatal(err)
	}
	b.b = b.b[n:]
	return b
}

func (b *BuffStruct) DecodeByteArray(dst []byte) *BuffStruct {
	length := len(dst)
	for i := 0; i < length; i++ {
		n, err := binary.Decode(b.b, binary.BigEndian, &dst[i])
		if err != nil {
			log.Fatal(err)
		}
		b.b = b.b[n:]
	}
	fmt.Println("Read " + strconv.Itoa(length) + " bytes: " + string(dst))
	return b
}

func (b *BuffStruct) DecodeUint32Array(dst []uint32) *BuffStruct {
	length := len(dst)
	for i := 0; i < length; i++ {
		n, err := binary.Decode(b.b, binary.BigEndian, &dst[i])
		if err != nil {
			log.Fatal(err)
		}
		b.b = b.b[n:]
	}
	return b
}

func (b *BuffStruct) DecodeRuneArray(dst []rune) *BuffStruct {
	length := len(dst)
	for i := 0; i < length; i++ {
		n, err := binary.Decode(b.b, binary.BigEndian, &dst[i])
		if err != nil {
			log.Fatal(err)
		}
		b.b = b.b[n:]
	}
	fmt.Println("Read " + strconv.Itoa(length*4) + " bytes: " + RuneToString(dst))
	return b
}

func (b *BuffStruct) DecodeArray(dst any) *BuffStruct {
	if reflect.TypeOf(dst).Elem().Kind() == reflect.Uint8 {
		return b.DecodeByteArray(dst.([]byte))
	} else if reflect.TypeOf(dst).Elem().Kind() == reflect.Uint32 {
		return b.DecodeUint32Array(dst.([]uint32))
	}
	return b
}

func (b *BuffStruct) DecodeTokenType(dst *TokenType) *BuffStruct {
	n, err := binary.Decode(b.b, binary.BigEndian, dst)
	if err != nil {
		log.Fatal(err)
	}
	b.b = b.b[n:]
	return b
}
