package splitter

import (
	"strings"
)

func findLastMatch(chars []rune, end int) int {
	length := len(chars)

	if end >= length {
		return length
	}

	index := strings.LastIndex(string(chars[:end]), " ")
	if index == -1 {
		index = end
	}

	return index

}

func findChunkSplitIndex(chars []rune, size int) int {
	splitIndex := findLastMatch(chars, size)

	if len(chars) > splitIndex && string(chars[splitIndex]) == " " {
		splitIndex = splitIndex + 1
	}

	return splitIndex

}

//Split splits given string into chunks of given size only at whitespace.
//If a word is longer than given size, the word is broken up into two or more pieces.
func Split(s string, size int) []string {
	var chunks []string
	chars := []rune(s)

	for len(chars) > 0 {
		splitAt := findChunkSplitIndex(chars, size)
		chunk := strings.TrimSpace(string(chars[0:splitAt]))
		if chunk != "" {
			chunks = append(chunks, chunk)
		}
		chars = chars[splitAt:]

	}

	return chunks
}
