package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	args := os.Args[1:]
	// reading first file
	givenText, _ := os.ReadFile(args[0])
	// This array will push the words into
	// words := strings.Split(strings.Replace(string(givenText), "\n", "", -1), " ")
	words := strings.Split(string(givenText), " ")
	for i, word := range words {
		switch word {
		case "(up)": // upper
			words[i-1] = strings.ToUpper(words[i-1])
			words = append(words[:i], words[i+1:]...)
		case "(down)": // lower
			words[i-1] = strings.ToLower(words[i-1])
			words = append(words[:i], words[i+1:]...)
		case "(cap)": // capitalize
			words[i-1] = strings.Title(words[i-1])
			words = append(words[:i], words[i+1:]...)
		case "(hex)": // hex
			words[i-1] = HextoInt(words[i-1])
			words = append(words[:i], words[i+1:]...)
		case "(bin)": // binary
			words[i-1] = BintoInt(words[i-1])
			words = append(words[:i], words[i+1:]...)
		case "(up,": // upper with number
			b := strings.Trim(string(words[i+1]), words[i+1][1:])
			number, _ := strconv.Atoi(string(b))
			for j := 1; j <= number; j++ {
				words[i-j] = strings.ToUpper(words[i-j])
			}
			words = append(words[:i], words[i+2:]...)
		case "(low,": // lower with number
			b := strings.Trim(string(words[i+1]), words[i+1][1:])
			number, _ := strconv.Atoi(string(b))
			for j := 1; j <= number; j++ {
				words[i-j] = strings.ToLower(words[i-j])
			}
			words = append(words[:i], words[i+2:]...)
		case "(down,": // lower with number
			b := strings.Trim(string(words[i+1]), words[i+1][1:])
			number, _ := strconv.Atoi(string(b))
			for j := 1; j <= number; j++ {
				words[i-j] = strings.ToLower(words[i-j])
			}
			words = append(words[:i], words[i+2:]...)
		case "(cap,": // capitalize with num
			b := strings.Trim(string(words[i+1]), words[i+1][1:])
			number, _ := strconv.Atoi(string(b))
			for j := 1; j <= number; j++ {
				words[i-j] = strings.Title(words[i-j])
			}
			words = append(words[:i], words[i+2:]...)
		}
	}
	ChangeA(words)
	// join slice
	needed := strings.Join(Punctuations(words), " ")
	// write file, automatically updates manipulated file.
	man := os.WriteFile(args[1], []byte(needed), 0o644)
	if man != nil {
		panic(man)
	}
}

// conv hex to int
func HextoInt(hex string) string {
	number, _ := strconv.ParseInt(hex, 16, 64)
	return fmt.Sprint(number)
}

// conv binary to int
func BintoInt(bin string) string {
	number, _ := strconv.ParseInt(bin, 2, 64)
	return fmt.Sprint(number)
}

func ChangeA(s []string) []string {
	vowels := []string{"a", "e", "i", "o", "u", "h", "A", "E", "I", "O", "U", "H"}
	for i, word := range s {
		for _, letter := range vowels {
			if word == "a" && string(s[i+1][0]) == letter {
				s[i] = "an"
			} else if word == "A" && string(s[i+1][0]) == letter {
				s[i] = "An"
			}
		}
	}
	return s
}

func Punctuations(s []string) []string {
	puncs := []string{",", ".", "!", "?", ":", ";"}
	// punc in the middle of a string connecting to word after
	for i, word := range s {
		for _, punc := range puncs {
			if string(word[0]) == punc && string(word[len(word)-1]) != punc {
				s[i-1] += punc
				s[i] = word[1:]
			}
		}
	}
	// punc at end of string
	for i, word := range s {
		for _, punc := range puncs {
			if (string(word[0]) == punc) && (s[len(s)-1] == s[i]) {
				s[i-1] += word
				s = s[:len(s)-1]
			}
		}
	}
	// punc in middle of string
	for i, word := range s {
		for _, punc := range puncs {
			if string(word[0]) == punc && string(word[len(word)-1]) == punc && s[i] != s[len(s)-1] {
				s[i-1] += word
				s = append(s[:i], s[i+1:]...)
			}
		}
	}
	// for apostrophe
	count := 0
	for i, word := range s {
		if word == "'" && count == 0 {
			count += 1
			s[i+1] = word + s[i+1]
			s = append(s[:i], s[i+1:]...)
		}
	}
	//  for second apostrophe
	for i, word := range s {
		if word == "'" {
			// print("here")
			s[i-1] = s[i-1] + word
			// print(s[i-1])
			s = append(s[:i], s[i+1:]...)
		}
	}
	return s
}