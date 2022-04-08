package katakana

type Katakana struct {
	japanese []string
	english  []string
}

func (r Katakana) JapaneseById(i int) string {
	return r.japanese[i]
}

func (r Katakana) EnglishById(i int) string {
	return r.english[i]
}

func (r Katakana) Japanese() []string {
	return r.japanese
}

func (r Katakana) English() []string {
	return r.english
}

func NewKatakana() *Katakana {
	return &Katakana{
		japanese: []string{"ア", "イ", "ウ", "エ", "オ", "カ", "キ", "ク", "ケ", "コ", "サ", "シ", "ス", "セ", "ソ", "タ", "チ", "ツ", "テ", "ト", "ナ", "ニ", "ヌ", "ネ", "ノ", "ハ", "ヒ", "フ", "ヘ", "ホ", "マ", "ミ", "ム", "メ", "モ", "ヤ", "ユ", "ヨ", "ラ", "リ", "ル", "レ", "ロ", "ワ", "ヰ", "ヱ", "ヲ", "ン"},
		english:  []string{"a", "i", "u", "e", "o", "ka", "ki", "ku", "ke", "ko", "sa", "shi", "su", "se", "so", "ta", "chi", "tsu", "te", "to", "na", "ni", "nu", "ne", "no", "ha", "hi", "fu", "e", "ho", "ma", "mi", "mu", "me", "mo", "ya", "yu", "Yo", "ra", "Ri", "ru", "re", "Ro", "wa", "wi", "we", "o", "n"},
	}
}
