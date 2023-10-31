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

func (o *OptSet) Int(ptr *int64, name byte, value int64, usage string) {
	if i, _ := o.find(name); i != -1 {
		panic("option with the same name defined multiple times")
	}

	*ptr = value

	o.ints = append(o.ints, intOpt{
		opt: opt{
			name:  name,
			typ:   optInt,
			usage: usage,
		},
		value: ptr,
	})
}

func (o *OptSet) Bool(ptr *bool, name byte, value bool, usage string) {
	if i, _ := o.find(name); i != -1 {
		panic("option with the same name defined multiple times")
	}

	*ptr = value

	o.bools = append(o.bools, boolOpt{
		opt: opt{
			name:  name,
			typ:   optBool,
			usage: usage,
		},
		value: ptr,
	})
}

func (o *OptSet) String(ptr *string, name byte, value string, usage string) {
	if i, _ := o.find(name); i != -1 {
		panic("option with the same name defined multiple times")
	}

	*ptr = value

	o.strings = append(o.strings, stringOpt{
		opt: opt{
			name:  name,
			typ:   optString,
			usage: usage,
		},
		value: ptr,
	})
}

func (o *OptSet) Parse(args []string) error {
	o.args = o.args[:0]

	for _, arg := range args {
		if len(arg) == 0 {
			continue
		}
		if arg[0] == '-' && len(arg) > 1 {
			name := arg[1]

			i, t := o.find(name)
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
				*o.ints[i].value = v
			case optBool:
				if len(arg[2:]) != 0 {
					return fmt.Errorf("boolean options cannot have a value")
				}
				v := *o.bools[i].value
				*o.bools[i].value = !v
			case optString:
				if len(arg[2:]) == 0 {
					return fmt.Errorf("string options must have a value")
				}
				*o.strings[i].value = arg[2:]
			case optUndefined:
				panic("unreachable")
			default:
				panic("unreachable")
			}

		} else {
			o.args = append(o.args, arg)
		}
	}

	return nil
}

func (o OptSet) Args() []string {
	return o.args
}

func (o OptSet) find(name byte) (int, optType) {
	for i, v := range o.ints {
		if v.name == name {
			return i, optInt
		}
	}
	for i, v := range o.bools {
		if v.name == name {
			return i, optBool
		}
	}
	for i, v := range o.strings {
		if v.name == name {
			return i, optString
		}
	}
	return -1, optUndefined
}
