package src

type Buffer struct{}

/*


func ReadToken(data []byte) (Token, error) {
	var t Token
	var err error
	b := new(src.BuffStruct)
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

func (b *BuffStruct) DecodeTokenType(dst *tokens.TokenType) *BuffStruct {
	n, err := binary.Decode(b.b, binary.BigEndian, dst)
	if err != nil {
		log.Fatal(err)
	}
	b.b = b.b[n:]
	return b
}
*/
