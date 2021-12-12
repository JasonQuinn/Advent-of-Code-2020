package utils

type MultiMap struct {
	Value map[string]map[string]bool
}

func NewMultiMap() MultiMap {
	return MultiMap{Value: make(map[string]map[string]bool)}
}

func (m MultiMap) GetValues(key string) map[string]bool {
	return m.Value[key]
}

func (m MultiMap) AddValue(key string, value string) {
	if m.Value[key] == nil {
		m.Value[key] = make(map[string]bool)
	}
	m.Value[key][value] = true
}
