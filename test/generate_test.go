package namegen_test

import (
	"strings"
	"testing"

	ng "github.com/a-random-lemurian/go-namegen"
)

func TestRecursion(t *testing.T) {
	phrase := ng.OpenPhraseFile(getTestPhraseSet())

	defer func() {
		if r := recover(); r != nil {
			expectedPrefix := "Excessive recursion!"
			panicMsg := r.(string)
			if !strings.HasPrefix(panicMsg, expectedPrefix) {
				t.Errorf("Unexpected panic: %s", panicMsg)
			} else {
				t.Logf("Panic occurs as expected.\n")
			}
		}
	}()

	phrase.GenerateString("bad recursion")
}
