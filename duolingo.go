package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

type CompleteItem struct {
	LexemeId          string
	MatchedLanguageId string
	MatchedText       string
}

type CompleteResult struct {
	En []CompleteItem
	Zh []CompleteItem
}

func BuildGetLexemeIDURL(word string) string {
	u, _ := url.Parse("https://duolingo-lexicon-prod.duolingo.com/api/1/complete?languageId=zh&uiLanguageId=en")
	q := u.Query()
	q.Set("query", word)
	u.RawQuery = q.Encode()
	return u.String()
}

func GetLexemeID(u, word string) string {
	resp, err := http.Get(u)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	var cr CompleteResult
	err = json.NewDecoder(resp.Body).Decode(&cr)
	if err != nil {
		log.Fatal(err)
	}
	for _, ci := range cr.Zh {
		if ci.MatchedText == word {
			return ci.LexemeId
		}
	}
	return ""
}

func main() {
	word := "多"
	u := BuildGetLexemeIDURL(word)
	lexemeId := GetLexemeID(u, word)
	fmt.Println(lexemeId)
}
