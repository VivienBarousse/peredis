package storage

import "strconv"

func (self *MemoryStorage) Incr(key string) (int) {
  value, ok := self.Get(key)
  if ok {
    new_value, _ := strconv.Atoi(value)
    new_value += 1
    self.Set(key, strconv.Itoa(new_value))
    return new_value
  } else {
    self.Set(key, "1")
    return 1
  }
}

