package namegen_test

import (
	"fmt"
	"testing"

	ng "github.com/a-random-lemurian/go-namegen"
)

type testCase struct {
	phrase   string
	expected int
}

var testCases = []testCase{
	{phrase: "3 possibilities", expected: 3},
	{phrase: "3 possibilities with embedded phrase", expected: 3},
}

func TestEnumerate(t *testing.T) {
	phrase := ng.OpenPhraseFile(getTestPhraseSet())

	for _, tcase := range testCases {
		permuts := phrase.Enumerate(tcase.phrase)
		phraseTestResultString := fmt.Sprintf("phrase '%v': got %v - expected %v",
			tcase.phrase, permuts, tcase.expected)
		if permuts != int64(tcase.expected) {
			t.Fatalf("bad  -- %v", phraseTestResultString)
		} else {
			t.Logf("good -- %v", phraseTestResultString)
		}
	}
}

func BenchmarkEnumerate(b *testing.B) {
	phrase := ng.DefaultPhraseSet()
	for i := 0; i < b.N; i++ {
		phrase.Enumerate("shipname")
	}
}

func BenchmarkGenerate(b *testing.B) {
	phrase := ng.DefaultPhraseSet()
	for i := 0; i < b.N; i++ {
		phrase.GenerateString("shipname")
	}
}
