package namegen

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func makeStringSlice(interf interface{}) []string {
	processed := interf.([]interface{})
	var slice []string
	for _, word := range processed {
		slice = append(slice, word.(string))
	}
	return slice
}

type ParseError struct {
	Err error
}

func (e *ParseError) Error() string {
	return fmt.Sprintf("%v", e.Err)
}

func jsonParsePhrase(phrase interface{}) ([]PhraseFragment, error) {
	phraseFrags, ok := phrase.([]interface{})
	if !ok {
		return nil, &ParseError{
			Err: fmt.Errorf("JSON error: Failed to parse phrase"),
		}
	}

	var pfSlice []PhraseFragment

	for _, v := range phraseFrags {
		var pf PhraseFragment
		m := v.(map[string]interface{})

		if m["word"] != nil && m["phrase"] != nil {
			return nil, &ParseError{
				Err: fmt.Errorf("You may not specify both a word and phrase at the same time"),
			}
		}

		if m["word"] != nil {
			pf.Words = makeStringSlice(m["word"])
		}
		if m["phrase"] != nil {
			pf.Phrase = makeStringSlice(m["phrase"])
		}

		pfSlice = append(pfSlice, pf)
	}

	return pfSlice, nil
}

/*
Unmarshal JSON text into a PhraseSet.
*/
func Unmarshal(jsonText []byte) (*PhraseSet, error) {
	return jsonToPhraseSet(jsonText)
}

/*
Convert a JSON phrase set file into a usable PhraseSet.
*/
func jsonToPhraseSet(jsonText []byte) (*PhraseSet, error) {
	var data map[string]interface{}
	err := json.Unmarshal(jsonText, &data)
	if err != nil {
		log.Fatalf("JSON parse error")
	}

	var pset PhraseSet
	pset.Phrases = make(map[string]Phrase)

	rawPhrases, ok := data["phrases"].(map[string]interface{})
	if !ok {
		fatal("JSON error: the key \"phrases\" does not exist!")
	}
	for phrName, phrVal := range rawPhrases {
		var phrase Phrase
		phrase.Name = phrName
		phrase.Fragments, err = jsonParsePhrase(phrVal)
		if err != nil {
			return nil, err
		}
		pset.Phrases[phrName] = phrase
	}

	pset.defaults()

	return &pset, nil
}

// Open a .json file containing phrases, and convert it into a PhraseSet.
func OpenPhraseFile(file string) *PhraseSet {
	data, err := os.ReadFile(file)
	if err != nil {
		log.Fatalf("Failed to read file %s", file)
	}

	phrases, err := jsonToPhraseSet(data)

	phrases.defaults()

	return phrases
}
