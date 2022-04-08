package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type KanaLetter struct {
	Japanese string `json:"japanese"`
	English  string `json:"english"`
	Repeat   int64  `json:"repeat"`
	Score    int64  `json:"score"`
}

func main() {
	fmt.Println(" ⮀  choose a kana system or kanji words")
	fmt.Println(" 1.  hiragana")
	fmt.Println(" 2.  katakana")
	fmt.Println(" 3.  kanji")
	fmt.Print(" ⮀ ")
	var userResponse int
	fmt.Scanln(&userResponse)

	kanaList := []KanaLetter{}
	var listType string

	switch userResponse {
	case 1:
		listType = "hiragana"
	case 2:
		listType = "katakana"
	case 3:
		fmt.Println("in development")
		/*
			 TODO : read the json file
			 first pick what u wanna do
			 1. review kanji words
			 2. add kanji
			 3. print all kanji words

			 for add kanji : https://dev.to/evilcel3ri/append-data-to-json-in-go-5gbj
			 for review kanji words just read the json file

			 use this to print all kanji words :

				w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)
				for _, j := range kanaList {
					fmt.Fprintln(w, j.English+"\t"+j.Japanese)
				}

				w.Flush()
		*/
		return
	default:
		fmt.Println("Enter a number between 1 to 3")
		return
	}

	kanaList = getListFromJson("./answers/", listType+".json")

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

	// answers := []Answer{}

	for allAnswersCorrect == true {
		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)
		randonNumber := r1.Intn(len(kanaList))

		fmt.Print("   " + kanaList[randonNumber].Japanese)
		questionsCounter = questionsCounter + 1
		kanaList[randonNumber].Repeat = kanaList[randonNumber].Repeat + 1

		fmt.Print(" ⮕  ")
		var userResponse string
		fmt.Scanln(&userResponse)

		if strings.ToLower(userResponse) != kanaList[randonNumber].English {

			fmt.Println(" ✘ " + kanaList[randonNumber].English)
			if heartsCounter <= 0 {
				allAnswersCorrect = false
			} else {
				fmt.Println(" 💔 ", heartsCounter, "/", allHearts)
			}
			heartsCounter = heartsCounter - 1

			kanaList[randonNumber].Score = kanaList[randonNumber].Score - 1
		} else {
			fmt.Println(" ✔  correct ")
			correctAnswersCounter = correctAnswersCounter + 1

			kanaList[randonNumber].Score = kanaList[randonNumber].Score + 1
		}
	}

	file, _ := json.MarshalIndent(kanaList, "", " ")
	jsonFileName := "answers/" + listType
	_ = ioutil.WriteFile(jsonFileName+".json", file, 0644)

	fmt.Println(" answered correctly to ", correctAnswersCounter, " from ", questionsCounter)
}

func getListFromJson(path string, fileName string) []KanaLetter {
	filename := path + fileName
	doesErr := checkFile(filename)
	if doesErr != true {
		filename = fileName
	}

	file, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	}

	data := []KanaLetter{}
	json.Unmarshal(file, &data)

	return data
}

func checkFile(filename string) bool {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return true
}
