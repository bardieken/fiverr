package main

import (
	"strconv"
	"strings"
	"os"
)

func main() {
	// read file content, write it into string and new output file
	input, _ := os.ReadFile(os.Args[1])
	finalText := ChangeInput(strings.Split(string(input), " "))
	writeFile(finalText)
}

var (
	// temporary file content holder
	Newtext   = []string{}
)

func ChangeInput(textArray []string) string{
	//make a required changes on input text and return the string of input text
	Newtext = append(Newtext, textArray[0])
	for i := 1; i < len(textArray); i++ {
		stringFromArray := Newtext[len(Newtext)-1]
		pos := len(Newtext) - 1
			switch strings.ToLower(textArray[i]) {
			case "(hex)":
				Newtext[pos] = convertNum(stringFromArray, 16)
			case "(bin)":
				Newtext[pos] = convertNum(stringFromArray, 2)
			case "(up)":
				Newtext[pos] = strings.ToUpper(stringFromArray)
			case "(low)":
				Newtext[pos] = strings.ToLower(stringFromArray)
			case "(cap)":
				Newtext[pos] = strings.Title(stringFromArray)
			case "(cap,":
				for j := 1; j <= findNumber(textArray[i+1]); j++ {
					Newtext[len(Newtext)-j] = strings.Title(Newtext[len(Newtext)-j])
				}
				i++
			case "(low,":
				for j := 1; j <= findNumber(textArray[i+1]); j++ {
					Newtext[len(Newtext)-j] = strings.ToLower(Newtext[len(Newtext)-j])
				}
				i++
			case "(up,":
				for j := 1; j <= findNumber(textArray[i+1]); j++ {
					Newtext[len(Newtext)-j] = strings.ToUpper(Newtext[len(Newtext)-j])
				}
				i++
			case "a"  :
				Newtext = append(Newtext, checkArticel(textArray[i], string(textArray[i+1][0])))
			default:
				Newtext = append(Newtext, textArray[i])
			}
	}
	//join an array of strings into a string, separate with a space. 
	return correctPunctuation(strings.Join(Newtext, " "))
}

func checkArticel(article string, firstChar string) string {
	//check if the next word starts with a vowel or "h" and return the correct article
	if strings.ContainsAny(firstChar, "aeiouhAEIOUH") {
		return article + "n"
	}
	return article
}

func findNumber(w string) int {
	//convert the given number in string into an integer
	intVar, _ := strconv.Atoi(w[:len(w)-1])
	return intVar
}

func convertNum(inputNum string, i int) string {
	//cconvert the given hexadecimal or binary number into decimal and return the string of this number
	num, _ := strconv.ParseInt(inputNum, i, 64)
	return strconv.Itoa(int(num))
}

func correctPunctuation(finalText string) string {
	//replace all wrong spacing and return correct final text
	for range finalText {
		finalText = strings.ReplaceAll(finalText, " ,", ", ")
		finalText = strings.ReplaceAll(finalText, " .", ". ")
		finalText = strings.ReplaceAll(finalText, " !", "! ")
		finalText = strings.ReplaceAll(finalText, " ?", "? ")
		finalText = strings.ReplaceAll(finalText, " :", ": ")
		finalText = strings.ReplaceAll(finalText, " ;", "; ")
		finalText = strings.ReplaceAll(finalText, "  ", " ")
	}
	//handel exceptions with upper comma
	finalText = strings.ReplaceAll(finalText, " '", "'")
	finalText = strings.ReplaceAll(finalText, ":' ", ": '")
	return finalText
}

func writeFile(fileContent string) {
	//get an output filename, convert it to bytes and write a new file
	outputFile := os.Args[2]
	output := []byte(fileContent)
	os.WriteFile(outputFile, output, 0o664)
}