package main

import (
	"fmt"
	"math"
)

func longestSubstring(str string) int {
	if len(str) <= 1 {
		return len(str)
	}

	leftPointer := 0
	longestSubstringCount := 0

	seenCharactersMap := map[string]int{}

	for rightPointer := 0; rightPointer < len(str); rightPointer++ {

		currentCharacter := string(str[rightPointer])
		seenCharacterIndex := seenCharactersMap[currentCharacter]

		if seenCharacterIndex >= leftPointer {
			leftPointer = seenCharacterIndex + 1
		}

		seenCharactersMap[currentCharacter] = rightPointer
		longestSubstringCount = int(math.Max(float64(longestSubstringCount), float64(rightPointer-leftPointer+1)))
	}
	return longestSubstringCount

}

func main() {
	longestSubstringCount := longestSubstring("abcdabc")

	fmt.Printf("The longest substring count is %d", longestSubstringCount)
}
