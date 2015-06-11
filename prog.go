package prog

import "fmt"

// Widget is a thing we make.
type Widget struct {
	Name  string
	Color string
}

// Describe returns a description about the widget's name and color.
func (w *Widget) Describe() string {
	if w.Color == "" {
		return fmt.Sprintf("My name is %q and I have no color", w.Name)
	}
	return fmt.Sprintf("My name is %q and I am %q", w.Name, w.Color)
}
