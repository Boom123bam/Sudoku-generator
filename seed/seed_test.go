package seed

import (
	"testing"
)

func TestValid(t *testing.T) {
	seed := GenerateSeed()
	if !seed.Isvalid() {
		t.Errorf("generated seed is not valid: %s", seed)
	}
}

func TestRandom(t *testing.T) {
	seed1 := GenerateSeed()
	seed2 := GenerateSeed()
	if string(seed1) == string(seed2) {
		t.Errorf("generated the same seed twice: %s", seed1)
	}
}
