package devnull

import (
  "testing"
)

type writerCase struct {
  s []byte
  n int
}

var writerCases = []writerCase{
  {[]byte(""), 0},
  {[]byte("a"), 1},
  {[]byte("xyz"), 3},
  {nil, 0},
}

func TestWriter(t *testing.T) {
  for _, wc := range writerCases {
    n, err := Writer.Write(wc.s)
    if err != nil {
      t.Errorf("TestWriter: failed Write(", wc.s, ") error = ", err)
    }
    if n != wc.n {
      t.Errorf("TestWriter: failed Write(", wc.s, ") ", wc.n, " != ", n)
    }
  }
}

const oneMiB = 1024 * 1024

var oneMiBBuf = make([]byte, oneMiB)

func BenchmarkWriter1MiB(b *testing.B) {
  b.SetBytes(oneMiB)
  for i := 0; i < b.N; i++ {
    Writer.Write(oneMiBBuf)
  }
}
