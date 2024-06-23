package namegen

import (
	"log"
	"os"
	"path/filepath"
	"runtime"
)

func getPackagePath() string {
	_, filename, _, _ := runtime.Caller(0)
	return filepath.Dir(filename)
}

func DefaultPhraseSet() *PhraseSet {
	defaultFilename := filepath.Join(getPackagePath(), "default.json")
	data, err := os.ReadFile(defaultFilename)
	if err != nil {
		log.Fatalf("Failed to read file %s", defaultFilename)
	}

	phrases, err := jsonToPhraseSet(data)
	if err != nil {
		log.Fatalf("Error: %s -- This is an error in the default phraseset.", err)
	}

	return phrases
}
