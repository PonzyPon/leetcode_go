package no0169

import "testing"

func TestMajorityElement(t *testing.T) {
	want := 0
	got := majorityElement([]int{1, 2})
	if got != want {
		t.Errorf("majorityElement() = %d; want %d", got, want)
	}
}
