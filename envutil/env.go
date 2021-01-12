package envutil

import (
	"strconv"
	"strings"
)

// Env is a environment variable mapper.
type Env struct {
	Name  string
	Value string
}

func (e *Env) String() string {
	if strings.Contains(e.Value, "\"") {
		return e.Name + "=" + strconv.Quote(e.Value)
	}
	return e.Name + "=" + e.Value
}
