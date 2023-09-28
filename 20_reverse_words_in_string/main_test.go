package main

import "testing"

func Test_reverseWords(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "",
			want: "",
		},
		{
			name: "Используйте этот текстовый инструмент, чтобы перевернуть список слов в обратном порядке. Первая строка становится последней, а последняя - первой",
			want: "первой - последняя а последней, становится строка Первая порядке. обратном в слов список перевернуть чтобы инструмент, текстовый этот Используйте",
		},
		{
			name: "every hunter wants to know where pheasant sits",
			want: "sits pheasant where know to wants hunter every",
		},
		{
			name: "Ожидаемое_поведение_программы_это_бездействие_ведь_пробелов_в_этой_строке_совсем_нет",
			want: "Ожидаемое_поведение_программы_это_бездействие_ведь_пробелов_в_этой_строке_совсем_нет",
		},
		{
			name: "Voyage Voyage",
			want: "Voyage Voyage",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseWords(tt.name); got != tt.want {
				t.Errorf("reverseWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
