package hiragana

type Hiragana struct {
	japanese []string
	english  []string
}

func (r Hiragana) JapaneseById(i int) string {
	return r.japanese[i]
}

func (r Hiragana) EnglishById(i int) string {
	return r.english[i]
}

func (r Hiragana) Japanese() []string {
	return r.japanese
}

func (r Hiragana) English() []string {
	return r.english
}

func NewHiragana() *Hiragana {
	return &Hiragana{
		japanese: []string{"あ", "い", "う", "え", "お", "か", "が", "き", "ぎ", "く", "ぐ", "け", "げ", "こ", "ご", "さ", "ざ", "し", "じ", "す", "ず", "せ", "ぜ", "そ", "ぞ", "た", "だ", "ち", "ぢ", "つ", "づ", "て", "で", "と", "ど", "な", "に", "ぬ", "ね", "の", "は", "ば", "ぱ", "ひ", "び", "ぴ", "ふ", "ぶ", "ぷ", "へ", "べ", "ぺ", "ほ", "ぼ", "ぽ", "ま", "み", "む", "め", "も", "や", "ゆ", "よ", "ら", "り", "る", "れ", "ろ", "わ", "を", "ん"},
		english:  []string{"A", "i", "u", "e", "o", "ka", "ga", "ki", "gi", "ku", "gu", "ke", "ge", "ko", "go", "sa", "za", "shi", "ji", "su", "zu", "se", "ze", "so", "zo", "ta", "da", "chi", "dji", "tsu", "dzu", "te", "de", "to", "do", "na", "ni", "nu", "ne", "no", "wa", "ba", "pa", "hi", "bi", "pi", "fu", "bu", "pu", "e", "be", "pe", "ho", "bo", "po", "ma", "mi", "mu", "me", "mo", "ya", "yu", "yo", "ra", "ri", "ru", "re", "ro", "wa", "o", "n"},
	}
}
