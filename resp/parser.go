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

func (self *Parser) ReadBulkArray() ([]interface{}, error) {
  self.input.ReadByte()
  line, _ := self.input.ReadString('\n')
  length, _ := strconv.Atoi(line[:len(line) - 2])
  value := make([]interface{}, length)

  for i := 0; i < length; i++ {
    object, err := self.ReadObject()
    if err != nil {
      return make([]interface{}, 0), err
    }

    value[i] = object
  }

  return value, nil
}

func (self *Parser) ReadInlineCommand() ([]string, error) {
  line, _ := self.input.ReadString('\n')
  line = line[:len(line) - 2]
  return strings.Split(line, " "), nil
}

func (self *Parser) ReadObject() (interface{}, error) {
  read, err := self.input.Peek(1)
  if err != nil {
    return nil, err
  }

  next := read[0]
  if next == '$' {
    return self.ReadBulkString()
  } else if (next == '*') {
    return self.ReadBulkArray()
  } else {
    return self.ReadInlineCommand()
  }
}

