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

