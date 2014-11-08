package storage

import "testing"

func TestIncr(t *testing.T) {
  s := NewMemoryStorage()
  s.Set("key", "11")

  returned_value := s.Incr("key")

  new_value, _ := s.Get("key")
  if returned_value != 12 {
    t.Errorf("Expected `returned_value` to be \"12\", got \"%v\"", returned_value)
  }
  if new_value != "12" {
    t.Errorf("Expected `new_value` to be \"12\", got \"%v\"", new_value)
  }
}

func TestIncrWithNonexistingKey(t *testing.T) {
  s := NewMemoryStorage()

  returned_value := s.Incr("key")

  new_value, _ := s.Get("key")
  if returned_value != 1 {
    t.Errorf("Expected `returned_value` to be \"1\", got \"%v\"", returned_value)
  }
  if new_value != "1" {
    t.Errorf("Expected `new_value` to be \"1\", got \"%v\"", new_value)
  }
}

