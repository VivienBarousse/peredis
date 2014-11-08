package resp

import "io"
import "bufio"
import "strings"
import "strconv"

type Parser struct {
  input *bufio.Reader
}

func NewParser(reader io.Reader) (*Parser) {
  p := new(Parser)
  p.input = bufio.NewReader(reader)
  return p
}

func (self *Parser) ReadBulkString() (string) {
  self.input.ReadByte()
  line, _ := self.input.ReadString('\n')
  length, _ := strconv.Atoi(line[:len(line) - 2])
  value := make([]byte, length)
  io.ReadFull(self.input, value)
  self.input.ReadString('\n')
  return string(value)
}

func (self *Parser) ReadInlineCommand() ([]string) {
  line, _ := self.input.ReadString('\n')
  line = line[:len(line) - 2]
  return strings.Split(line, " ")
}

func (self *Parser) ReadObject() (interface{}) {
  read, _ := self.input.Peek(1)
  next := read[0]
  if next == '$' {
    return self.ReadBulkString()
  } else {
    return self.ReadInlineCommand()
  }
}

