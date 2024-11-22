package database

import "strings"

func StopContain(sentence string) (bool, string) {
	sentence = strings.ReplaceAll(sentence, " ", "")
	sentence = strings.ReplaceAll(sentence, "\n", "")
	sentence = strings.ReplaceAll(sentence, "*", "")
	sentence = strings.ReplaceAll(sentence, "\\", "")
	sentence = strings.ReplaceAll(sentence, "/", "")
	sentence = strings.ReplaceAll(sentence, ".", "")
	sentence = strings.ReplaceAll(sentence, ",", "")
	sentence = strings.ReplaceAll(sentence, "。", "")
	sentence = strings.ReplaceAll(sentence, "，", "")
	var stops []StopTable
	DB.Model(&StopTable{}).Find(&stops)
	for _, i := range stops {
		if strings.Contains(sentence, i.Text) {
			return true, i.Text
		}
	}
	return false, ""
}

func StopAdd(text string) bool {
	b, _ := StopContain(text)
	if b {
		return false
	}
	var i StopTable
	i.Text = text
	DB.Create(&i)
	return true
}
