package leetcodego

import (
	"strings"
	"unicode"
)

/*
 * @Author: deylr1c
 * @Email: linyugang7295@gmail.com
 * @Description: https://leetcode.cn/problems/valid-palindrome/?envType=study-plan-v2&envId=top-interview-150 125. 验证回文串
 * 输入: s = "A man, a plan, a canal: Panama"
 * 输出：true
 * @Date: 2024-09-25 21:19
 */
func isPalindrome(s string) bool {
	noSpacesStr := strings.ReplaceAll(s, " ", "")
	lowerStr := strings.ToLower(noSpacesStr)
	removeNonLowercase := func(s string) string {
		var builder strings.Builder
		for _, char := range s {
			if unicode.IsLower(char) || unicode.IsDigit(char) {
				builder.WriteRune(char)
			}
		}
		return builder.String()
	}
	lowerStr = removeNonLowercase(lowerStr)
	left, right := 0, len(lowerStr)-1
	for left < right {
		if lowerStr[left] != lowerStr[right] {
			return false
		}
		left++
		right--
	}
	return true
}
func isPalindrome_(s string) bool { // 官方题解
	var sgood string
	for i := 0; i < len(s); i++ {
		if isalnum(s[i]) {
			sgood += string(s[i])
		}
	}

	n := len(sgood)
	sgood = strings.ToLower(sgood)
	for i := 0; i < n/2; i++ {
		if sgood[i] != sgood[n-1-i] {
			return false
		}
	}
	return true
}

func isalnum(ch byte) bool {
	return (ch >= 'A' && ch <= 'Z') || (ch >= 'a' && ch <= 'z') || (ch >= '0' && ch <= '9')
}
