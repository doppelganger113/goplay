package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func Test_Lex(t *testing.T) {
	t.Log("Started")
	fmt.Println("Started")
	data := `
++       Cell c0 = 2
> +++++  Cell c1 = 5

[        Start your loops with your cell pointer on the loop counter (c1 in our case)
< +      Add 1 to c0
> -      Subtract 1 from c1
]        End your loops with the cell pointer on the loop counter

At this point our program has added 5 to 2 leaving 7 in c0 and 0 in c1
but we cannot output this value to the terminal since it is not ASCII encoded!

To display the ASCII character "7" we must add 48 to the value 7
48 = 6 * 8 so let's use another loop to help us!

++++ ++++  c1 = 8 and this will be our loop counter again
[
< +++ +++  Add 6 to c0
> -        Subtract 1 from c1
]
< .        Print out c0 which has the value 55 which translates to "7"!
`
	reader := strings.NewReader(data)

	tokens, err := lex(reader)
	if err != nil {
		t.Fatal("Error lexing: ", err)
	}

	expected := []*token{
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
	assert.ElementsMatchf(t, *tokens, expected, "error message %s", "formatted")
}
