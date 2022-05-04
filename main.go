package main

import (
	"fmt"
	"strings"
)

func main() {
	var number int
	s := "man i need a taxi up to ubud"
	words := strings.Fields(s)
	number = searchWord(words, 0, 0, 0, 0)
	fmt.Println(number)
	fmt.Println(words[number-1])
}

func searchWord(words []string, number, sum, length, count int) int {
	var sumTest int
	count++
	if len(words) > 0 {
		wordCheck := []byte(words[0])
		for i := 0; i < len(wordCheck); i++ {
			sumTest += int(wordCheck[i])
		}
		sumTest -= 96 * len(wordCheck)
		if sum < sumTest {
			length = len(wordCheck)
			sum = sumTest
			number = count
		}

		number = searchWord(words[1:], number, sum, length, count)
	}
	return number
}

//Given a string of words, you need to find the highest scoring word.

//Each letter of a word scores points according to its position in the alphabet: a = 1, b = 2, c = 3 etc.

//You need to return the highest scoring word as a string.

//If two words score the same, return the word that appears earliest in the original string.

//All letters will be lowercase and all inputs will be valid.

//var _ = Describe("Example tests",func() {
//    arr := [...][2]string{
//        {"man i need a taxi up to ubud","taxi"},
//        {"what time are we climbing up the volcano","volcano"},
//        {"take me to semynak","semynak"},
//        {"aa b", "aa"},
//        {"b aa", "b"},
//        {"bb d", "bb"},
//        {"d bb", "d"},
//        {"aaa b", "aaa"},
//    }
//
