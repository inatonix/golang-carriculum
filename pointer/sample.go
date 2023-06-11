package main

type myStruct struct {
	value int
}

var testStruct = myStruct{value: 0}

func addTwoNumbers(a int, b int) int {
	return a + b
}
func myFunction() {
	testVar1 := 123
	testVar2 := 456
	testStruct.value = addTwoNumbers(testVar1, testVar2)
}
func someOtherFunction() {
	// some other code
	myFunction()
	// some more code
}
