package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type writerMock struct {
	IsCalled bool
	Value    interface{}
}

func (w *writerMock) Output(str interface{}) {
	w.IsCalled = true
	w.Value = str
}

func TestParse(t *testing.T) {
	tokens := []*token{
		{incrementCommand, location{1, 2}},
		{incrementCommand, location{2, 2}},
		{shiftRightCommand, location{1, 3}},
		{incrementCommand, location{3, 3}},
		{incrementCommand, location{4, 3}},
		{incrementCommand, location{5, 3}},
		{incrementCommand, location{6, 3}},
		{incrementCommand, location{7, 3}},
		{whileStartCommand, location{1, 5}},
		{shiftLeftCommand, location{1, 6}},
		{incrementCommand, location{3, 6}},
		{shiftRightCommand, location{1, 7}},
		{decrementCommand, location{3, 7}},
		{whileEndCommand, location{1, 8}},
		{incrementCommand, location{1, 16}},
		{incrementCommand, location{2, 16}},
		{incrementCommand, location{3, 16}},
		{incrementCommand, location{4, 16}},
		{incrementCommand, location{6, 16}},
		{incrementCommand, location{7, 16}},
		{incrementCommand, location{8, 16}},
		{incrementCommand, location{9, 16}},
		{whileStartCommand, location{1, 17}},
		{shiftLeftCommand, location{1, 18}},
		{incrementCommand, location{3, 18}},
		{incrementCommand, location{4, 18}},
		{incrementCommand, location{5, 18}},
		{incrementCommand, location{7, 18}},
		{incrementCommand, location{8, 18}},
		{incrementCommand, location{9, 18}},
		{shiftRightCommand, location{1, 19}},
		{decrementCommand, location{3, 19}},
		{whileEndCommand, location{1, 20}},
		{shiftLeftCommand, location{1, 21}},
		{putCommand, location{3, 21}},
	}

	mock := writerMock{}

	err := Parse(&tokens, &mock)
	assert.NoError(t, err)
	assert.Equal(t, mock.Value, 55)
}
