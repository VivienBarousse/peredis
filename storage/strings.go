package storage

func (self *MemoryStorage) Get(key string) (string, bool) {
  value, ok := self.keys[key];
  return value, ok;
}

func (self *MemoryStorage) Set(key string, value string) (string) {
  self.keys[key] = value;
  return value;
}

