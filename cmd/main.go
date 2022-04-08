package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"japaneseapp/domain/hiragana"
	"japaneseapp/domain/katakana"
	"math/rand"
	"strings"
	"time"
)

type Type int

const (
	HIRAGANA Type = iota
	KATAKANA
	KANJI
)

type Answer struct {
	Id     int   `json:"id"`
	Type   Type  `json:"type"`
	Answer bool  `json:"answer"`
	Time   int64 `json:"time"`
}

func main() {
	fmt.Println(" ⮀  choose a kana system or kanji words")
	fmt.Println(" 1.  hiragana")
	fmt.Println(" 2.  katakana")
	fmt.Println(" 3.  kanji")
	fmt.Print(" ⮀ ")
	var userResponse int
	fmt.Scanln(&userResponse)

	japaneseList := []string{}
	englishList := []string{}
	var listType Type

	switch userResponse {
	case 1:
		hiraganaLetters := hiragana.NewHiragana()
		japaneseList = hiraganaLetters.Japanese()
		englishList = hiraganaLetters.English()
		listType = HIRAGANA
	case 2:
		katakanaLetters := katakana.NewKatakana()
		japaneseList = katakanaLetters.Japanese()
		englishList = katakanaLetters.English()
		listType = KATAKANA
	case 3:
		fmt.Println("in development")
		return
	default:
		fmt.Println("Enter a number between 1 to 3")
		return
	}
	fmt.Println("-------------------------------------------")

	fmt.Println(" ⮀  choose mod (default easy) ")
	fmt.Println(" 1.  easy (3 hearts)")
	fmt.Println(" 2.  hard (1 heart)")
	fmt.Print(" ⮀ ")
	var userResponsePickMod int
	fmt.Scanln(&userResponsePickMod)

	heartsCounter := 0
	allHearts := 0

	switch userResponsePickMod {
	case 1:
		heartsCounter = 3
	case 2:
		heartsCounter = 1
	default:
		fmt.Println(" you picked default which is easy")
		heartsCounter = 3
	}
	allHearts = heartsCounter
	fmt.Println("-------------------------------------------")

	allAnswersCorrect := true
	questionsCounter := 0
	correctAnswersCounter := 0

	answers := []Answer{}

	for allAnswersCorrect == true {
		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)
		randonNumber := r1.Intn(len(japaneseList))

		fmt.Print("   " + japaneseList[randonNumber])
		questionsCounter = questionsCounter + 1

		fmt.Print(" ⮕  ")
		var userResponse string
		fmt.Scanln(&userResponse)

		if strings.ToLower(userResponse) != englishList[randonNumber] {

			fmt.Println(" ✘ " + englishList[randonNumber])
			heartsCounter = heartsCounter - 1
			if heartsCounter < 0 {
				allAnswersCorrect = false
			} else {
				fmt.Println(" 💔 ", heartsCounter, "/", allHearts)
			}

			answers = append(answers, Answer{
				Id:     randonNumber,
				Answer: false,
				Type:   listType,
				Time:   time.Now().Unix(),
			})
		} else {
			fmt.Println(" ✔  correct ")
			correctAnswersCounter = correctAnswersCounter + 1

			answers = append(answers, Answer{
				Id:     randonNumber,
				Answer: true,
				Type:   listType,
				Time:   time.Now().Unix(),
			})
		}
	}

	file, _ := json.MarshalIndent(answers, "", " ")
	jsonFileName := "answers/" + fmt.Sprint(time.Now().Unix()) + "-answer"

	_ = ioutil.WriteFile(jsonFileName+".json", file, 0644)

	fmt.Println(" answered correctly to ", correctAnswersCounter, " from ", questionsCounter)
}
