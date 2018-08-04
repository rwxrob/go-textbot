package textbot

import (
	"encoding/json"
	"sync"
)

type Map struct {
	mu   sync.Mutex
	data Data
}

func NewMap() Map {
	m := Map{}
	m.data = Data{}
	return m
}

func (m Map) Set(p ...interface{}) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.data.Set(p...)
}

func (m Map) Get(keys ...string) interface{} {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.data.Get(keys...)
}

func (m Map) String() string {
	return m.data.String()
}

func (m Map) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.data)
}

func (m Map) UnmarshalJSON(byt []byte) error {
	return json.Unmarshal(byt, &m.data)
}
