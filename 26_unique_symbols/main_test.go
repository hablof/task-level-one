package main

import "testing"

func Test_isUinqueSymbolsCaseInsensitive(t *testing.T) {
	tests := []struct {
		str  string
		want bool
	}{
		{
			str:  "Test_isUinqueSymbolsCaseInsensitive",
			want: false,
		},
		{
			str:  "aA",
			want: false,
		},
		{
			str:  "false",
			want: true,
		},
		{
			str:  "йцукенгшщзхъфывапролджэ\\ячсмитьбю.qwertyuiop[]{}asdfghjklxcvbnm,/'\"0123456789",
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.str, func(t *testing.T) {
			if got := isUinqueSymbolsCaseInsensitive(tt.str); got != tt.want {
				t.Errorf("isUinqueSymbolsCaseInsensitive() = %v, want %v", got, tt.want)
			}
		})
	}
}
