package lex

import (
	"bytes"
	"go/ast"
	"io"
	"sync"
)

type noCopy struct{}

func (*noCopy) Lock()   {}
func (*noCopy) Unlock() {}

type fragLen int // utf8 rune len sum

func (fl fragLen) Len() int {
	return int(fl)
}

type Fragmenter interface {
	io.ByteWriter
	io.RuneScanner
	io.Closer
	Next() *Fragment
	More() bool
}

type Fragger struct {
	src io.ByteScanner
	buf []byte
	cur int
	off int
	nc  noCopy
}

var fragBufs = sync.Pool{
	New: func() interface{} {
		return make([]byte, 16)
	},
}

func getBuf() []byte {
	return fragBufs.Get().([]byte)
}

func putBuf(b []byte) {
	b = b[:0]
	fragBufs.Put(b)
}

func NewFragger(src io.ByteScanner) *Fragger {
	f := &Fragger{src: src}
	f.buf = getBuf()
	return f
}

func (f *Fragger) Next() *Fragment {
	return nil
}

/*func (f *Fragger)

}*/

func (f *Fragger) Close() error {
	putBuf(f.buf)
	return nil
}

type Fragment struct {
	b    []rune
	bLen int
	nc   noCopy
}

var yeet = bytes.Buffer{}

type buffer interface {
	io.ReadWriter
	io.WriterTo
	io.ByteReader
	io.RuneScanner
	io.RuneReader
}

var _ ast.ArrayType

var _ buffer = &bytes.Buffer{}

type Reader struct {
	b   buffer
	off int64
}
