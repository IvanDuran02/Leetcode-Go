package longestcommonprefix

import "strings"

func longestCommonPrefix(strs []string) string {
	word_len := int(^uint(0) >> 1)
	word := ""

	for _, word2 := range strs {
		if len(word) < word_len {
			word = word2
			word_len = len(word)
		}
	}

	for i := 0; i < len(strs); i++ {
		if strings.HasPrefix(strs[i], word[0]) {
			word = ""
		}
	}

	return word
}
