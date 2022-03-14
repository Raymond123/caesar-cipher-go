package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	// List of messages
	messages := []string{"Csi2520", "Csi2120", "3 Paradigms",
		"Go is 1st", "Prolog is 2nd", "Scheme is 3rd",
		"uottawa.ca", "csi/elg/ceg/seg", "800 King Edward"}

	done := make(chan int)
	item := make(chan string)

	// PART 1 // PRINT IS IN FUNCTION
	caesarCipher("I love CS!", 5)

	// PART 2, COMMENTED OUT CAUSE USED THE SAME CHANNELS AS PART 3
	//go caesarCipherList(messages[:], 2, item, done)
	fmt.Println()

	// PART 3
	go caesarCipherList(messages[:3], 2, item, done)
	go caesarCipherList(messages[3:6], 2, item, done)
	go caesarCipherList(messages[6:], 2, item, done)

	// CHANNEL SYNCING AND RESULT PRINTING
	sync := 0
	for {
		msg, _ := <-item
		fmt.Println(msg)

		if sync >= len(messages)-1 {
			break
		}
		sync = sync + <-done
	}

	// CLOSE CHANNELS
	close(done)
	close(item)
}

// CIPHER FUNCTION FOR PART 1
func caesarCipher(msg string, shift int) {
	char := toRunes(msg)
	var encode []rune

	for _, i := range char {
		if !(unicode.IsLetter(i)) {
			continue
		}
		rn := int(i) + shift
		if !(unicode.IsLetter(rune(rn))) {
			rn -= 26
		}
		encode = append(encode, rune(rn))
	}

	fmt.Println(strings.ToUpper(string(encode)))
}

// CIPHER FUNCTION FOR PARTS 2 & 3
func caesarCipherList(msg []string, shift int, item chan string, done chan int) {
	for _, m := range msg {
		char := toRunes(m)
		var encode []rune

		for _, i := range char {
			if !(unicode.IsLetter(i)) {
				continue
			}
			rn := int(i) + shift
			if !(unicode.IsLetter(rune(rn))) {
				rn -= 26
			}
			encode = append(encode, rune(rn))
		}
		item <- strings.ToUpper(string(encode))
		done <- 1
	}
}

func toRunes(s string) []rune {
	var r []rune // slice of unicode chars
	// from string to slice of unicode
	for _, c := range s {
		r = append(r, c)
	}

	return r
}
