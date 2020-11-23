package main

import "testing"

var buildGetLexemeIDURLTestCases = []struct {
	word     string
	expected string
}{
	{
		word:     "多",
		expected: "https://duolingo-lexicon-prod.duolingo.com/api/1/complete?languageId=zh&query=%E5%A4%9A&uiLanguageId=en",
	},
	{
		word:     "鼓励",
		expected: "https://duolingo-lexicon-prod.duolingo.com/api/1/complete?languageId=zh&query=%E9%BC%93%E5%8A%B1&uiLanguageId=en",
	},
	{
		word:     "猫头鹰",
		expected: "https://duolingo-lexicon-prod.duolingo.com/api/1/complete?languageId=zh&query=%E7%8C%AB%E5%A4%B4%E9%B9%B0&uiLanguageId=en",
	},
	{
		word:     "宫保鸡丁",
		expected: "https://duolingo-lexicon-prod.duolingo.com/api/1/complete?languageId=zh&query=%E5%AE%AB%E4%BF%9D%E9%B8%A1%E4%B8%81&uiLanguageId=en",
	},
}

func TestBuildGetLexemeIDURL(t *testing.T) {
	for _, testCase := range buildGetLexemeIDURLTestCases {
		got := BuildGetLexemeIDURL(testCase.word)
		if got != testCase.expected {
			t.Errorf("BuildGetLexemeIDURL(\"%s\")\nwant: %s, \ngot:  %s", testCase.word, testCase.expected, got)
		}
	}
}
