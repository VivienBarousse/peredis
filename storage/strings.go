package storage

func (self *MemoryStorage) Get(key string) (string, bool) {
  value, ok := self.keys[key];
  return value, ok;
}

func (self *MemoryStorage) Set(key string, value string) (string) {
  self.keys[key] = value;
  return value;
}

func (self *MemoryStorage) Mset(arguments ...string) {
  if len(arguments) % 2 != 0 {
    panic("Invalid number of arguments")
  }

  for i := 0; i < len(arguments); i += 2 {
    self.Set(arguments[i], arguments[i + 1])
  }
}

