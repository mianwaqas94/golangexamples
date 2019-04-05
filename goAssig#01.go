package golangexamples

import (
	"fmt"

	"github.com/ehteshamz/greetings"
)

//ConcatSlice The first function
func ConcatSlice(sliceToConcat []byte) string {

	var concatSlice string
	for _, slice := range sliceToConcat {

		concatSlice += string(slice) + "-"
	}

	return concatSlice
}

//Encrpt Function
func Encrypt(sliceToEncrypt []byte, ceaserCount int) {
	encrypted := ""
	for i := 0; i < len(sliceToEncrypt); i++ {
		encrypted = encrypted + string(int(sliceToEncrypt[i])+ceaserCount)
	}
	fmt.Println(encrypted)

}

//EZGreetings is a global function
func EZGreetings(name string) string {
	return (greetings.PrintGreetings(name))

}
