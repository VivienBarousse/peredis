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

func (self *Parser) ReadBulkString() (string, error) {
  self.input.ReadByte()
  line, _ := self.input.ReadString('\n')
  length, _ := strconv.Atoi(line[:len(line) - 2])
  value := make([]byte, length)
  io.ReadFull(self.input, value)
  self.input.ReadString('\n')
  return string(value), nil
}

func (self *Parser) ReadInlineCommand() ([]string, error) {
  line, _ := self.input.ReadString('\n')
  line = line[:len(line) - 2]
  return strings.Split(line, " "), nil
}

func (self *Parser) ReadObject() (interface{}, error) {
  read, _ := self.input.Peek(1)
  next := read[0]
  if next == '$' {
    return self.ReadBulkString()
  } else {
    return self.ReadInlineCommand()
  }
}

