package locale

import "testing"

func TestDetect(t *testing.T) {
	if lc, err := Detect(); err != nil {
		t.Fatal(err)
	} else {
		t.Logf("detected locale: %q", lc)
	}
}
