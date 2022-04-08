package kanji

type Kanji struct {
	japanese string
	english  string
}

func (r Kanji) Japanese() string {
	return r.japanese
}

func (r Kanji) English() string {
	return r.english
}

func NewKanji(
	japanese string,
	english string,
) *Kanji {
	return &Kanji{
		japanese: japanese,
		english:  english,
	}
}
