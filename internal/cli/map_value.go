package cli

import (
	"strings"
)

// MapValue implements flag.Value for map options set
// -c name1=value1 -c name2=value2
type MapValue struct {
	values map[string]string
}

// NewMapValue creates empty MapValue
func NewMapValue() *MapValue {
	return &MapValue{make(map[string]string)}
}

// Values returns all values map
func (m *MapValue) Values() map[string]string {
	return m.values
}

// String implements flag.Value.String() method
func (m *MapValue) String() string {
	return ""
}

// Set implements flag.Value.Set() method
func (m *MapValue) Set(arg string) error {
	pair := strings.SplitN(arg, "=", 2)
	value := ""
	if len(pair) > 1 {
		value = pair[1]
	}
	m.values[pair[0]] = value
	return nil
}
