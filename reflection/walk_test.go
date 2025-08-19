package reflection

import (
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {
	type Profile struct {
		Age  int
		City string
	}
	type Person struct {
		Name    string
		Profile Profile
	}

	cases := []struct {
		Name     string
		Input    interface{}
		Expected []string
	}{
		{
			"a struct with one field",
			struct {
				Name string
				City string
			}{"Chris", "Australia"},
			[]string{"Chris", "Australia"},
		},
		{
			"struct with non string field",
			struct {
				Name string
				Age  int
			}{"Chris", 33},
			[]string{"Chris"},
		},
		{
			"nested fields",
			Person{"Chris", struct {
				Age  int
				City string
			}{33, "London"}},
			[]string{"Chris", "London"},
		},
	}

	for _, test := range cases {
		got := []string{}
		t.Run(test.Name, func(t *testing.T) {
			walk(test.Input, func(input string) {
				got = append(got, input)
			})
		})
		if !reflect.DeepEqual(got, test.Expected) {
			t.Errorf("expected %v got %v", test.Expected, got)
		}
	}
}
