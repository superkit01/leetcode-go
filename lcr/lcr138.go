/*
*

有效数字（按顺序）可以分成以下几个部分：

若干空格
一个 小数 或者 整数
（可选）一个 'e' 或 'E' ，后面跟着一个 整数
若干空格
小数（按顺序）可以分成以下几个部分：

（可选）一个符号字符（'+' 或 '-'）
下述格式之一：
至少一位数字，后面跟着一个点 '.'
至少一位数字，后面跟着一个点 '.' ，后面再跟着至少一位数字
一个点 '.' ，后面跟着至少一位数字
整数（按顺序）可以分成以下几个部分：

（可选）一个符号字符（'+' 或 '-'）
至少一位数字
部分有效数字列举如下：["2", "0089", "-0.1", "+3.14", "4.", "-.9", "2e10", "-90E3", "3e+7", "+6e-1", "53.5e93", "-123.456e789"]

部分无效数字列举如下：["abc", "1a", "1e", "e3", "99e2.5", "--6", "-+3", "95a54e53"]

给你一个字符串 s ，如果 s 是一个 有效数字 ，请返回 true 。

# DFA 确定有限状态自动机

*
*/
package lcr

func validNumber(s string) bool {
	state := STATE_START

	for _, c := range s {
		charType := getCharType(byte(c))
		if charType == -1 {
			return false
		}
		if nextState, ok := stateMachine[state][charType]; ok {
			state = nextState
		} else {
			return false
		}
	}

	if nextState, ok := stateMachine[state][SPACE]; ok {
		return nextState == STATE_END
	} else {
		return false
	}
}

var stateMachine = buildDFA() 

const (
	SYMBOL   = iota //  + -
	NUMBER          // 0-9
	POINT           // .
	EXPONENT        // e E
	SPACE           // ' '
)

func getCharType(c byte) int {
	switch c {
	case '+', '-':
		return SYMBOL
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		return NUMBER
	case '.':
		return POINT
	case 'e', 'E':
		return EXPONENT
	case ' ':
		return SPACE
	default:
		return -1
	}
}

const (
	STATE_START      = iota
	STATE_SYMBOL     // + -
	STATE_NUMBER     // 0-9
	STATE_POINT      // .
	STATE_FRACTION   // 小数部分
	STATE_EXPONENT   // e E
	STATE_EXP_SYMBOL // e E 后的 + -
	STATE_EXP_NUMBER // e E 后的 0-9
	STATE_END
)

func buildDFA() map[int]map[int]int {
	stateMachine := make(map[int]map[int]int)

	stateMachine[STATE_START] = map[int]int{
		SYMBOL: STATE_SYMBOL,
		NUMBER: STATE_NUMBER,
		POINT:  STATE_POINT,
		SPACE:  STATE_START,
	}
	stateMachine[STATE_SYMBOL] = map[int]int{
		NUMBER: STATE_NUMBER,
		POINT:  STATE_POINT,
	}
	stateMachine[STATE_NUMBER] = map[int]int{
		NUMBER:   STATE_NUMBER,
		POINT:    STATE_FRACTION,
		EXPONENT: STATE_EXPONENT,
		SPACE:    STATE_END,
	}
	stateMachine[STATE_POINT] = map[int]int{
		NUMBER: STATE_FRACTION,
	}
	stateMachine[STATE_FRACTION] = map[int]int{
		NUMBER:   STATE_FRACTION,
		EXPONENT: STATE_EXPONENT,
		SPACE:    STATE_END,
	}
	stateMachine[STATE_EXPONENT] = map[int]int{
		SYMBOL: STATE_EXP_SYMBOL,
		NUMBER: STATE_EXP_NUMBER,
	}
	stateMachine[STATE_EXP_SYMBOL] = map[int]int{
		NUMBER: STATE_EXP_NUMBER,
	}
	stateMachine[STATE_EXP_NUMBER] = map[int]int{
		NUMBER: STATE_EXP_NUMBER,
		SPACE:  STATE_END,
	}
	stateMachine[STATE_END] = map[int]int{
		SPACE: STATE_END,
	}

	return stateMachine

}
