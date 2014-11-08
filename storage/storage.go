package storage

type MemoryStorage struct {
  keys map[string]string
}

func NewMemoryStorage() (*MemoryStorage) {
  storage := new(MemoryStorage)
  storage.keys = make(map[string]string)
  return storage
}

func (*MemoryStorage) Ping() (string) {
  return "PONG";
}

func (self *MemoryStorage) Get(key string) (string, bool) {
  value, ok := self.keys[key];
  return value, ok;
}

func (self *MemoryStorage) Set(key string, value string) (string) {
  self.keys[key] = value;
  return value;
}

