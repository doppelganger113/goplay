package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type loop struct {
	startIndex uint
	endIndex   uint
}

type outWriter interface {
	Output(str interface{})
}

func Parse(tokens *[]*token, writer outWriter) error {
	reader := bufio.NewReader(os.Stdin)
	cells := [maxArrayStack]int{}
	column := 0
	ptr := &cells[column]

	loopStack := []*loop{}

	for i := uint(0); i < uint(len(*tokens)); i++ {
		if *ptr < 0 {
			return fmt.Errorf("ptr should not be bellow 0")
		}

		tok := (*tokens)[i]
		switch tok.value {
		case incrementCommand:
			*ptr++
		case decrementCommand:
			*ptr--
		case shiftLeftCommand:
			if column < 1 {
				return fmt.Errorf("err: cannot shift left bellow 0\t%d:%d", tok.loc.line, tok.loc.col)
			}
			column--
			ptr = &cells[column]
		case shiftRightCommand:
			if column >= maxArrayStack {
				return fmt.Errorf(
					"err: cannot shift right above %d\t%d:%d",
					maxArrayStack, tok.loc.line, tok.loc.col,
				)
			}
			column++
			ptr = &cells[column]
		case getCommand:
			text, err := reader.ReadString('\n')
			if err != nil {
				return err
			}
			trimmed := strings.Trim(text, "\n")
			num, err := strconv.Atoi(trimmed)
			if err != nil {
				return err
			}
			*ptr = num
		case putCommand:
			writer.Output(*ptr)
		case whileStartCommand:
			// If ptr is 0, then we skip to the end of the loop
			if *ptr == 0 {
				// Jump after the end of the loop
				lastLoop := loopStack[len(loopStack)-1]
				// i will get incremented because of loop
				i = lastLoop.endIndex
				// Pop it
				loopStack = loopStack[:len(loopStack)-1]
				continue
			}
			loopStack = append(loopStack, &loop{
				startIndex: i,
				endIndex:   0,
			})
		case whileEndCommand:
			// Set last stack loop info to know where the end of the loop is
			lastLoop := loopStack[len(loopStack)-1]
			// Always redirect back to start of the loop
			// Decrement because loop increments i
			if lastLoop.endIndex == uint(0) {
				lastLoop.endIndex = i
			}
			i = lastLoop.startIndex - 1
		}
	}

	return nil
}
