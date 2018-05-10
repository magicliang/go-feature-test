package main

import "fmt"

func main()  {
	compositeKeyNamespace := "\x00" // \x00
	minUnicodeRuneValue := 0 // => 正好对应 U+0000 这个码点
	fmt.Println(string(minUnicodeRuneValue) == compositeKeyNamespace)
}
