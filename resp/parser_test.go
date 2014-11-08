package resp

import "testing"
import "strings"

func TestReadBulkString(t *testing.T) {
  input := strings.NewReader("$6\r\nfoobar\r\n")
  parser := NewParser(input)

  bulk_string, err := parser.ReadBulkString()

  if bulk_string != "foobar" {
    t.Errorf("Expected `bulk_string` to be \"foobar\", was \"%v\"", bulk_string)
  }
  if err != nil {
    t.Errorf("Expected `err` to be \"nil\", was \"%v\"", err)
  }
}

func TestReadInlineCommand(t *testing.T) {
  input := strings.NewReader("SET foo bar\r\n")
  parser := NewParser(input)

  command, err := parser.ReadInlineCommand()
  if len(command) != 3 {
    t.Errorf("Expected `len(command)` to be \"3\", was \"%v\"", len(command))
  }
  if command[0] != "SET" {
    t.Errorf("Expected `command[0]` to be \"SET\", was \"%v\"", command[0])
  }
  if command[1] != "foo" {
    t.Errorf("Expected `command[1]` to be \"foo\", was \"%v\"", command[1])
  }
  if command[2] != "bar" {
    t.Errorf("Expected `command[2]` to be \"bar\", was \"%v\"", command[2])
  }
  if err != nil {
    t.Errorf("Expected `err` to be \"nil\", was \"%v\"", err)
  }
}

func TestReadObjectWithBulkString(t *testing.T) {
  input := strings.NewReader("$6\r\nfoobar\r\n")
  parser := NewParser(input)

  object, err := parser.ReadObject()

  if object != "foobar" {
    t.Errorf("Expected `object` to be \"foobar\", was \"%v\"", object)
  }
  if err != nil {
    t.Errorf("Expected `err` to be \"nil\", was \"%v\"", err)
  }
}

func TestReadObjectWithInlineCommand(t *testing.T) {
  input := strings.NewReader("SET foo bar\r\n")
  parser := NewParser(input)

  object, err := parser.ReadObject()
  command := object.([]string)
  if len(command) != 3 {
    t.Errorf("Expected `len(command)` to be \"3\", was \"%v\"", len(command))
  }
  if command[0] != "SET" {
    t.Errorf("Expected `command[0]` to be \"SET\", was \"%v\"", command[0])
  }
  if command[1] != "foo" {
    t.Errorf("Expected `command[1]` to be \"foo\", was \"%v\"", command[1])
  }
  if command[2] != "bar" {
    t.Errorf("Expected `command[2]` to be \"bar\", was \"%v\"", command[2])
  }
  if err != nil {
    t.Errorf("Expected `err` to be \"nil\", was \"%v\"", err)
  }
}

