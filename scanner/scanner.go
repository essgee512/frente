package scanner

import (
	"bufio"
	"bytes"
	"io"
)

// Token represents a lexical token.
type Token int

const (
	ILLEGAL Token = iota
	EOF

	DELIM

	EOL
	CR
	NL

	// <space> | <tab>
	WS

	// <non-empty-space>
	// !EOL && !CR && !NL && !WS
	NES
)

// TODO; remove before "release"
func (tok Token) String() string {
	return [...]string{"ILLEGAL", "EOF", "DELIM", "EOL", "CR", "NL", "WS", "NES"}[tok]
}

const (
	eof   = "" // peek returns EOF as <empty-string>
	eol   = "\r\n"
	cr    = "\r"
	nl    = "\n"
	space = " "
	tab   = "\t"
)

type Scanner struct {
	r     *bufio.Reader
	delim string
}

func NewScanner(r io.Reader, delim string) *Scanner {
	return &Scanner{
		r:     bufio.NewReader(r),
		delim: delim,
	}
}

func (s *Scanner) Scan() (Token, string) {

	if s.isNext(eof) {

		return EOF, eof

	} else if s.isNext(s.delim) {

		return DELIM, s.readN(len(s.delim))

	} else if s.isNext(eol) {

		return EOL, s.readContiguous(eol)

	} else if s.isNext(cr) {

		return CR, s.readContiguous(cr)

	} else if s.isNext(nl) {

		return NL, s.readContiguous(nl)

	} else if s.isNextWS() {

		return WS, s.readWS()

	} else if s.isNextNES() {

		return NES, s.readNES()

	}

	return ILLEGAL, s.readN(1)
}

func (s *Scanner) readN(n int) string {
	var buf bytes.Buffer
	for i := 0; i < n; i++ {
		buf.WriteRune(s.read())
	}
	return buf.String()
}

func (s *Scanner) readContiguous(str string) string {
	var buf bytes.Buffer
	buf.WriteRune(s.read())
	for {
		if s.isNext(str) {
			for i := 0; i < len(str); i++ {
				buf.WriteRune(s.read())
			}
		}
		break
	}
	return buf.String()
}

func (s *Scanner) isNext(str string) bool {
	n := len(str)
	if str == eof {
		n = 1
	}
	peek, _ := s.r.Peek(n)
	if string(peek) == str {
		return true
	}
	return false
}

func (s *Scanner) isNextWS() bool {
	if s.isNext(space) || s.isNext(tab) {
		return true
	}
	return false
}

func (s *Scanner) readWS() string {
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	for {
		if s.isNextWS() {
			buf.WriteRune(s.read())
		} else {
			break
		}
	}
	return buf.String()
}

func (s *Scanner) isNextNES() bool {
	if !s.isNext(eof) &&
		!s.isNext(s.delim) &&
		!s.isNext(eol) &&
		!s.isNext(cr) &&
		!s.isNext(nl) &&
		!s.isNextWS() {
		return true
	}
	return false
}

func (s *Scanner) readNES() string {
	var buf bytes.Buffer
	buf.WriteRune(s.read())
	for {
		if s.isNextNES() {
			buf.WriteRune(s.read())
		} else {
			break
		}
	}
	return buf.String()
}

func (s *Scanner) read() rune {
	ch, _, _ := s.r.ReadRune()
	return ch
}
