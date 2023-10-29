package getopt

import (
	"fmt"
	"strconv"
)

type OptSet struct {
	ints    []intOpt
	strings []stringOpt
	bools   []boolOpt

	args []string
}

type opt struct {
	usage string
	name  byte
	typ   optType
}

type intOpt struct {
	value *int64
	opt
}

type boolOpt struct {
	value *bool
	opt
}

type stringOpt struct {
	value *string
	opt
}

type optType uint8

const (
	optUndefined optType = iota
	optInt
	optBool
	optString
)

func (self *OptSet) Int(ptr *int64, name byte, value int64, usage string) {
	if i, _ := self.find(name); i != -1 {
		panic("option with the same name defined multiple times")
	}

	*ptr = value

	self.ints = append(self.ints, intOpt{
		opt: opt{
			name:  name,
			typ:   optInt,
			usage: usage,
		},
		value: ptr,
	})
}

func (self *OptSet) Bool(ptr *bool, name byte, value bool, usage string) {
	if i, _ := self.find(name); i != -1 {
		panic("option with the same name defined multiple times")
	}

	*ptr = value

	self.bools = append(self.bools, boolOpt{
		opt: opt{
			name:  name,
			typ:   optBool,
			usage: usage,
		},
		value: ptr,
	})
}

func (self *OptSet) String(ptr *string, name byte, value string, usage string) {
	if i, _ := self.find(name); i != -1 {
		panic("option with the same name defined multiple times")
	}

	*ptr = value

	self.strings = append(self.strings, stringOpt{
		opt: opt{
			name:  name,
			typ:   optString,
			usage: usage,
		},
		value: ptr,
	})
}

func (self *OptSet) Parse(args []string) error {
	self.args = self.args[:0]

	for _, arg := range args {
		if len(arg) == 0 {
			continue
		}
		if arg[0] == '-' && len(arg) > 1 {
			name := arg[1]

			i, t := self.find(name)
			if i == -1 {
				return fmt.Errorf("unknown option '%c'", name)
			}

			switch t {
			case optInt:
				if len(arg[2:]) == 0 {
					return fmt.Errorf("int options must have a value")
				}
				v, err := strconv.ParseInt(arg[2:], 0, 64)
				if err != nil {
					return err
				}
				*self.ints[i].value = v
			case optBool:
				if len(arg[2:]) != 0 {
					return fmt.Errorf("boolean options cannot have a value")
				}
				v := *self.bools[i].value
				*self.bools[i].value = !v
			case optString:
				if len(arg[2:]) == 0 {
					return fmt.Errorf("string options must have a value")
				}
				*self.strings[i].value = arg[2:]
			case optUndefined:
				panic("unreachable")
			default:
				panic("unreachable")
			}

		} else {
			self.args = append(self.args, arg)
		}
	}

	return nil
}

func (self OptSet) Args() []string {
	return self.args
}

func (self OptSet) find(name byte) (int, optType) {
	for i, v := range self.ints {
		if v.name == name {
			return i, optInt
		}
	}
	for i, v := range self.bools {
		if v.name == name {
			return i, optBool
		}
	}
	for i, v := range self.strings {
		if v.name == name {
			return i, optString
		}
	}
	return -1, optUndefined
}
