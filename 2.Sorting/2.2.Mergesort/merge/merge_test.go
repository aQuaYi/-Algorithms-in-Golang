package merge

import "testing"

func isSorted(a []int) bool {
	for i := 1; i < len(a); i++ {
		if a[i] < a[i-1] {
			return false
		}
	}
	return true
}

func Test_Sort(t *testing.T) {
	s := make([]int, 10)
	for i := 0; i < len(s); i++ {
		s[i] = len(s) - i - 1
	}
	t.Log(s)
	s = Sort(s)
	t.Log(s)
	if !isSorted(s) {
		t.Error("没能排序好")
	}
}

func Test_Sort_1000000(t *testing.T) {
	s := make([]int, 1000000)
	for i := 0; i < len(s); i++ {
		s[i] = len(s) - i - 1
	}
	s = Sort(s)
	if !isSorted(s) {
		t.Error("没能排序好")
	}
}

func Test_BUSort(t *testing.T) {
	s := make([]int, 10)
	for i := 0; i < len(s); i++ {
		s[i] = len(s) - i - 1
	}
	t.Log(s)
	s = BUSort(s)
	t.Log(s)
	if !isSorted(s) {
		t.Error("没能排序好")
	}
}

func Test_BUSort_1000000(t *testing.T) {
	s := make([]int, 1000000)
	for i := 0; i < len(s); i++ {
		s[i] = len(s) - i - 1
	}
	s = BUSort(s)
	if !isSorted(s) {
		t.Error("没能排序好")
	}
}