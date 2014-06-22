package devnull

// conforms to io.Writer
var Writer = devNull{}

type devNull struct{}

func (dn devNull) Write(p []byte) (n int, err error) {
  return len(p), nil
}
