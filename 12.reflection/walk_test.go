package reflection

import (
	"reflect"
	"slices"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func assertContains(t testing.TB, haystack []string, needle string) {
	t.Helper()
	contains := false
	for _, x := range haystack {
		if x == needle {
			contains = true
			break
		}
	}
	if !contains {
		t.Errorf("expected %v to contain %q but it didn't", haystack, needle)
	}
}

func TestWalk(t *testing.T) {

	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"struct with one field",
			struct {
				Name string
			}{"Matt"},
			[]string{"Matt"},
		},
		{
			"struct with two fields",
			struct {
				Name string
				City string
			}{"Matt", "Stockholm"},
			[]string{"Matt", "Stockholm"},
		},
		{
			"struct with non string fields",
			struct {
				Name string
				Age  int
			}{"Matt", 49},
			[]string{"Matt"},
		},
		{
			"nested struct with non string fields",
			Person{"Matt", Profile{49, "Stockholm"}},
			[]string{"Matt", "Stockholm"},
		},
		{
			"pointers to things",
			&Person{"Matt", Profile{49, "Stockholm"}},
			[]string{"Matt", "Stockholm"},
		},
		{"slices",
			[]Profile{
				{33, "London"},
				{49, "Stockholm"},
			},
			[]string{"London", "Stockholm"},
		},
		{"arrays",
			[2]Profile{
				{33, "London"},
				{49, "Stockholm"},
			},
			[]string{"London", "Stockholm"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, expected %v", got, test.ExpectedCalls)
			}
		})
	}

	t.Run("with maps", func(t *testing.T) {
		aMap := map[string]string{
			"Cow": "Moo",
			"Cat": "Meaow",
		}

		var got []string

		walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Moo")
		assertContains(t, got, "Meaow")
	})

	t.Run("with channels", func(t *testing.T) {
		aChannel := make(chan Profile)

		go func() {
			aChannel <- Profile{33, "Krakow"}
			aChannel <- Profile{44, "Oslo"}
			close(aChannel)
		}()

		var got []string
		want := []string{"Krakow", "Oslo"}

		walk(aChannel, func(input string) {
			got = append(got, input)
		})

		if !slices.Equal(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("with function", func(t *testing.T) {
		someFunc := func() (Profile, Profile) {
			return Profile{33, "Krakow"}, Profile{44, "Oslo"}
		}

		var got []string
		want := []string{"Krakow", "Oslo"}

		walk(someFunc, func(input string) {
			got = append(got, input)
		})

		if !slices.Equal(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}
