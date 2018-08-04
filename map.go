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

func (m Map) Set(key string, val interface{}) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.data[key] = val
}

func (m Map) Get(key string) interface{} {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.data[key]
}

func (m Map) String() string {
	return JSONString(m.data)
}

func (m Map) MarshalJSON() ([]byte, error) {
	return []byte(JSONString(m.data)), nil
}

func (m Map) UnmarshalJSON(byt []byte) error {
	return json.Unmarshal(byt, &m.data)
}
