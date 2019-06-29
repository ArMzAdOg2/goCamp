package rotate

import (
	"reflect"
	"testing"
)

func TestRotateMe(t *testing.T) {
	t.Run("Rotate when input (1,2,3,4,5) expect (5,1,2,3,4)", func(t *testing.T) {
		n := []int{1, 2, 3, 4, 5}

		expected := []int{5, 1, 2, 3, 4}
		r := Rotate(n)
		if !reflect.DeepEqual(r, expected) {
			t.Errorf("expect % #v but got % #v", expected, r)
		}
	})

	t.Run("Rotate when input (5,1,2,3,4) expect (4,5,1,2,3)", func(t *testing.T) {
		n := []int{5, 1, 2, 3, 4}

		expected := []int{4, 5, 1, 2, 3}
		r := Rotate(n)
		if !reflect.DeepEqual(r, expected) {
			t.Errorf("expect % #v but got % #v", expected, r)
		}
	})

	t.Run("Rotate when input (4,5,1,2,3) expect (3,4,5,1,2)", func(t *testing.T) {
		n := []int{4, 5, 1, 2, 3}

		expected := []int{3, 4, 5, 1, 2}
		r := Rotate(n)
		if !reflect.DeepEqual(r, expected) {
			t.Errorf("expect % #v but got % #v", expected, r)
		}
	})

	t.Run("Rotate when input () expect ()", func(t *testing.T) {
		n := []int{}

		expected := []int{}
		r := Rotate(n)
		if !reflect.DeepEqual(r, expected) {
			t.Errorf("expect % #v but got % #v", expected, r)
		}
	})

	t.Run("Rotate when input (1) expect (1)", func(t *testing.T) {
		n := []int{1}

		expected := []int{1}
		r := Rotate(n)
		if !reflect.DeepEqual(r, expected) {
			t.Errorf("expect % #v but got % #v", expected, r)
		}
	})

}

func BenchmarkRotateMe(b *testing.B) {
	b.Run("Rotate when input (1,2,3,4,5) expect (5,1,2,3,4)", func(b *testing.B) {
		n := []int{1, 2, 3, 4, 5}
		i := 0
		for i < 1000 {
			i++
			Rotate(n)
		}

	})

}
