package mostcommon

import (
	"errors"
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
func (m *Most) Chocolate() (Common, error) {
	if len(m.ListCommon) == 0 {
		return Common{"@", -1}, errors.New("sorry the chocolate wasnt shaked")
	}
	return m.ListCommon[0], nil
}

// Strawberry -
func (m *Most) Strawberry() (Common, error) {
	if len(m.ListCommon) == 1 {
		return Common{"@", -1}, errors.New("sorry the strawberry wasnt shaked")
	}
	return m.ListCommon[1], nil
}

// Vanilla -
func (m *Most) Vanilla() (Common, error) {
	if len(m.ListCommon) < 2 {
		return Common{"@", -1}, errors.New("sorry the vanilla wasnt shaked")
	}
	return m.ListCommon[2], nil
}
