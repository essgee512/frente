package frente_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/essgee512/frente"
)

func TestSplit(t *testing.T) {
	const delim string = "---"
	var scenarios = []struct {
		in    []byte
		front []byte
		body  []byte
		err   error
	}{
		{
			// nominal case
			in:    []byte("---\nfront\n---\nbody"),
			front: []byte("front"),
			body:  []byte("body"),
			err:   nil,
		},
		{
			// body has CDL
			in:    []byte("---\nfront\n---\nbody\n---\n"),
			front: []byte("front"),
			body:  []byte("body\n---\n"),
			err:   nil,
		},
		{
			// body has CDL edge case
			in:    []byte("---\nfront\n---\n---\n"),
			front: []byte("front"),
			body:  []byte("---\n"),
			err:   nil,
		},
		{
			in:    []byte("  ---\nfront\n---\nbody"),
			front: []byte(""),
			body:  []byte(""),
			err:   fmt.Errorf("matter must begin with ODL := '---[NL]'."),
		},
		{
			in:    []byte("---  \nfront\n---\nbody"),
			front: []byte(""),
			body:  []byte(""),
			err:   fmt.Errorf("matter must begin with ODL := '---[NL]'."),
		},
		{
			in:    []byte("---\nfront\nmissing CDL\n"),
			front: []byte(""),
			body:  []byte(""),
			err:   fmt.Errorf("matter must have a CDL := '[NL]---[NL]'."),
		},
		{
			in:    []byte("---\nfront\n ---\nbody\n"),
			front: []byte(""),
			body:  []byte(""),
			err:   fmt.Errorf("matter must have a CDL := '[NL]---[NL]'."),
		},
		{
			in:    []byte("---\nfront\n--- \nbody\n"),
			front: []byte(""),
			body:  []byte(""),
			err:   fmt.Errorf("matter must have a CDL := '[NL]---[NL]'."),
		},
	}

	for i, expected := range scenarios {
		f, b, err := frente.Split(expected.in, delim)
		if !reflect.DeepEqual(f, expected.front) {
			t.Errorf("sc %d. front mismatch:\n  expected %q\n  got      %q", i+1, expected.front, f)
		}
		if !reflect.DeepEqual(b, expected.body) {
			t.Errorf("sc %d. body  mismatch:\n  expected %q\n  got      %q", i+1, expected.body, b)
		}
		if !reflect.DeepEqual(err, expected.err) {
			t.Errorf("sc %d. error mismatch:\n  expected %q\n  got      %q", i+1, expected.err, err)
		}
	}
}
