package main

import (
	"fmt"
	"io"
)

type location struct {
	col  uint
	line uint
}

type command byte

const (
	shiftRightCommand command = '>'
	shiftLeftCommand  command = '<'
	incrementCommand  command = '+'
	decrementCommand  command = '-'
	putCommand        command = '.'
	getCommand        command = ','
	whileStartCommand command = '['
	whileEndCommand   command = ']'
)

type token struct {
	value command
	loc   location
}

func isCommand(b byte) bool {
	switch b {
	case '<':
		break
	case '>':
		break
	case '-':
		break
	case '+':
		break
	case '.':
		break
	case ',':
		break
	case '[':
		break
	case ']':
		break
	default:
		return false
	}

	return true
}

func newToken(b byte, loc *location) (*token, bool) {
	t := token{loc: *loc}

	switch b {
	case '<':
		t.value = shiftLeftCommand
	case '>':
		t.value = shiftRightCommand
	case '-':
		t.value = decrementCommand
	case '+':
		t.value = incrementCommand
	case '.':
		t.value = putCommand
	case ',':
		t.value = getCommand
	case '[':
		t.value = whileStartCommand
	case ']':
		t.value = whileEndCommand
	default:
		return nil, false
	}

	return &t, true
}

func (t *token) toString() string {
	return string(t.value)
}

func lex(source io.Reader) (*[]*token, error) {
	tokens := []*token{}
	buff := make([]byte, 1)
	var line uint = 1
	var column uint = 0
	var dontAllowSameRowNextToken bool
	var prevChar byte
	var prevToken token

	for {
		n, err := source.Read(buff)
		if err != nil && err != io.EOF {
			return nil, err
		}
		if err != nil && err == io.EOF && n == 0 {
			break
		}

		c := buff[0]

		switch c {
		case '\n':
			line++
			column = 0
			continue
		default:
			column++
		}

		currToken, ok := newToken(c, &location{column, line})
		if ok {
			// We moved to new line so we don't care about previous line invalid chars
			if dontAllowSameRowNextToken == true && prevToken.loc.line != currToken.loc.line {
				dontAllowSameRowNextToken = false
			}
			// We found token in the same line after some invalid characters
			if dontAllowSameRowNextToken == true {
				return nil, fmt.Errorf(
					"unexpected token '%s' at %d:%d",
					currToken.toString(), currToken.loc.line, currToken.loc.col,
				)
			}
			tokens = append(tokens, currToken)
		}

		// If the previous character and current are not tokens
		// If the previous token is in the same line as in the invalid token
		// All next tokens in this line are invalid
		// input is invalid
		// example:
		// +>++ some text    (this is valid)
		// +>++ some text <- (this is invalid)
		if prevChar != 0 && !isCommand(prevChar) && !isCommand(c) {
			dontAllowSameRowNextToken = true
		}

		prevChar = c

		if err == io.EOF {
			break
		}
	}

	return &tokens, nil
}

const maxArrayStack = 3000

func PrettyPrint(tokens *[]*token) {
	str := ""
	var isInLoop bool
	var prevToken *token

	for _, t := range *tokens {
		if t.value == whileEndCommand {
			isInLoop = false
		}
		if prevToken == nil {
			prevToken = t
			str += prevToken.toString()
		} else {
			if prevToken.loc.line != t.loc.line {
				str += "\n"
			}
			if isInLoop == true && t.loc.col == 1 {
				str += "  "
			}
			str += t.toString()
			prevToken = t
		}

		if t.value == whileStartCommand {
			isInLoop = true
		}
	}
	str += "\n"
}
