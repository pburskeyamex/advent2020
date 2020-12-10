package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type program struct {
	callStack         []*operationContext
	orderedOperations []*operationPayload

	acc_operation *acc
	noc_operation *command
	jmp_operation *command
}

func (p *program) lastCall() *operationContext {
	var context *operationContext
	length := len(p.callStack)
	if length > 0 {
		context = p.callStack[length-1]
	}
	return context
}

type operationContext struct {
	accumulator      int
	operationPayload *operationPayload
}

type command interface {
	Process(aProgram program, payload *operationPayload) *operationContext
}
type acc struct {
}

func (op *acc) Process(aProgram program, payload *operationPayload) *operationContext {

	lastCall := aProgram.lastCall()

	context := operationContext{
		accumulator:      0,
		operationPayload: payload,
	}

}

type nop struct {
}

func (op *nop) Process(aProgram program, payload *operationPayload) *operationContext {

}

type jmp struct {
}

func (op *jmp) Process(aProgram program, payload *operationPayload) *operationContext {

}

type operation struct {
	name    string
	command *command
}

type operationPayload struct {
	operation *operation
	argument  int
}

func main() {

	data := Parse("day_8_sample_data.txt")

	program := compile(data)

}

func compile(data []string) *program {

	aProgram := program{
		callStack:         nil,
		orderedOperations: nil,
		acc_operation:     nil,
		noc_operation:     nil,
		jmp_operation:     nil,
	}

	return &aProgram
}

func Parse(filename string) []string {
	file, err := os.Open(fmt.Sprintf("data/%w", filename))
	if err != nil {
		panic(err)
	}
	defer file.Close()

	data := make([]string, 0)
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	for i := 0; fileScanner.Scan(); i++ {

		aString := fileScanner.Text()
		data = append(data, aString)

	}

	file.Close()

	return data
}
