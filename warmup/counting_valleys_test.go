package warmup

import "testing"

func TestCountingValleys(t *testing.T) {

	result := countingValleys(8, "UDDDUDUU")
	if result != 1 {
		t.Errorf("Expected: 1, got %d", result)
	}

}
