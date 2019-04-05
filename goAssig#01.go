package golangexamples

import "github.com/ehteshamz/greetings"

//ConcatSlice The first function
func ConcatSlice(sliceToConcat []byte) string {

	var concatSlice string
	for _, slice := range sliceToConcat {

		concatSlice += string(slice) + "-"
	}

	return concatSlice
}

//Encrypt.... THe Second Function
func Enrypt(sliceToEncrypt []byte, ceaserCount int) []byte {
	for i := 0; i < len(sliceToEncrypt); i++ {
		if sliceToEncrypt[i]+byte(ceaserCount) >= 122 {
			sliceToEncrypt[i] = ((sliceToEncrypt[i]+byte(ceaserCount))%122 + 96)
		} else {
			sliceToEncrypt[i] = (sliceToEncrypt[i] + byte(ceaserCount))
		}

	}

	return sliceToEncrypt
}

//EZGreetings is a global function
func EZGreetings(name string) string {
	return (greetings.PrintGreetings(name))

}
