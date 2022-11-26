package main

import (
	"os"
	"strconv"
	"strings"
)

// Read files
func main() {
	givenText, _ := os.ReadFile(os.Args[1])
	finalText := NewInput(strings.Split(string(givenText), " "))
	output(finalText)
}

// create a new string
var TempString = []string{}

// make a required changes on input text and return the string of input text
func NewInput(articleArray []string) string {
	TempString = append(TempString, articleArray[0])
	for i := 1; i < len(articleArray); i++ {
		getStringArray := TempString[len(TempString)-1]
		pos := len(TempString) - 1
		switch strings.ToLower(articleArray[i]) {
		case "(hex)":
			TempString[pos] = numberConvert(getStringArray, 16)
		case "(bin)":
			TempString[pos] = numberConvert(getStringArray, 2)
		case "(up)":
			TempString[pos] = strings.ToUpper(getStringArray)
		case "(low)":
			TempString[pos] = strings.ToLower(getStringArray)
		case "(cap)":
			TempString[pos] = strings.Title(getStringArray)
		case "(cap,":
			for j := 1; j <= setNumber(articleArray[i+1]); j++ {
				TempString[len(TempString)-j] = strings.Title(TempString[len(TempString)-j])
			}
			i++
		case "(low,":
			for j := 1; j <= setNumber(articleArray[i+1]); j++ {
				TempString[len(TempString)-j] = strings.ToLower(TempString[len(TempString)-j])
			}
			i++
		case "(up,":
			for j := 1; j <= setNumber(articleArray[i+1]); j++ {
				TempString[len(TempString)-j] = strings.ToUpper(TempString[len(TempString)-j])
			}
			i++
		case "a":
			TempString = append(TempString, articleControl(articleArray[i], string(articleArray[i+1][0])))
		default:
			TempString = append(TempString, articleArray[i])
		}
	}

	return puncCorrection(strings.Join(TempString, " "))
}

func articleControl(article string, firstChar string) string {
	// check if the word is an article
	if strings.ContainsAny(firstChar, "aeiouhAEIOUH") {
		return article + "n"
	}
	return article
}

func setNumber(w string) int {
	// find number
	intVar, _ := strconv.Atoi(w[:len(w)-1])
	return intVar
}

func numberConvert(inputNum string, i int) string {
	// convert to int
	num, _ := strconv.ParseInt(inputNum, i, 64)
	return strconv.Itoa(int(num))
}

func puncCorrection(finalText string) string {
	// correct punctuation
	for range finalText {
		finalText = strings.ReplaceAll(finalText, " ,", ", ")
		finalText = strings.ReplaceAll(finalText, " .", ". ")
		finalText = strings.ReplaceAll(finalText, " !", "! ")
		finalText = strings.ReplaceAll(finalText, " ?", "? ")
		finalText = strings.ReplaceAll(finalText, " :", ": ")
		finalText = strings.ReplaceAll(finalText, " ;", "; ")
		finalText = strings.ReplaceAll(finalText, "  ", " ")
	}
	// correct capitalization
	finalText = strings.ReplaceAll(finalText, " '", "'")
	finalText = strings.ReplaceAll(finalText, ":' ", ": '")
	return finalText
}

func output(fileContent string) {
	// write to file
	outputFile := os.Args[2]
	output := []byte(fileContent)
	os.WriteFile(outputFile, output, 0o664)
}
