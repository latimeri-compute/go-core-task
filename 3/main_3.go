package main

type StringIntMap struct {
	innerMap map[string]int
}

func (m *StringIntMap) Add(key string, val int) {
	m.innerMap[key] = val
}

func (m *StringIntMap) Remove(key string) {
	delete(m.innerMap, key)
}

func (m *StringIntMap) Copy() map[string]int {
	return map[string]int{}
}

func (m *StringIntMap) Exists(key string) bool {
	_, exists := m.innerMap[key]
	return exists
}

func (m *StringIntMap) Get(key string) (int, bool) {
	val, exists := m.innerMap[key]
	return val, exists
}
