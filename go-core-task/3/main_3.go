package main

type StringIntMap struct {
	data map[string]int
}

func (m *StringIntMap) Get(key string) (int, bool) {
	v, exist := m.data[key]
	return v, exist
}

func (m *StringIntMap) Add(key string, value int) {
	m.data[key] = value
}

func (m *StringIntMap) Remove(key string) {
	delete(m.data, key)
}

func (m *StringIntMap) Copy() map[string]int {
	copied := make(map[string]int, len(m.data))
	for key, value := range m.data {
		copied[key] = value
	}

	return copied
}

func (m *StringIntMap) Exists(key string) bool {
	_, exist := m.data[key]

	return exist
}

func New(data map[string]int) *StringIntMap {
	if data == nil {
		data = make(map[string]int)
	}
	return &StringIntMap{
		data: data,
	}
}
