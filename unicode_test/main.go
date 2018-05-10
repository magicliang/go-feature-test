package main

import "fmt"

func main()  {
	// 一个 unicode 是定长32字节的。
	compositeKeyNamespace := "\x00" // 但这种字面量却不超过单字节的表示范围。
	minUnicodeRuneValue := 0 // => 正好对应 U+0000 这个码点
	fmt.Println(string(minUnicodeRuneValue) == compositeKeyNamespace)

	/**
	"世界"
	//不带u开头的只能做单字节字面量。
"\xe4\xb8\x96\xe7\x95\x8c"
	// 带u的可以输入双字节字面量
"\u4e16\u754c"
"\U00004e16\U0000754c"
	 */
}
