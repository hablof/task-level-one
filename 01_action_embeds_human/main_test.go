package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Embedding(t *testing.T) {
	actionStruct := Action{
		actionField1: "если бы мы знали, что это такое",
		actionField2: "но мы не знаем, что это такое",
		Human: Human{
			humanField1: 5,
			humanField2: 10,
			humanField3: "пятнадцатого",
		},
	}

	actionStruct.ChangeHumanFields()

	// Поля, относящиеся к типу структуры Human изменятся,
	// несмотря на то, что метод был вызван у структуры типа Action
	assert.Equal(t, 15, actionStruct.humanField1)
	assert.Equal(t, -5, actionStruct.humanField2)

	info := actionStruct.HumanInfo()

	assert.Equal(t, "human info: 15:пятнадцатого:-5", info)
}
