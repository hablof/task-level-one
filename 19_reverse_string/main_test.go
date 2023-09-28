package main

import "testing"

func Test_reverseUnicodeString(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "",
			want: "",
		},
		{
			name: "Онлайн конвертер, помогающий перевернуть текст задом наперед. Удобен при оформлении художественных текстов и исправлении очередности букв.",
			want: ".вкуб итсондеречо иинелварпси и вотскет хынневтсежодух иинелмрофо ирп небодУ .дерепан модаз тскет ьтунревереп йищюагомоп ,ретревнок нйалнО",
		},
		{
			name: "главрыба",
			want: "абырвалг",
		},
		{
			name: "паттерн «адаптер»",
			want: "»ретпада« нреттап",
		},
		{
			name: "Ну давай разберем по частям, тобою написанное ))",
			want: ")) еоннасипан юобот ,мятсач оп меребзар йавад уН",
		},
		{
			name: "Запасы копиума",
			want: "амуипок ысапаЗ",
		},
		{
			name: "Intel(R) Core(TM) i3-2100 CPU @ 3.10GHz",
			want: "zHG01.3 @ UPC 0012-3i )MT(eroC )R(letnI",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseUnicodeString(tt.name); got != tt.want {
				t.Errorf("reverseUnicodeString() = %v, want %v", got, tt.want)
			}
		})
	}
}
