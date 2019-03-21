package frente

import (
  "bufio"
  "bytes"
  "fmt"
  "io"
)

const (
  eof   = ""    // peek returns empty string at EOF
  nl    = "\n"
)

func Split(matter []byte, delim string) ([]byte, []byte, error) {
  p := newParser(bufio.NewReader(bytes.NewBuffer(matter)), delim)
  openDelim := delim + nl
  closeDelim := nl + delim + nl
  front := bytes.NewBufferString("")
  body := bytes.NewBufferString("")
  // var front, body bytes.Buffer

  // matter must begin with an opening delimeter
  if !p.isNext(openDelim) {
    found, _ := p.r.Peek(len(openDelim))
    expected := delim + nl
    return front.Bytes(), body.Bytes(), fmt.Errorf("matter must begin with the delimiter. found %q, expected %q", found, expected)
  }

  // handle case of empty front
  if p.isNext(openDelim + openDelim) {        // 2 contiguous openDelims implies and empty front
    p.readN(2 * len(openDelim))               // consume delims
    b, _ := p.readUpto(eof)                   // capture body
    body.WriteString(b)
    return front.Bytes(), body.Bytes(), nil   // front is empty
  }

  p.readN(len(openDelim))                     // consume opening delimiter
  f, err := p.readUpto(closeDelim)            // capture front
  front.WriteString(f)                        // write front

  // matter must have a closing delimeter
  if err != nil {
    return front.Bytes(), body.Bytes(), fmt.Errorf("closing delimeter not found")
  }

  p.readN(len(closeDelim))                    // consume closing delimeter
  b, _ := p.readUpto(eof)                     // capture body
  body.WriteString(b)                         // write body

  return front.Bytes(), body.Bytes(), nil     // success
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
    return buf.String(), fmt.Errorf("%q not found", str)
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
