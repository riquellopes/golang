package mostcommon

import (
	"sort"
)

// Common -
type Common struct {
	Key   string
	Value int
}

// Most -
type Most struct {
	Char       string
	ListCommon []Common
}

// Shake -
func (m *Most) Shake() {
	dict := make(map[string]int)

	for _, char := range m.Char {
		if _, ok := dict[string(char)]; ok {
			dict[string(char)]++
		} else {
			dict[string(char)] = 1
		}
	}

	for key, value := range dict {
		m.ListCommon = append(m.ListCommon, Common{key, value})
	}

	sort.Slice(m.ListCommon, func(i int, j int) bool {
		return m.ListCommon[i].Value > m.ListCommon[j].Value
	})
}

// Chocolate -
func (m *Most) Chocolate() Common {
	return m.ListCommon[0]
}

// Strawberry -
func (m *Most) Strawberry() Common {
	return m.ListCommon[1]
}

// Vanilla -
func (m *Most) Vanilla() Common {
	return m.ListCommon[2]
}
