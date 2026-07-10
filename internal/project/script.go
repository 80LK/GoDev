package project

import (
	"fmt"

	"github.com/80LK/modlike"
	"github.com/mattn/go-shellwords"
)

type Script struct {
	Command string   `modlike:"command"`
	Args    []string `modlike:"args"`
	WorkDir *string  `modlike:"workdir"`
	Env     []string `modlike:"env"`
}

func parseFromString(input string, s *Script) error {
	parts, err := shellwords.Parse(input)
	if err != nil {
		return err
	}

	if len(parts) == 0 {
		return nil
	}

	s.Command = parts[0]
	s.Args = parts[1:]

	return nil
}

const (
	_COMMAND_KEY = "command"
	_ARGS_KEY    = "args"
	_WORKDIR_KEY = "workdir"
	_ENV_KEY     = "env"
)

func parseFromMap(value modlike.Value, s *Script) error {
	_map, _ := value.ToMap()

	var err error

	value = _map.Get(_COMMAND_KEY)
	if value == nil || value.Kind() == modlike.K_NONE {
		return notFoundKey(_COMMAND_KEY)
	}
	str, err := parseString(value, _COMMAND_KEY)
	if err != nil {
		return err
	}
	s.Command = str

	value = _map.Get(_ARGS_KEY)
	if value != nil {
		list, err := parseList(value, _ARGS_KEY)
		if err != nil {
			return err
		}
		s.Args = list
	}

	value = _map.Get(_WORKDIR_KEY)
	if value != nil {
		str, err := parseString(value, _WORKDIR_KEY)
		if err != nil {
			return err
		}
		s.WorkDir = &str
	}

	value = _map.Get(_ENV_KEY)
	if value != nil {
		list, err := parseList(value, _ENV_KEY)
		if err != nil {
			return err
		}
		s.Env = list
	}

	return nil
}

func parseList(value modlike.Value, key string) ([]string, error) {
	kind := value.Kind()
	if value.Kind() != modlike.K_LIST {
		return nil, missmatchType(key, kind, modlike.K_LIST)
	}
	list, _ := value.ToList()

	slice := []string{}
	for _, val := range list.All {
		kind := val.Kind()
		if kind != modlike.K_STRING {
			return nil, missmatchType(_ARGS_KEY+"[int]", kind, modlike.K_LIST)
		}

		str, _ := val.ToString()
		slice = append(slice, str.Get())
	}
	return slice, nil
}

func parseString(value modlike.Value, key string) (string, error) {
	kind := value.Kind()
	if kind != modlike.K_LIST && kind != modlike.K_STRING {
		return "", missmatchType(key, kind, modlike.K_STRING)
	}

	if kind == modlike.K_LIST {
		list, _ := value.ToList()
		return list.GetStringFirst()
	} else {
		str, _ := value.ToString()
		return str.Get(), nil
	}
}

func notFoundKey(key string) error {
	return fmt.Errorf("[Script] DecodeModlike: not found key %q", key)
}

func missmatchType(key string, got, expect modlike.Kind) error {
	return fmt.Errorf("[Script] DecodeModlike: missmatch type key %q. Expect %s, but got %s", key, expect, got)
}

func (s *Script) DecodeModlike(value modlike.Value) error {
	s.Args = []string{}
	s.Env = []string{}

	switch value.Kind() {
	case modlike.K_LIST:
		list, _ := value.ToList()
		str, err := list.GetStringFirst()
		if err != nil {
			return err
		}
		return parseFromString(str, s)
	case modlike.K_STRING:
		str, _ := value.ToString()
		return parseFromString(str.Get(), s)
	case modlike.K_MAP:
		if err := parseFromMap(value, s); err != nil {
			return err
		}
	}
	return nil
}
