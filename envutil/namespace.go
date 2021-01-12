package envutil

import (
	"os"
	"strconv"
	"strings"
)

// Namespace is a binder which is used for binding environment variables.
type Namespace struct {
	s string
}

func (n *Namespace) new(s string) *Env {
	ss := []string{n.s, strings.ReplaceAll(s, " ", "_")}
	return &Env{Name: strings.ToUpper(strings.Join(ss, "_"))}
}

// BindString binds string into ptr with a optional default value.
func (n *Namespace) BindString(name string, ptr *string, def ...string) *Env {
	e := n.new(name)
	val, ok := os.LookupEnv(e.Name)
	if ok {
		e.Value = val
		goto BIND
	}
	if len(def) > 0 {
		e.Value = def[0]
	}

BIND:
	*ptr = e.Value
	return e
}

// BindInt binds integer into ptr with a optional default value.
func (n *Namespace) BindInt(name string, ptr *int64, def ...int64) *Env {
	e := n.new(name)
	val, ok := os.LookupEnv(e.Name)
	if ok {
		e.Value = val
		goto BIND
	}
	if len(def) > 0 {
		e.Value = strconv.FormatInt(def[0], 10)
	}

BIND:
	i, err := strconv.ParseInt(e.Value, 10, 64)
	if err != nil {
		if len(def) > 0 {
			*ptr = def[0]
		}
		return e
	}
	*ptr = i
	return e
}

// BindUint binds unassigned integer into ptr with a optional default value.
func (n *Namespace) BindUint(name string, ptr *uint64, def ...uint64) *Env {
	e := n.new(name)
	val, ok := os.LookupEnv(e.Name)
	if ok {
		e.Value = val
		goto BIND
	}
	if len(def) > 0 {
		e.Value = strconv.FormatUint(def[0], 10)
	}

BIND:
	i, err := strconv.ParseUint(e.Value, 10, 64)
	if err != nil {
		if len(def) > 0 {
			*ptr = def[0]
		}
		return e
	}
	*ptr = i
	return e
}

// BindFloat binds float into ptr with a optional default value.
func (n *Namespace) BindFloat(name string, ptr *float64, def ...float64) *Env {
	e := n.new(name)
	val, ok := os.LookupEnv(e.Name)
	if ok {
		e.Value = val
		goto BIND
	}
	if len(def) > 0 {
		e.Value = strconv.FormatFloat(def[0], 'f', -1, 64)
	}

BIND:
	i, err := strconv.ParseFloat(e.Value, 64)
	if err != nil {
		if len(def) > 0 {
			*ptr = def[0]
		}
		return e
	}
	*ptr = i
	return e
}

// BindBool binds boolean into ptr with a optional default value.
func (n *Namespace) BindBool(name string, ptr *bool, def ...bool) *Env {
	e := n.new(name)
	val, ok := os.LookupEnv(e.Name)
	if ok {
		e.Value = val
		goto BIND
	}
	if len(def) > 0 {
		e.Value = strconv.FormatBool(def[0])
	}

BIND:
	switch strings.ToLower(strings.TrimSpace(e.Value)) {
	case "1", "true":
		*ptr = true
	case "0", "false":
		*ptr = false
	default:
		if len(def) > 0 {
			*ptr = def[0]
		}
	}
	return e
}

// BindFunc binds value with given fn.
func (n *Namespace) BindFunc(name string, fn EnvBindFunc) *Env {
	e := n.new(name)
	val, ok := os.LookupEnv(e.Name)
	e.Value = fn(val, ok)
	return e
}

// NewNamespace defines a new namespace of environment variable.
func NewNamespace(s string) *Namespace {
	return &Namespace{strings.ToUpper(strings.ReplaceAll(s, " ", "_"))}
}

// EnvBindFunc is a function for binding value into variables. Applied value
// must be returned.
type EnvBindFunc func(value string, exists bool) string
