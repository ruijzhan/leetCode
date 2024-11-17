package string

import "math"

const (
	INT_MAX = math.MaxInt32
	INT_MIN = math.MinInt32
)

const (
	START = iota
	SIGNED
	IN_NUMBER
	END
)

type Automaton struct {
	state int
	sign  int
	ans   int
	table map[int][]int
}

func NewAutomaton() *Automaton {
	return &Automaton{
		state: START,
		sign:  1,
		ans:   0,
		table: map[int][]int{
			START:     {START, SIGNED, IN_NUMBER, END},
			SIGNED:    {END, END, IN_NUMBER, END},
			IN_NUMBER: {END, END, IN_NUMBER, END},
			END:       {END, END, END, END},
		},
	}
}

func (a *Automaton) getCol(c byte) int {
	if c == ' ' {
		return 0
	}
	if c == '+' || c == '-' {
		return 1
	}
	if c >= '0' && c <= '9' {
		return 2
	}
	return 3
}

func (a *Automaton) transition(c byte) {
	a.state = a.table[a.state][a.getCol(c)]
	if a.state == IN_NUMBER {
		a.ans = a.ans*10 + int(c-'0')
		if a.sign == 1 {
			a.ans = min(a.ans, INT_MAX)
		} else {
			a.ans = min(a.ans, -INT_MIN)
		}
	} else if a.state == SIGNED {
		if c == '+' {
			a.sign = 1
		} else {
			a.sign = -1
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func myAtoi(s string) int {
	automaton := NewAutomaton()
	for i := 0; i < len(s); i++ {
		automaton.transition(s[i])
	}
	return automaton.sign * automaton.ans
}
