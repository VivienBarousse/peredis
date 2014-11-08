package storage

import "testing"

func TestGet(t *testing.T) {
  s := NewMemoryStorage()

  s.Set("foo", "bar")
  value, ok := s.Get("foo")

  if value != "bar" {
    t.Errorf("Expected `value` to be \"foo\", got \"%v\"", value)
  }
  if ok != true {
    t.Errorf("Expected `ok` to be \"true\", got \"%v\"", ok)
  }
}

func TestGetNoValue(t *testing.T) {
  s := NewMemoryStorage();

  value, ok := s.Get("foo")

  if value != "" {
    t.Errorf("Expected `value` to be \"\", got \"%v\"", value)
  }
  if ok != false {
    t.Errorf("Expected `ok` to be \"false\", got \"%v\"", ok)
  }
}

func TestGetEmptyString(t *testing.T) {
  s := NewMemoryStorage();
  s.Set("foo", "")

  value, ok := s.Get("foo")
  if value != "" {
    t.Errorf("Expected `value` to be \"\", got \"%v\"", value)
  }
  if ok != true {
    t.Errorf("Expected `ok` to be \"true\", got \"%v\"", ok)
  }
}

func TestSet(t *testing.T) {
  s := NewMemoryStorage();

  s.Set("foo", "bar")

  value, _ := s.Get("foo")
  if value != "bar" {
    t.Errorf("Expected `value` to be \"bar\", got \"%v\"", value)
  }
}

func TestSetValueAlreadyExists(t *testing.T) {
  s := NewMemoryStorage();

  s.Set("foo", "bar")
  s.Set("foo", "baz")

  value, _ := s.Get("foo")
  if value != "baz" {
    t.Errorf("Expected `value` to be \"baz\", got \"%v\"", value)
  }
}

