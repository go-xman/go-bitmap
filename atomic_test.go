package bitmap

import "testing"

func TestSetAtomic(t *testing.T) {
	slice := make([]byte, 4)
	SetAtomic(slice, 31, true)
	if !Get(slice, 31) {
		t.Fatal("should return true")
	}
}
