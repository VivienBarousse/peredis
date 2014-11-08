package storage

import "testing"

func TestPing(t *testing.T) {
  s := NewMemoryStorage()
  value := s.Ping();

  if value != "PONG" {
    t.Error("Expected `value` to be \"PONG\", got \"%v\"", value);
  }
}

