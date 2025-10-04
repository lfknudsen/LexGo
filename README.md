# LexGo

Simple lexer and lexer generator written in Go.

It takes two things as input: a **ruleset** and a **code file**.
It produces a binary file as output.

## Usage

Build with `go build`. Then:

```
$ ./LexGo <options> <filename(s)>                       Tokenise code file(s) based on a ruleset.
$ ./LexGo decode <filename>                             Pretty print contents of a binary file written by this programme. 

Options:
    --no-bom  -n                                        Do not use a BOM.
    --use-bom -u                                        Use a BOM (default)
    
    --endian -e     <endianness>                        Set the endianness.
                    little/little-endian
                    big/big-endian                      (default)
                    native/native-endian/machine
                    
    --rule -r       <ruleset filename>                  Specify ruleset filename (default is "ruleset.txt")
    
    --format -f     <format>                            Choose output format
                    bin/binary/b
                    plain/plaintext/text/p
    
    --output -o     <output filename>                   Set output filename.

```

### Example

```
$ ./LexGo -e big code.txt code2.txt
```

## Ruleset format

The ruleset format is described within the included example `ruleset.txt` (slightly truncated here):

```
/* Each line below is a separate rule.
 The format is <name><whitespace><regexp>.
 Note that the programme uses Google's RE2 engine, which is slightly limited in
 functionality in order to maintain O(n) time complexity;
 This means look-back and look-ahead is _not_ possible.

 Rules have higher priority than ones listed below them. If two rules would
 theoretically match the same text, whichever one is highest will be selected. */
PKG                             package
TYPE                            type
STRUCT                          struct
STRING                          "(?:(?:(?:(?:\\")|[^"])|\s)*[^\\])?"
STRING_SINGLE                   '(?:(?:(?:(?:\\')|[^'])|\s)*[^\\])?'
STRING_BACK                     `(?:(?:(?:(?:\\`)|[^`])|\s)*[^\\])?`
IDENT                           \p{L}+
INTEGER                         \d+
FLOAT                           \d+(\.\d+)?
# Remember to escape special symbols if you mean to use them as characters!
LPAREN                          \(
RPAREN                          \)
LBRace                          {
RBRACE                          }
OP_PLUS                         \+
OP_MINUS                        \-
OP_EQ                           \=
OP_DEQ                          \=\=
COMMENT_LINE_START              \#
COMMENT_BLOCK_START             \/\*
/* Some characters can be problematic, and do not follow normal rules of escaping with
 a backslash. You can enclose them in brackets instead. */
COMMENT_BLOCK_END               [*][/]
/* ?: instead of a name means the regex will not be captured, meaning no token will
 be created based on it. The regex output of the following will be (?:\s+) */
?:(none)                        \s+
MISTAKE                         .+
```

## Binary output format

The format is described below. Naturally, this programme is able to decode this format
itself; feel free to look at the source code for inspiration.\
The binary file contains a header and an array of "Token Sets". Each *Token Set* represents
a file (or translation unit, or however you wish to use/think of it), which 
contains an array of *Tokens*.

Below, I've preceded each line with one or two characters to conceptually categorise them (this does not
indicate how they were implemented, nor the author's opinions of how they ought to be);
```
 >  Struct/class
 +  Primitive value
:>  Array of structs/classes
:+  Array of primitive values
```

### Binary File

```
 + BOM (2 bytes); 0xFEFF if read as correct endianness; reader should switch if read as 0xFFFE
 > File Header (12 bytes)
:> File Content (Array of Token Sets)
```

#### Binary File Header

```
:+ Sentinel (5 bytes, ASCII characters; note that they are not zero-terminated): L E X G O
:+ Version (3 bytes. Semantic versioning. Each an unsigned integer): <major> <minor> <patch>
 + TokenSetCount (4 bytes, signed integer); Number of *Token Sets* in the File Content.
```

#### Token Set

```
 > Token Set Header
:> Tokens (array, each element represents a Token)
```

##### Token Set Header

```
:+ Version (3 bytes): <major> <minor> <patch>
 + Token count (4 bytes, unsigned integer): Number of *Tokens* in this Token Set.
 + Filename Length (2 bytes, unsigned integer)
:+ Filename (byte array, not zero-terminated, number of elements contained in Filename Length).
```

##### Token

```
 + Total Length (2 bytes, unsigned integer); number of bytes in this token 
    (including the total length)
 + ID (1 byte): 0 = first rule in ruleset, 1 = second rule, and so on...
 + Type (1 byte, unsigned integer): Undefined.
 + Value Length (2 bytes, unsigned integer): Length of the byte array which contains the
    actual text which succesfully matched against a rule.
:+ Value (byte array, not zero-terminated, number of elements is equal to value of Value Length).
    This is encoded as plaintext.
 + Row (4 bytes, unsigned integer); the row for the first character of the token.
 + Column (4 bytes, unsigned integer); the column for the first character of the token.
```
