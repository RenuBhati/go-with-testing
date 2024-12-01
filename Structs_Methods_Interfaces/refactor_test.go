package main

// // 1

// func Perimeter(width float64, height float64) float64 {
// 	return 0
// }

// func TestPerimeter(t *testing.T) {
// 	got := Perimeter(10.0, 10.0)
// 	want := 40.0

// 	if got != want {
// 		t.Errorf("got %.2f want %.2f", got, want)
// 	}
// }

// // 2
// func Perimeter(width float64, height float64) float64 {
// 	return 2 * (width + height)
// }

// func TestPerimeter(t *testing.T) {
// 	got := Perimeter(10.0, 10.0)
// 	want := 40.0

// 	if got != want {
// 		t.Errorf("got %.2f want %.2f", got, want)
// 	}
// }

// // 3 - Refactor
// type Rectangle struct {
// 	Width  float64
// 	Height float64
// }

// func Perimeter(r Rectangle) float64 {
// 	return 2 * (r.Width + r.Height)
// }

// func TestPerimeter(t *testing.T) {
// 	r := Rectangle{10.0, 10.0}
// 	got := Perimeter(r)
// 	want := 40.0

// 	if got != want {
// 		t.Errorf("got %.2f want %.2f", got, want)
// 	}
// }

// 4 - New type

// type Circle struct {
// 	Radius float64
// }

// func (c Circle) Perimeter() float64 {
// 	return 3.14 * c.Radius * c.Radius
// }

// func TestPerimeter(t *testing.T) {
// 	t.Run("I am testing fir Circle ", func(t *testing.T) {
// 		c := Circle{10}
// 		got := c.Perimeter()
// 		want := 314.0

// 		if got != want {
// 			t.Errorf("got %g want %g", got, want)
// 		}
// 	})
// }

// //5 Methods

// type Rectangle struct {
// 	Width  float64
// 	Height float64
// }

// func (r Rectangle) Area() float64 {
// 	return r.Width * r.Height
// }

// type Circle struct {
// 	Radius float64
// }

// func (c Circle) Area() float64 {
// 	return math.Pi * c.Radius * c.Radius
// }

// // 6 RE-FACTOR - Interface

// type Shape interface {
// 	Area() float64
// }

// type Triangle struct {
// 	Base   float64
// 	Height float64
// }

// func (t Triangle) Area() float64 {
// 	return (t.Base * t.Height) * 0.5
// }
