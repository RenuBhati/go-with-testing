package main

// func Perimeter(width, height float64) float64 {
// 	return 0
// }

// func TestPerimeter(t *testing.T) {
// 	got := Perimeter(10.0, 10.0)
// 	want := 40.0
// 	if got != want {
// 		t.Errorf("got %.2f want %.2f", got, want)
// 	}
// }

// func Perimeter(width float64, height float64) float64 {
// 	return 2 * (width + height)
// }

// func Area(width float64, height float64) float64 {
// 	return width * height
// }

// func TestPerimeter(t *testing.T) {
// 	got := Perimeter(10.0, 10.0)
// 	want := 40.0
// 	if got != want {
// 		t.Errorf("got %.2f want %.2f", got, want)
// 	}
// }

// func TestArea(t *testing.T) {
// 	got := Area(10.0, 10.0)
// 	want := 100.0
// 	if got != want {
// 		t.Errorf("got %.2f want %.2f", got, want)
// 	}
// }

// type Rectangle struct {
// 	Width  float64
// 	Height float64
// }

// func Perimeter(rectangle Rectangle) float64 {
// 	return 2 * (rectangle.Height + rectangle.Width)
// }

// func Area(rectangle Rectangle) float64 {
// 	return rectangle.Height * rectangle.Width
// }

// func TestPerimeter(t *testing.T) {
// 	rectangle := Rectangle{10.0, 10.0}
// 	got := Perimeter(rectangle)
// 	want := 40.0

// 	if got != want {
// 		t.Errorf("got %.2f want %.2f", got, want)
// 	}

// }

// func TestArea(t *testing.T) {
// 	rectangle := Rectangle{12.0, 6.0}
// 	got := Area(rectangle)
// 	want := 72.0

// 	if got != want {
// 		t.Errorf("got %.2f want %.2f", got, want)
// 	}

// }

// func TestArea(t *testing.T) {

// 	t.Run("rectangles", func(t *testing.T) {
// 		rectangle := Rectangle{12, 6}
// 		got := Area(rectangle)
// 		want := 72.0

// 		if got != want {
// 			t.Errorf("got %g want %g", got, want)
// 		}
// 	})

// 	t.Run("circles", func(t *testing.T) {
// 		circle := Circle{10}
// 		got := Area(circle)
// 		want := 314.1592653589793

// 		if got != want {
// 			t.Errorf("got %g want %g", got, want)
// 		}
// 	})

// }


type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return 0
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 0
}

