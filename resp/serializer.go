package resp

import "io"
import "fmt"

type Serializer struct {
  output io.Writer
}

func NewSerializer(output io.Writer) (*Serializer) {
  s := new(Serializer)
  s.output = output
  return s
}

func (self *Serializer) WriteNil() {
  self.output.Write([]byte("$-1\r\n"))
}

func (self *Serializer) WriteString(value string) {
  fmt.Fprintf(self.output, "$%d\r\n%s\r\n", len(value), value)
}

