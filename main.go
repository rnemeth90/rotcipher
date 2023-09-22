package main

import "fmt"

func main() {
	fmt.Println(RotationalCipher("A,bc", 1))
}

func RotationalCipher(plain string, shiftKey int) string {
	returnStr := ""

	for _, v := range plain {
		val := rotateRune(v, shiftKey)
		returnStr += string(val)
	}

	return returnStr
}

func rotateRune(c rune, n int) rune {
	if c >= 'a' && c <= 'z' {
		return ((c-'a'+rune(n))%26+26)%26 + 'a'
	} else if c >= 'A' && c <= 'Z' {
		return ((c-'A'+rune(n))%26+26)%26 + 'A'
	}
	return c
}
