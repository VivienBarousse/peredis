package resp

import "testing"
import "bytes"

func TestWriteNil(t *testing.T) {
  output := new(bytes.Buffer)
  serializer := NewSerializer(output)

  serializer.WriteNil()

  if output.String() != "$-1\r\n" {
    t.Errorf("Expected `output` to be \"$-1\r\n\", was \"%v\"", output)
  }
}

func TestWriteInteger(t *testing.T) {
  output := new(bytes.Buffer)
  serializer := NewSerializer(output)

  serializer.WriteInteger(123)
  if output.String() != ":123\r\n" {
    t.Errorf("Expected `output` to be \"$6\r\nfoobar\r\n\", was \"%v\"", output)
  }
}

func TestWriteString(t *testing.T) {
  output := new(bytes.Buffer)
  serializer := NewSerializer(output)

  serializer.WriteString("foobar")
  if output.String() != "$6\r\nfoobar\r\n" {
    t.Errorf("Expected `output` to be \"$6\r\nfoobar\r\n\", was \"%v\"", output)
  }
}

func TestWriteEmptyString(t *testing.T) {
  output := new(bytes.Buffer)
  serializer := NewSerializer(output)

  serializer.WriteString("")
  if output.String() != "$0\r\n\r\n" {
    t.Errorf("Expected `output` to be \"$0\r\n\r\n\", was \"%v\"", output)
  }
}

