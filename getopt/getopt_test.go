package getopt

import (
	"testing"
)

func TestParse(t *testing.T) {
	t.Run("NoArgs", func(t *testing.T) {
		var o OptSet

		var (
			i int64
			b bool
			s string
		)

		o.Int(&i, 'i', 42, "int")
		o.Bool(&b, 'b', false, "bool")
		o.String(&s, 's', "foo", "string")

		if err := o.Parse(nil); err != nil {
			t.Fatal(err)
		}

		if i != 42 {
			t.Fatal("wrong value")
		}
		if b != false {
			t.Fatal("wrong value")
		}
		if s != "foo" {
			t.Fatal("wrong value")
		}

	})

	t.Run("SpecialCases", func(t *testing.T) {
		var o OptSet

		var (
			i int64
			b bool
			s string
		)

		o.Int(&i, 'i', 42, "int")
		o.Bool(&b, 'b', false, "bool")
		o.String(&s, 's', "foo", "string")

		args := []string{"", "-"}
		if err := o.Parse(args); err != nil {
			t.Fatal(err)
		}

		expectedArgs := []string{"-"}
		actualArgs := o.Args()
		if len(actualArgs) != len(expectedArgs) {
			t.Fatal("unexpected len of args")
		}

		for i := range actualArgs {
			if actualArgs[i] != expectedArgs[i] {
				t.Fatal("expected:", expectedArgs[i], "have:", actualArgs[i])
			}
		}

		if i != 42 {
			t.Fatal("wrong value")
		}
		if b != false {
			t.Fatal("wrong value")
		}
		if s != "foo" {
			t.Fatal("wrong value")
		}
	})

	t.Run("UnknownOption", func(t *testing.T) {
		var o OptSet

		args := []string{"-f23"}

		if err := o.Parse(args); err == nil {
			t.Fatal("expected error")
		}
	})

	t.Run("Int", func(t *testing.T) {
		var o OptSet

		var v int64
		o.Int(&v, 'f', 42, "usage")

		args := []string{"-f23"}

		if err := o.Parse(args); err != nil {
			t.Fatal(err)
		}

		if v != 23 {
			t.Fatal("wrong value")
		}
	})

	t.Run("IntWithoutValue", func(t *testing.T) {
		var o OptSet

		var v int64
		o.Int(&v, 'f', 42, "usage")

		args := []string{"-f"}

		if err := o.Parse(args); err == nil {
			t.Fatal("expected error")
		}
	})

	t.Run("IntWithInvalidValue", func(t *testing.T) {
		var o OptSet

		var v int64
		o.Int(&v, 'f', 42, "usage")

		args := []string{"-fabc"}

		if err := o.Parse(args); err == nil {
			t.Fatal("expected error")
		}
	})

	t.Run("IntDuplicate", func(t *testing.T) {
		var o OptSet

		var v int64
		o.Int(&v, 'f', 42, "usage")

		defer func() {
			r := recover()
			if r == nil {
				t.Fatal("expected panic")
			}
		}()

		var v2 int64
		o.Int(&v2, 'f', 42, "usage")
	})

	t.Run("Bool", func(t *testing.T) {
		var o OptSet

		var v bool
		o.Bool(&v, 'f', false, "usage")

		args := []string{"-f"}

		if err := o.Parse(args); err != nil {
			t.Fatal(err)
		}

		if v != true {
			t.Fatal("wrong value")
		}
	})

	t.Run("BoolWithValue", func(t *testing.T) {
		var o OptSet

		var v bool
		o.Bool(&v, 'f', false, "usage")

		args := []string{"-ftrue"}

		if err := o.Parse(args); err == nil {
			t.Fatal("expected error")
		}
	})

	t.Run("BoolDuplicate", func(t *testing.T) {
		var o OptSet

		var v int64
		o.Int(&v, 'f', 42, "usage")

		defer func() {
			r := recover()
			if r == nil {
				t.Fatal("expected panic")
			}
		}()

		var v2 bool
		o.Bool(&v2, 'f', false, "usage")
	})

	t.Run("String", func(t *testing.T) {
		var o OptSet

		var v string
		o.String(&v, 'f', "foo", "usage")

		args := []string{"-fbar"}

		if err := o.Parse(args); err != nil {
			t.Fatal(err)
		}

		if v != "bar" {
			t.Fatal("wrong value")
		}
	})

	t.Run("StringWithoutValue", func(t *testing.T) {
		var o OptSet

		var v string
		o.String(&v, 'f', "foo", "usage")

		args := []string{"-f"}

		if err := o.Parse(args); err == nil {
			t.Fatal("expected error")
		}
	})

	t.Run("StringDuplicate", func(t *testing.T) {
		var o OptSet

		var v int64
		o.Int(&v, 'f', 42, "usage")

		defer func() {
			r := recover()
			if r == nil {
				t.Fatal("expected panic")
			}
		}()

		var v2 string
		o.String(&v2, 'f', "foo", "usage")
	})

	t.Run("MixedArgsAndOptions", func(t *testing.T) {
		var o OptSet

		var (
			i int64
			b bool
			s string
		)

		o.Int(&i, 'i', 42, "int")
		o.Bool(&b, 'b', false, "bool")
		o.String(&s, 's', "xxx", "string")

		args := []string{"-b", "foo", "-i23", "bar", "baz", "-syyy"}
		if err := o.Parse(args); err != nil {
			t.Fatal(err)
		}

		expectedArgs := []string{"foo", "bar", "baz"}
		actualArgs := o.Args()
		if len(actualArgs) != len(expectedArgs) {
			t.Fatal("unexpected len of args")
		}

		for i := range actualArgs {
			if actualArgs[i] != expectedArgs[i] {
				t.Fatal("expected:", expectedArgs[i], "have:", actualArgs[i])
			}
		}

		if i != 23 {
			t.Fatal("wrong value")
		}
		if b != true {
			t.Fatal("wrong value")
		}
		if s != "yyy" {
			t.Fatal("wrong value")
		}
	})

}
