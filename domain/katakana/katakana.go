package katakana

type Katakana struct {
	japanese string
	english  string
}

func (r Katakana) Japanese() string {
	return r.japanese
}

func (r Katakana) English() string {
	return r.english
}

func NewKatakana(
	japanese string,
	english string,
) *Katakana {
	return &Katakana{
		japanese: japanese,
		english:  english,
	}
}