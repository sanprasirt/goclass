package goclass // import "github.com/sanprasirt/goclass"

import "testing"

func TestSayHello(t *testing.T) {
	subtests := []struct {
		items []string
		result string
	}{
		{
			result: "Hello, World!",
		},
		{
			items: []string{"Bob"},
			result: "Hello, Bob!",
		},
		{
			items: []string{"Bob", "Joe"},
			result: "Hello, Bob Joe!",
		},
	}

	for _, test := range subtests {
		if Say(test.items) != test.result {
			t.Errorf("Say(%v) = %q, want %q", test.items, Say(test.items), test.result)
		}
	}
	// want := "Hello, World!"
	// got := Say([]string{"World"})
	// if got != want {
	// 	t.Errorf("Say() = %q, want %q", got, want)
	// }
}