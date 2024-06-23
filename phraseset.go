package namegen

func (ps *PhraseSet) defaults() {
	ps.MaxPhraseUseCount = 5
	ps.MaxRecursionLevel = 200
}
