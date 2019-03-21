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
    { // empty front
      in:    []byte("---\n---\nbody"),
      front: []byte(""),
      body:  []byte("body"),
      err:   nil,
    },
    { // empty body
      in:    []byte("---\nfront\n---\n"),
      front: []byte("front"),
      body:  []byte(""),
      err:   nil,
    },
    { // empty front and empty body
      in:    []byte("---\n---\n"),
      front: []byte(""),
      body:  []byte(""),
      err:   nil,
    },
    {
      in:    []byte(" ---\n---\n"),
      front: []byte(""),
      body:  []byte(""),
      err:   fmt.Errorf(`matter must begin with the delimiter. found " ---", expected "---\n"`),
    },
    {
      in:    []byte("---front\n---\nbody"),
      front: []byte(""),
      body:  []byte(""),
      err:   fmt.Errorf(`the delimiter must be the only string on the line.`),
    },
    // {
    //   in:    []byte("---\nfront\nmissing\n"),
    //   front: []byte("front---\n"),
    //   body:  []byte(""),
    //   err:   fmt.Errorf("closing delimeter not found"),
    // },   
    // {
    //   in:    []byte("---\nfront---\n"),
    //   front: []byte("front---\n"),
    //   body:  []byte(""),
    //   err:   fmt.Errorf("closing delimeter not found"),
    // },

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










