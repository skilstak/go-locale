package locale

import "testing"

func TestDetect(t *testing.T) {
	loc := Detect()
	if loc == "" {
		t.Fatal("Empty")
	} else {
		t.Logf("detected locale: %q", loc)
	}
}
