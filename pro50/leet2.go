package leetcodego

/*
 * @Author: deylr1c
 * @Email: linyugang7295@gmail.com
 * @Description: https://leetcode.cn/problems/valid-parentheses/?envType=study-plan-v2&envId=2024-spring-sprint-100 20.有效括号
 * 输入：s = "()[]{}"
 * 输出：true
 * @Date: 2024-09-23 17:03
 */
type Stack[T any] struct { // 手搓stack
	len     int32
	index   int32
	element []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		len:     0,
		index:   -1,
		element: make([]T, 0),
	}
}
func (stack *Stack[T]) IsEmpty() bool {
	return stack.index < 0
}
func (stack *Stack[T]) Push(data T) {
	stack.element = append(stack.element, data)
	stack.index++
	stack.len++
}
func (stack *Stack[T]) Pop() (T, bool) {
	if stack.IsEmpty() {
		var zero T
		return zero, false
	}
	data := stack.element[stack.index]
	stack.element = stack.element[:stack.index]
	stack.index--
	stack.len--
	return data, true
}
func (stack *Stack[T]) Check() (T, bool) {
	// 如果栈为空，返回零值和 false
	if stack.IsEmpty() {
		var zero T
		return zero, false
	}
	return stack.element[stack.index], true
}
func IsValid(s string) bool {
	stack := NewStack[rune]() // 使用泛型的栈来存储括号字符
	mapping := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}
	for _, char := range s {
		if closingBracket, ok := mapping[char]; ok {
			top, have := stack.Pop()
			if !have && top != closingBracket {
				return false
			}
		} else {
			stack.Push(char)
		}
	}
	return stack.IsEmpty()
}
func isValid(s string) bool { // 官方题解
	n := len(s)
	if n%2 == 1 {
		return false
	}
	pairs := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	stack := []byte{}
	for i := 0; i < n; i++ {
		if pairs[s[i]] > 0 {
			if len(stack) == 0 || stack[len(stack)-1] != pairs[s[i]] {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}
	return len(stack) == 0
}
