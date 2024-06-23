package namegen

type Phrase struct {
	Name      string
	Fragments []PhraseFragment
}

type PhraseFragment struct {
	Words  []string
	Phrase []string
}

type PhraseSet struct {
	Phrases map[string]Phrase

	MaxRecursionLevel int
	MaxPhraseUseCount int
}
