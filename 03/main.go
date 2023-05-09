package main

import (
	"log"
)

func main() {

	//i := lengthOfLongestSubstring("abcabcbb")
	//if 3 != i {
	//	log.Fatalf("%d should be 3", i)
	//}
	//i = lengthOfLongestSubstring("bbbbb")
	//if 1 != i {
	//	log.Fatalf("%d should be 1", i)
	//}
	//
	//i = lengthOfLongestSubstring("pwwkew")
	//if 3 != i {
	//	log.Fatalf("%d should be 3", i)
	//}

	i := lengthOfLongestSubstring("dvdf")
	if 3 != i {
		log.Fatalf("%d should be 3", i)
	}
}

// Create a dict that holds every letter as key and the index of the letter as value
// When the letter repeats itself
// we start a new dict from index +1
// if newDict.size() > prevDict.size() forget about prevDict
// At the end, return max(newDict.size(), prevDict.size())
func lengthOfLongestSubstring(s string) int {
	bestDict := make(map[string]int, 256)
	currentDict := make(map[string]int, 256)
	for i := 0; i < len(s); i++ {
		key := string(s[i])
		prevPosition, exists := currentDict[key]
		if exists {
			if len(currentDict) > len(bestDict) {
				bestDict = currentDict
			}
			currentDict = make(map[string]int, 256)
			i = prevPosition
			continue
		}
		currentDict[key] = i
	}

	return max(len(bestDict), len(currentDict))
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
