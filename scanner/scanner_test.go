package scanner_test

import (
	"strings"
	"testing"

	fsc "github.com/essgee512/frente/scanner"
)

func TestScan(t *testing.T) {
	const delim string = "---"
	var scenarios = []struct {
		in  string
		tok fsc.Token
		lit string
	}{
		{
			in:  "",
			tok: fsc.EOF,
			lit: "",
		},
		{
			in:  delim,
			tok: fsc.DELIM,
			lit: "---",
		},
		{
			in:  delim + delim,
			tok: fsc.DELIM,
			lit: "------",
		},
		{
			in:  "\r\n",
			tok: fsc.EOL,
			lit: "\r\n",
		},
		{
			in:  "\r\n\r\n\r\n",
			tok: fsc.EOL,
			lit: "\r\n\r\n\r\n",
		},
		{
			in:  "\r\r\r",
			tok: fsc.CR,
			lit: "\r\r\r",
		},
		{
			in:  "\n\n\n",
			tok: fsc.NL,
			lit: "\n\n\n",
		},
		{
			in:  " ",
			tok: fsc.WS,
			lit: " ",
		},
		{
			in:  "   ",
			tok: fsc.WS,
			lit: "   ",
		},
		{
			in:  "\t",
			tok: fsc.WS,
			lit: "\t",
		},
		{
			in:  "\t\t",
			tok: fsc.WS,
			lit: "\t\t",
		},
		{
			in:  " \t",
			tok: fsc.WS,
			lit: " \t",
		},
		{
			in:  `nes135\3:@#%^`,
			tok: fsc.NES,
			lit: `nes135\3:@#%^`,
		},
	}

	for i, sc := range scenarios {
		s := fsc.NewScanner(strings.NewReader(sc.in), delim)
		tok, lit := s.Scan()
		if sc.tok != tok {
			t.Errorf("%d. in: %q, mismatch: expected [%v, %q], got [%v, %q]", i, sc.in, sc.tok, sc.lit, tok, lit)
		}
	}
}
