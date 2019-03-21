package frente

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
)

const (
	eof = "" // peek returns empty string at EOF
	nl  = "\n"
)

func Split(matter []byte, delim string) ([]byte, []byte, error) {
	p := newParser(bufio.NewReader(bytes.NewBuffer(matter)), delim)
	openDelim := delim + nl
	closeDelim := nl + delim + nl
	empty := []byte("")
	front := bytes.NewBufferString("")
	body := bytes.NewBufferString("")

	// matter must begin with an opening delimeter
	if !p.isNext(openDelim) {
		return empty, empty, fmt.Errorf("matter must begin with ODL := '%s[NL]'.", delim)
	}

	p.readN(len(openDelim))          // consume opening delimiter
	f, err := p.readUpto(closeDelim) // capture front
	front.WriteString(f)             // write front

	// matter must have a closing delimeter
	if err != nil {
		return empty, empty, fmt.Errorf("matter must have a CDL := '[NL]%s[NL]'.", delim)
	}

	p.readN(len(closeDelim)) // consume closing delimeter
	b, _ := p.readUpto(eof)  // capture body
	body.WriteString(b)      // write body

	return front.Bytes(), body.Bytes(), nil // success
}

type parser struct {
	r *bufio.Reader
}

func newParser(r io.Reader, delim string) *parser {
	return &parser{bufio.NewReader(r)}
}

func (p *parser) isNext(str string) bool {
	n := len(str)
	if str == eof {
		n = 1
	}
	peek, _ := p.r.Peek(n)
	if string(peek) == str {
		return true
	}
	return false
}

func (p *parser) readN(n int) string {
	var buf bytes.Buffer
	for i := 0; i < n; i++ {
		buf.WriteRune(p.read())
	}
	return buf.String()
}

func (p *parser) readUpto(str string) (string, error) {
	var buf bytes.Buffer
	if p.isNext(eof) {
		return eof, fmt.Errorf("%q not found", str)
	}
	buf.WriteRune(p.read())
	for {
		if p.isNext(eof) {
			break
		} else if p.isNext(str) {
			return buf.String(), nil
		}
		buf.WriteRune(p.read())
	}
	return buf.String(), fmt.Errorf("%q not found", str)
}

func (p *parser) read() rune {
	ch, _, _ := p.r.ReadRune()
	return ch
}

func (p *parser) unread() { _ = p.r.UnreadRune() }
