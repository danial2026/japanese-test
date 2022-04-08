package hiragana

type Hiragana struct {
	japanese string
	english  string
}

func (r Hiragana) Japanese() string {
	return r.japanese
}

func (r Hiragana) English() string {
	return r.english
}

func NewHiragana(
	japanese string,
	english string,
) *Hiragana {
	return &Hiragana{
		japanese: japanese,
		english:  english,
	}
}
