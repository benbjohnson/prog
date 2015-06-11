package prog_test

import (
	"fmt"
	"testing"

	"github.com/benbjohnson/prog"
)

// Ensure that a widget can output a correct description of itself.
func TestWidget_Describe(t *testing.T) {
	// Create a named, colorful widget.
	w := &prog.Widget{Name: "bob", Color: "purple"}

	// Verify that it can describe itself.
	desc := w.Describe()
	if desc != `My name is "bob" and I am "purple"` {
		t.Fatalf("unexpected description: %s", desc)
	}
}

// Ensure that a widget with no color can describe itself.
func TestWidget_Describe_Colorless(t *testing.T) {
	// Create a named widget without a color.
	w := &prog.Widget{Name: "susy"}

	// Verify that it can describe itself.
	desc := w.Describe()
	if desc != `My name is "susy" and I have no color` {
		t.Fatalf("unexpected description: %s", desc)
	}
}

func TestWidget_Describe_TableTest(t *testing.T) {
	var tests = []struct {
		Widget   *prog.Widget
		Expected string
	}{
		{
			Widget:   &prog.Widget{Name: "bob"},
			Expected: `-`,
		},

		{
			Widget:   &prog.Widget{Color: "purple"},
			Expected: `-`,
		},
	}

	for i, tt := range tests {
		got := tt.Widget.Describe()
		if got != tt.Expected {
			t.Errorf("%d. unexpected description: %s", i, got)
		}
	}
}

func testWidget_Describe(t *testing.T, w *prog.Widget, exp string) {
	got := w.Describe()
	if got != exp {
		t.Fatalf("unexpected description: %s", got)
	}
}

func TestWidget_Describe_WithNameOnly(t *testing.T) {
	testWidget_Describe(t, &prog.Widget{Name: "bob"}, `-`)
}

func BenchmarkWidget_Describe(b *testing.B) {
	w := &prog.Widget{Name: "susy"}

	for i := 0; i < b.N; i++ {
		if w.Describe() == "" {
			b.Fatal("description failed")
		}
	}
}

func ExampleWidget_Describe() {
	// Create a widget.
	w := &prog.Widget{
		Name:  "Susan",
		Color: "green",
	}

	// Print out a description of your widget.
	fmt.Println("Your widget:")
	fmt.Println(w.Describe())

	// Output:
	// Your widget:
	// My name is "Susan" and I am "green"
}
