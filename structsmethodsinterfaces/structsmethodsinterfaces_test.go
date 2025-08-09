package structsmethodsinterfaces

import "testing"

// func TestPerimeter(t *testing.T) {
// 	rectangle := Reactangle{Length: 4.0, Breadth: 5.0}
// 	got := rectangle.Perimeter()
// 	want := 18.0
// 	assertionHelper(t, got,want)
// }

// func TestArea(t *testing.T) {
// 	t.Run("rectangles", func(t *testing.T) {
// 		rectangle := Reactangle{Length: 12.0, Breadth: 6.0}
// 		got := rectangle.Area()
// 		want:= 72.0
// 		assertionHelper(t, got,want)
// 	})

// 	t.Run("circles", func(t *testing.T) {
// 		circle := Circle{10.0}
// 		got := circle.Area()
// 		want := 314.1592653589793
// 		assertionHelper(t, got, want)
// 	})
// }

// func assertionHelper(t testing.TB, got, want float64) {
// 	t.Helper()
// 	if got != want {
// 		t.Errorf("Expected '%.2f' Got '%.2f'", want, got)
// 	}
// }

func TestArea(t *testing.T) {
	testArea := []struct {
		name string
		shape Shape
		want	float64
	} {
		{"Rectangle", Rectangle{12,6}, 72.0},
		{"Circle", Circle{10}, 314.1592653589793},
		{"Triangle", Triangle{12,6}, 36.0},
	}

	for _, tt := range testArea {
		t.Run(tt.name, func(t *testing.T) {
			area := tt.shape.Area()
			if area != tt.want {
				t.Errorf("%#v Got %g want %g",tt.shape, area, tt.want)
			}
		})
	}

	// checkArea := func(t testing.TB, shape Shape, want float64) {
	// 	t.Helper()
	// 	got := shape.Area()
	// 	if got != want {
	// 		t.Errorf("got %g want %g", got, want)
	// 	}
	// }

	// t.Run("rectangles", func(t *testing.T) {
	// 	rectangle := Rectangle{12, 6}
	// 	checkArea(t, rectangle, 72.0)
	// })

	// t.Run("circles", func(t *testing.T) {
	// 	circle := Circle{10}
	// 	checkArea(t, circle, 314.1592653589793)
	// })

}