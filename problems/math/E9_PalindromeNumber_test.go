package math

import "testing"

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		input    int
		expected bool
	}{
		{121, true},          // 正数回文
		{-121, false},        // 负数不是回文
		{10, false},          // 非回文
		{12321, true},        // 长回文数
		{1, true},            // 单一数字
		{0, true},            // 0 是回文
		{1221, true},         // 偶数位回文
		{123, false},         // 非回文
		{-2147447412, false}, // 负数不是回文
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			result := isPalindrome(tt.input)
			if result != tt.expected {
				t.Errorf("isPalindrome(%d) = %v; want %v", tt.input, result, tt.expected)
			}
		})
	}
}
