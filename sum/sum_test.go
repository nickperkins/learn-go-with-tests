package sum

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

}

func TestSumAll(t *testing.T) {

	got := All([]int{1, 2, 3}, []int{2, 4, 6})
	want := []int{6, 12}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v given", got, want)
	}
}

func TestSumAllTails(t *testing.T) {

	checkSums := func(t *testing.T, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v given", got, want)
		}
	}
	t.Run("make the sums of some slices", func(t *testing.T) {
		got := AllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}
		checkSums(t, got, want)
	})
	t.Run("safely sum empty slices", func(t *testing.T) {
		got := AllTails([]int{}, []int{0, 9})
		want := []int{0, 9}
		checkSums(t, got, want)
	})
}
