package src

type BinFile struct {
	Header  BinFileHeader
	Content []TokenSet
}

type BinFileHeader struct {
	Sentinel         [4]byte
	Version          [3]byte
	TokenSetCount    uint16
	TokenSetHeaderSz byte
}

type TokenSet struct {
	Header TokenSetHeader
	Tokens []Token
}

type TokenSetHeader struct {
	Version        [3]byte
	FilenameLength uint16
	Filename       []rune
	TokenCount     uint32
}
