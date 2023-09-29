package concat

import "testing"

var (
	sample = []string{"lal", "lol", "sample", "kek", "What is your address?", "What is your first name?", "lal", "lol", "mem", "параноя", "pop", "hih", "lal", "lol", "mem", "kek", "pop", "hih", "lal", "How do you spend your weekends?", "mem", "kek", "pop", "hih",
		"Where do you live?", "lol", "Валентин Витальевич", "kek", "pop", "hih", "lal", "lol", "mem", "Is your family large?", "pop", "hih", "lal", "lol", "mem", "kek", "What are your household duties?", "What do you usually do in the evenings?", "lal", "lol", "mem", "kek", "pop", "hih",
		"lal", "What is your surname?", "mem", "kek", "pop", "hih", "lal", "lol", "mem", "kek", "golang", "hih", "lal", "lol", "mem", "kek", "pop", "hih", "lal", "lol", "mem", "kek", "pop", "Do you always plan your day beforehand?",
		"lal", "lol", "mem", "kek", "pop", "Where were you born?", "lal", "lol", "What is your telephone number?", "kek", "Who do you most take after, your mother or your father?", "hih", "lal", "lol", "mem", "kek", "pop", "hih", "Are you a stay-at-home or do you prefer to go out when you have some time to spare?", "lol", "mem", "kek", "pop", "hih"}
)

func Benchmark_Fixed(b *testing.B) {

	s1 := "What is your address?"
	s2 := "What is your telephone number?"
	s3 := "lol kek cheburek"
	s4 := "До свидания Auf Wiedersehen"
	s5 := "Очень такая интересная пятая строка"
	s6 := "https://github.com/hablof/"

	b.Run("concatStrings", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = concatStrings_Fixed(s1, s2, s3, s4, s5, s6)
		}
	})

	b.Run("formatString", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = formatString_Fixed(s1, s2, s3, s4, s5, s6)
		}
	})

	b.Run("Builder_NoPreAlloc", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = stringsBuilder_Fixed_NoPreAlloc(s1, s2, s3, s4, s5, s6)
		}
	})

	b.Run("Builder_PreAlloc", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = stringsBuilder_Fixed_PreAlloc(s1, s2, s3, s4, s5, s6)
		}
	})
}

func Benchmark_NonFixed(b *testing.B) {
	b.Run("concatStrings", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = concatStrings_NonFixed(sample...)
		}
	})

	b.Run("formatString", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = formatString_NonFixed(sample...)
		}
	})

	b.Run("Builder_NoPreAlloc", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = stringsBuilder_NonFixed_NoPreAlloc(sample...)
		}
	})

	b.Run("Builder_PreAlloc", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = stringsBuilder_NonFixed_PreAlloc(sample...)
		}
	})
}
