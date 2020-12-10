package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"
	"runtime"
	"strconv"
	"strings"
)

type command func(aProgram *program, operationID int, previousOperation *operationContext, currentOperation *operationPayload) (int, *operationContext)

type program struct {
	callStack         []*operationContext
	orderedOperations []*operationPayload
	accumulator       int
	acc               command
	nop               command
	jmp               command
}

func (p *program) execute() {

	p.accumulator = 0
	nextStep := 0

	success := p.executeWithinContext(nextStep, make([]int, len(p.orderedOperations)))

	iteration := 0
	for !success {
		p.accumulator = 0
		p.callStack = nil
		nextStep := 0
		p.switchOperation(iteration)

		success = p.executeWithinContext(nextStep, make([]int, len(p.orderedOperations)))

		if !success {
			p.switchOperation(iteration)
			iteration++
		} else {
			log.Println(fmt.Sprintf("Successful Munging! Current Accumulator value: %d Operation ID: %d", p.accumulator, humanOperationID(iteration)))
		}

		//if iteration > len(p.callStack){
		//	log.Fatal("Bummer")
		//}
	}

}

func (p *program) switchOperation(operationID int) bool {
	success := false

	accFunctionName := pointerName(p.acc)
	jmpFunctionName := pointerName(p.jmp)
	nopFunctionName := pointerName(p.nop)

	payload := p.orderedOperations[operationID]
	if pointerName(payload.instruction) == accFunctionName {
		// we are leaving it alone!
	} else if pointerName(payload.instruction) == jmpFunctionName {
		payload.instruction = p.nop
	} else if pointerName(payload.instruction) == nopFunctionName {
		payload.instruction = p.jmp
	}

	return success
}

func pointerName(aFunction interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(aFunction).Pointer()).Name()
}

func humanOperationID(anOperationID int) int {
	return anOperationID + 1
}
func (p *program) executeWithinContext(operationID int, payloadContext []int) bool {
	if operationID > len(p.orderedOperations)-1 {
		return true
	}
	payload := p.orderedOperations[operationID]

	var previousOperationContext *operationContext
	if (len(p.callStack) - 1) >= operationID {
		previousOperationContext = p.callStack[operationID]
	}

	payloadContext[operationID]++

	if payloadContext[operationID] > 1 {
		log.Println("endless loop detected")
		log.Println(fmt.Sprintf("Current Accumulator value: %d Operation ID: %d", p.accumulator, humanOperationID(operationID)))
		return false
	}

	var nextStep int
	var anOperationContext *operationContext
	nextStep, anOperationContext = payload.instruction(p, operationID, previousOperationContext, payload)
	p.callStack = append(p.callStack, anOperationContext)
	p.accumulator = anOperationContext.accumulator
	return p.executeWithinContext(nextStep, payloadContext)
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

func accProcess(aProgram *program, operationID int, previousOperation *operationContext, currentOperation *operationPayload) (int, *operationContext) {

	value := aProgram.accumulator
	value = value + currentOperation.argument

	context := operationContext{
		accumulator:      value,
		operationPayload: currentOperation,
	}

	nextOperationID := operationID + 1

	//log.Println(logOperation(operationID, "ACC", aProgram, &context))
	return nextOperationID, &context

}

func nopProcess(aProgram *program, operationID int, previousOperation *operationContext, currentOperation *operationPayload) (int, *operationContext) {

	value := aProgram.accumulator

	context := operationContext{
		accumulator:      value,
		operationPayload: currentOperation,
	}

	nextOperationID := operationID + 1

	//log.Println(logOperation(operationID, "NOP", aProgram, &context))
	return nextOperationID, &context
}

func logOperation(operationID int, operationName string, aProgram *program, instructionContext *operationContext) string {
	return fmt.Sprintf("Processing Operation ID: %d Instruction: %s %d Program Accumulator: %d Instruction Accumulator: %d", humanOperationID(operationID), operationName, instructionContext.operationPayload.argument, aProgram.accumulator, instructionContext.accumulator)
}

func jmpProcess(aProgram *program, operationID int, previousOperation *operationContext, currentOperation *operationPayload) (int, *operationContext) {

	value := aProgram.accumulator

	context := operationContext{
		accumulator:      value,
		operationPayload: currentOperation,
	}

	nextOperationID := operationID + currentOperation.argument

	//log.Println(logOperation(operationID, "JMP", aProgram, &context))
	return nextOperationID, &context
}

type operation struct {
	name    string
	command *command
}

type operationPayload struct {
	instruction command
	argument    int
	visits      int
}

func main() {

	data := Parse("day_8_data_2.txt")

	program := compile(data)

	program.execute()

}

func parseInstruction(aString string, aProgram *program) *operationPayload {
	anArray := strings.Split(aString, " ")
	instruction := anArray[0]
	value, _ := strconv.Atoi(anArray[1])

	var operation command
	if instruction == "acc" {
		operation = aProgram.acc
	} else if instruction == "jmp" {
		operation = aProgram.jmp
	} else if instruction == "nop" {
		operation = aProgram.nop
	}

	payload := operationPayload{
		instruction: operation,
		argument:    value,
	}
	return &payload

}

func compile(data []string) *program {

	aProgram := program{
		callStack:         make([]*operationContext, 0),
		orderedOperations: make([]*operationPayload, 0),
		nop:               nopProcess,
		jmp:               jmpProcess,
		acc:               accProcess,
	}

	for _, aString := range data {
		payload := parseInstruction(aString, &aProgram)
		aProgram.orderedOperations = append(aProgram.orderedOperations, payload)
	}

	return &aProgram
}

func Parse(aFilePart string) []string {
	filename := fmt.Sprintf("data/%s", aFilePart)
	file, err := os.Open(filename)
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
