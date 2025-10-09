package tokens_test

import (
	"bytes"
	"io"
	"math"
	"reflect"
	"testing"

	. "LexGo/src/tokens"
)

type fields struct {
	TotalLength uint16
	ID          byte
	Type        TokenType
	ValueLength uint16
	Value       []byte
	Row         uint32
	Column      uint32
}

var testTokens []fields = []fields{
	{TotalLength: 0, ID: 0, Type: 0,
		ValueLength: 0, Value: []byte{},
		Row: 0, Column: 0},
	{
		TotalLength: 0, ID: 0, Type: 0,
		ValueLength: 70, Value: []byte{},
		Row: 0, Column: 0,
	}, {
		TotalLength: 0, ID: 0, Type: 0,
		ValueLength: 0, Value: []byte{65, 66, 67, 68},
		Row: 0, Column: 0,
	}, {
		TotalLength: 0, ID: 0, Type: 0,
		ValueLength: 2, Value: []byte{65, 66, 67, 68},
		Row: 0, Column: 0,
	}, {
		TotalLength: 0, ID: 0, Type: 0,
		ValueLength: 8, Value: []byte{65, 66, 67, 68},
		Row: 0, Column: 0,
	}, {
		TotalLength: 1, ID: 0, Type: 0,
		ValueLength: 8, Value: []byte{65, 66, 67, 68},
		Row: 0, Column: 0,
	}, {
		TotalLength: 2, ID: 0, Type: 0,
		ValueLength: 8, Value: []byte{65, 66, 67, 68},
		Row: 0, Column: 0,
	}, {
		TotalLength: 2, ID: 1, Type: 0,
		ValueLength: 8, Value: []byte{65, 66, 67, 68},
		Row: 0, Column: 0,
	}, {
		TotalLength: 2, ID: 2, Type: 0,
		ValueLength: 8, Value: []byte{65, 66, 67, 68},
		Row: 0, Column: 0,
	}, {
		TotalLength: 2, ID: 2, Type: 1,
		ValueLength: 8, Value: []byte{65, 66, 67, 68},
		Row: 0, Column: 0,
	}, {
		TotalLength: 2, ID: 2, Type: 2,
		ValueLength: 8, Value: []byte{65, 66, 67, 68},
		Row: 0, Column: 0,
	}, {
		TotalLength: 2, ID: 2, Type: 2,
		ValueLength: 8, Value: []byte{65, 66, 67, 68},
		Row: 1, Column: 0,
	}, {
		TotalLength: 2, ID: 2, Type: 2,
		ValueLength: 8, Value: []byte{65, 66, 67, 68},
		Row: 2, Column: 0,
	}, {
		TotalLength: 2, ID: 2, Type: 2,
		ValueLength: 8, Value: []byte{65, 66, 67, 68},
		Row: 0, Column: 1,
	},
	{
		TotalLength: 2, ID: 2, Type: 2,
		ValueLength: 8, Value: []byte{65, 66, 67, 68},
		Row: 0, Column: 2,
	},
}

func TestDecompileToken(t *testing.T) {
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name string
		args args
		want Token
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DecompileToken(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DecompileToken() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToken_Equals(t1 *testing.T) {
	tokens := make([]Token, len(testTokens))
	for i := 0; i < len(tokens); i++ {
		tokens[i] = Token{
			TotalLength: testTokens[i].TotalLength,
			ID:          testTokens[i].ID,
			Type:        testTokens[i].Type,
			ValueLength: testTokens[i].ValueLength,
			Value:       testTokens[i].Value,
			Row:         testTokens[i].Row,
			Column:      testTokens[i].Column,
		}
	}

	for i, _ := range tokens {
		for j, _ := range tokens {
			comp := tokens[i].Equals(tokens[j])
			if i == j {
				if !comp {
					t1.Errorf("Equals() = %v, want %v:\ni = %d: %v\nj = %d: %v",
						comp, !comp, i, tokens[i].String(), j, tokens[j].String())
					t1.Fail()
				} else {
					continue
				}
			}
			if comp {
				t1.Errorf("Equals() = %v, want %v:\ni = %d: %v\nj = %d: %v",
					comp, !comp, i, tokens[i].String(), j, tokens[j].String())
				t1.Fail()
			}
		}
	}
}

func TestToken_Equals2(t1 *testing.T) {
	token1 := Token{
		TotalLength: testTokens[0].TotalLength,
		ID:          testTokens[0].ID,
		Type:        testTokens[0].Type,
		ValueLength: testTokens[0].ValueLength,
		Value:       testTokens[0].Value,
		Row:         testTokens[0].Row,
		Column:      testTokens[0].Column,
	}
	token2 := Token{
		TotalLength: testTokens[0].TotalLength,
		ID:          testTokens[0].ID,
		Type:        testTokens[0].Type,
		ValueLength: testTokens[0].ValueLength,
		Value:       testTokens[0].Value,
		Row:         testTokens[0].Row,
		Column:      testTokens[0].Column,
	}
	if !token1.Equals(token2) {
		t1.Errorf("Equals2() = false, expected true.\nToken 1: %v\nToken 2: %v", token1, token2)
	}
}

func TestToken_Equals3(t1 *testing.T) {
	token1 := Token{
		TotalLength: 10,
		ID:          testTokens[0].ID,
		Type:        testTokens[0].Type,
		ValueLength: testTokens[0].ValueLength,
		Value:       testTokens[0].Value,
		Row:         testTokens[0].Row,
		Column:      testTokens[0].Column,
	}
	token2 := Token{
		TotalLength: 11,
		ID:          testTokens[0].ID,
		Type:        testTokens[0].Type,
		ValueLength: testTokens[0].ValueLength,
		Value:       testTokens[0].Value,
		Row:         testTokens[0].Row,
		Column:      testTokens[0].Column,
	}
	if token1.Equals(token2) {
		t1.Errorf("Equals2() = true, expected false.\nToken 1: %v\nToken 2: %v", token1, token2)
	}
}

func TestToken_Equals4(t1 *testing.T) {
	token1 := Token{
		TotalLength: testTokens[0].TotalLength,
		ID:          3,
		Type:        testTokens[0].Type,
		ValueLength: testTokens[0].ValueLength,
		Value:       testTokens[0].Value,
		Row:         testTokens[0].Row,
		Column:      testTokens[0].Column,
	}
	token2 := Token{
		TotalLength: testTokens[0].TotalLength,
		ID:          7,
		Type:        testTokens[0].Type,
		ValueLength: testTokens[0].ValueLength,
		Value:       testTokens[0].Value,
		Row:         testTokens[0].Row,
		Column:      testTokens[0].Column,
	}
	if token1.Equals(token2) {
		t1.Errorf("Equals2() = true, expected false.\nToken 1: %v\nToken 2: %v", token1, token2)
	}
}

func TestToken_Equals5(t1 *testing.T) {
	token1 := Token{
		TotalLength: testTokens[0].TotalLength,
		ID:          testTokens[0].ID,
		Type:        5,
		ValueLength: testTokens[0].ValueLength,
		Value:       testTokens[0].Value,
		Row:         testTokens[0].Row,
		Column:      testTokens[0].Column,
	}
	token2 := Token{
		TotalLength: testTokens[0].TotalLength,
		ID:          testTokens[0].ID,
		Type:        92,
		ValueLength: testTokens[0].ValueLength,
		Value:       testTokens[0].Value,
		Row:         testTokens[0].Row,
		Column:      testTokens[0].Column,
	}
	if token1.Equals(token2) {
		t1.Errorf("Equals2() = true, expected false.\nToken 1: %v\nToken 2: %v", token1, token2)
	}
}

func TestToken_Equals6(t1 *testing.T) {
	token1 := Token{
		TotalLength: testTokens[0].TotalLength,
		ID:          testTokens[0].ID,
		Type:        testTokens[0].Type,
		ValueLength: 77,
		Value:       testTokens[0].Value,
		Row:         testTokens[0].Row,
		Column:      testTokens[0].Column,
	}
	token2 := Token{
		TotalLength: testTokens[0].TotalLength,
		ID:          testTokens[0].ID,
		Type:        testTokens[0].Type,
		ValueLength: 6334,
		Value:       testTokens[0].Value,
		Row:         testTokens[0].Row,
		Column:      testTokens[0].Column,
	}
	if token1.Equals(token2) {
		t1.Errorf("Equals2() = true, expected false.\nToken 1: %v\nToken 2: %v", token1, token2)
	}
}

func TestToken_Equals7(t1 *testing.T) {
	token1 := Token{
		TotalLength: testTokens[0].TotalLength,
		ID:          testTokens[0].ID,
		Type:        testTokens[0].Type,
		ValueLength: testTokens[0].ValueLength,
		Value:       []byte{123, 124, 122},
		Row:         testTokens[0].Row,
		Column:      testTokens[0].Column,
	}
	token2 := Token{
		TotalLength: testTokens[0].TotalLength,
		ID:          testTokens[0].ID,
		Type:        testTokens[0].Type,
		ValueLength: testTokens[0].ValueLength,
		Value:       []byte{123, 124, 122, 199},
		Row:         testTokens[0].Row,
		Column:      testTokens[0].Column,
	}
	if token1.Equals(token2) {
		t1.Errorf("Equals2() = true, expected false.\nToken 1: %v\nToken 2: %v", token1, token2)
	}
}

func TestToken_Equals8(t1 *testing.T) {
	token1 := Token{
		TotalLength: testTokens[0].TotalLength,
		ID:          testTokens[0].ID,
		Type:        testTokens[0].Type,
		ValueLength: testTokens[0].ValueLength,
		Value:       []byte{123, 124, 122, 199},
		Row:         testTokens[0].Row,
		Column:      testTokens[0].Column,
	}
	token2 := Token{
		TotalLength: testTokens[0].TotalLength,
		ID:          testTokens[0].ID,
		Type:        testTokens[0].Type,
		ValueLength: testTokens[0].ValueLength,
		Value:       []byte{123, 124, 122},
		Row:         testTokens[0].Row,
		Column:      testTokens[0].Column,
	}
	if token1.Equals(token2) {
		t1.Errorf("Equals2() = true, expected false.\nToken 1: %v\nToken 2: %v", token1, token2)
	}
}

func TestToken_Equals9(t1 *testing.T) {
	token1 := Token{
		TotalLength: testTokens[0].TotalLength,
		ID:          testTokens[0].ID,
		Type:        testTokens[0].Type,
		ValueLength: testTokens[0].ValueLength,
		Value:       []byte{},
		Row:         testTokens[0].Row,
		Column:      testTokens[0].Column,
	}
	token2 := Token{
		TotalLength: testTokens[0].TotalLength,
		ID:          testTokens[0].ID,
		Type:        testTokens[0].Type,
		ValueLength: testTokens[0].ValueLength,
		Value:       []byte{123, 124, 122},
		Row:         testTokens[0].Row,
		Column:      testTokens[0].Column,
	}
	if token1.Equals(token2) {
		t1.Errorf("Equals2() = true, expected false.\nToken 1: %v\nToken 2: %v", token1, token2)
	}
}

func TestToken_Equals10(t1 *testing.T) {
	token1 := Token{
		TotalLength: testTokens[0].TotalLength,
		ID:          testTokens[0].ID,
		Type:        testTokens[0].Type,
		ValueLength: testTokens[0].ValueLength,
		Value:       []byte{123, 124, 122},
		Row:         testTokens[0].Row,
		Column:      testTokens[0].Column,
	}
	token2 := Token{
		TotalLength: testTokens[0].TotalLength,
		ID:          testTokens[0].ID,
		Type:        testTokens[0].Type,
		ValueLength: testTokens[0].ValueLength,
		Value:       []byte{},
		Row:         testTokens[0].Row,
		Column:      testTokens[0].Column,
	}
	if token1.Equals(token2) {
		t1.Errorf("Equals2() = true, expected false.\nToken 1: %v\nToken 2: %v", token1, token2)
	}
}

func TestToken_Equals11(t1 *testing.T) {
	token1 := Token{
		TotalLength: testTokens[0].TotalLength,
		ID:          testTokens[0].ID,
		Type:        testTokens[0].Type,
		ValueLength: testTokens[0].ValueLength,
		Value:       []byte{123, 124, 122},
		Row:         238,
		Column:      testTokens[0].Column,
	}
	token2 := Token{
		TotalLength: testTokens[0].TotalLength,
		ID:          testTokens[0].ID,
		Type:        testTokens[0].Type,
		ValueLength: testTokens[0].ValueLength,
		Value:       []byte{123, 124, 122},
		Row:         math.MaxUint32,
		Column:      testTokens[0].Column,
	}
	if token1.Equals(token2) {
		t1.Errorf("Equals2() = true, expected false.\nToken 1: %v\nToken 2: %v", token1, token2)
	}
}

func TestToken_Equals12(t1 *testing.T) {
	token1 := Token{
		TotalLength: testTokens[0].TotalLength,
		ID:          testTokens[0].ID,
		Type:        testTokens[0].Type,
		ValueLength: testTokens[0].ValueLength,
		Value:       []byte{123, 124, 122},
		Row:         testTokens[0].Row,
		Column:      238,
	}
	token2 := Token{
		TotalLength: testTokens[0].TotalLength,
		ID:          testTokens[0].ID,
		Type:        testTokens[0].Type,
		ValueLength: testTokens[0].ValueLength,
		Value:       []byte{123, 124, 122},
		Row:         testTokens[0].Row,
		Column:      math.MaxUint32,
	}
	if token1.Equals(token2) {
		t1.Errorf("Equals2() = true, expected false.\nToken 1: %v\nToken 2: %v", token1, token2)
	}
}

func TestToken_PrintTo(t1 *testing.T) {
	tests := []struct {
		name    string
		fields  fields
		wantOut string
	}{
		{"All zero", testTokens[0], "ID: 0; Type: 0; Value length: 0; Value: \n"},
		{"Positive length but empty value", testTokens[1], "ID: 0; Type: 0; Value length: 70; Value: \n"},
		{"Zero length but non-empty value", testTokens[2], "ID: 0; Type: 0; Value length: 0; Value: ABCD\n"},
		{"Smaller claimed length than actual", testTokens[3], "ID: 0; Type: 0; Value length: 2; Value: ABCD\n"},
		{"Longer claimed length than actual", testTokens[4], "ID: 0; Type: 0; Value length: 8; Value: ABCD\n"},
		{"Total length of 1", testTokens[5], "ID: 0; Type: 0; Value length: 8; Value: ABCD\n"},
		{"Total length of 2", testTokens[6], "ID: 0; Type: 0; Value length: 8; Value: ABCD\n"},
		{"ID of 1", testTokens[7], "ID: 1; Type: 0; Value length: 8; Value: ABCD\n"},
		{"ID of 2", testTokens[8], "ID: 2; Type: 0; Value length: 8; Value: ABCD\n"},
		{"Type of 1", testTokens[9], "ID: 2; Type: 1; Value length: 8; Value: ABCD\n"},
		{"Type of 2", testTokens[10], "ID: 2; Type: 2; Value length: 8; Value: ABCD\n"},
		{"Row 1", testTokens[11], "ID: 2; Type: 2; Value length: 8; Value: ABCD\n"},
		{"Row 2", testTokens[12], "ID: 2; Type: 2; Value length: 8; Value: ABCD\n"},
		{"Col 1", testTokens[13], "ID: 2; Type: 2; Value length: 8; Value: ABCD\n"},
		{"Col 2", testTokens[14], "ID: 2; Type: 2; Value length: 8; Value: ABCD\n"},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Token{
				TotalLength: tt.fields.TotalLength,
				ID:          tt.fields.ID,
				Type:        tt.fields.Type,
				ValueLength: tt.fields.ValueLength,
				Value:       tt.fields.Value,
				Row:         tt.fields.Row,
				Column:      tt.fields.Column,
			}
			out := &bytes.Buffer{}
			t.PrintTo(out)
			if gotOut := out.String(); gotOut != tt.wantOut {
				t1.Errorf("PrintTo() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}

func TestToken_String(t1 *testing.T) {
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"All zero", testTokens[0], "[ 0 | 0 | 0 | 0 |  | 0 | 0 ]"},
		{"Positive length but empty value", testTokens[1], "[ 0 | 0 | 0 | 70 |  | 0 | 0 ]"},
		{"Zero length but non-empty value", testTokens[2], "[ 0 | 0 | 0 | 0 | ABCD | 0 | 0 ]"},
		{"Smaller claimed length than actual", testTokens[3], "[ 0 | 0 | 0 | 2 | ABCD | 0 | 0 ]"},
		{"Longer claimed length than actual", testTokens[4], "[ 0 | 0 | 0 | 8 | ABCD | 0 | 0 ]"},
		{"Total length of 1", testTokens[5], "[ 1 | 0 | 0 | 8 | ABCD | 0 | 0 ]"},
		{"Total length of 2", testTokens[6], "[ 2 | 0 | 0 | 8 | ABCD | 0 | 0 ]"},
		{"ID of 1", testTokens[7], "[ 2 | 1 | 0 | 8 | ABCD | 0 | 0 ]"},
		{"ID of 2", testTokens[8], "[ 2 | 10 | 0 | 8 | ABCD | 0 | 0 ]"},
		{"Type of 1", testTokens[9], "[ 2 | 10 | 1 | 8 | ABCD | 0 | 0 ]"},
		{"Type of 2", testTokens[10], "[ 2 | 10 | 2 | 8 | ABCD | 0 | 0 ]"},
		{"Row 1", testTokens[11], "[ 2 | 10 | 2 | 8 | ABCD | 1 | 0 ]"},
		{"Row 2", testTokens[12], "[ 2 | 10 | 2 | 8 | ABCD | 2 | 0 ]"},
		{"Col 1", testTokens[13], "[ 2 | 10 | 2 | 8 | ABCD | 0 | 1 ]"},
		{"Col 2", testTokens[14], "[ 2 | 10 | 2 | 8 | ABCD | 0 | 2 ]"},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Token{
				TotalLength: tt.fields.TotalLength,
				ID:          tt.fields.ID,
				Type:        tt.fields.Type,
				ValueLength: tt.fields.ValueLength,
				Value:       tt.fields.Value,
				Row:         tt.fields.Row,
				Column:      tt.fields.Column,
			}
			if got := t.String(); got != tt.want {
				t1.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToken_Write(t1 *testing.T) {
	tests := []struct {
		name             string
		fields           fields
		wantW            string
		wantBytesWritten int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &Token{
				TotalLength: tt.fields.TotalLength,
				ID:          tt.fields.ID,
				Type:        tt.fields.Type,
				ValueLength: tt.fields.ValueLength,
				Value:       tt.fields.Value,
				Row:         tt.fields.Row,
				Column:      tt.fields.Column,
			}
			w := &bytes.Buffer{}
			gotBytesWritten := t.Write(w)
			if gotW := w.String(); gotW != tt.wantW {
				t1.Errorf("Write() gotW = %v, want %v", gotW, tt.wantW)
			}
			if gotBytesWritten != tt.wantBytesWritten {
				t1.Errorf("Write() = %v, want %v", gotBytesWritten, tt.wantBytesWritten)
			}
		})
	}
}
