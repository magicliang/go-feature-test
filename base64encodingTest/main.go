package main

import (
"encoding/base64"
"fmt"
"log"
)

func main() {
	// 初始的字符串是 utf-8的 编码
	input := []byte("hello golang base64 王者天下 http://baidu.com +~")

	// base64 字符串才是真正的中间形式。encodeTo, decodeFrom。
	// encode 是字节 to string。
	encodeString := base64.StdEncoding.EncodeToString(input)
	// base 64 字符串是人类不可读的
	fmt.Println("encodeString: " + encodeString)
	// encodeString: aGVsbG8gZ29sYW5nIGJhc2U2NCDnjovogIXlpKnkuIsgaHR0cDovL2JhaWR1LmNvbSArfg==

	// decode 是string to 字节。解码的结果反而不是字符串。
	// 对上面的编码结果进行base64解码
	decodeBytes, err := base64.StdEncoding.DecodeString(encodeString)
	if err != nil {
		log.Fatalln(err)
	}

	// 打印字节不会自动转码为人类可读的encoding形式，就直接打出数字形式
	//fmt.Println(decodeBytes)

	// 原始字符串=原始字节转回string。是人类可读的
	fmt.Println("string(decodeBytes):" + string(decodeBytes))
	// 原字符串是很容易人类刻度的：string(decodeBytes):hello golang base64 王者天下 http://baidu.com +~

	fmt.Println()

	// 如果要用在url中，需要使用URLEncoding
	uEnc := base64.URLEncoding.EncodeToString(input)
	fmt.Println("uEnc: " + uEnc)

	// URLencoding的结果和 base64 encoding 的结果是一样的，因为都是base64包的编码函数。想要真正的url 编码要用url包的内容。
	// url encoding 里充满里百分号和两个十六进制数而 base64 编码则是连续26个字母/数字，以==结尾
	uDec, err := base64.URLEncoding.DecodeString(uEnc)
	if err != nil {
		log.Fatalln(err)
	}
	// 解码后的字符串还是人类刻度的
	fmt.Println("string(uDec):" + string(uDec))

	// 然而base64 是不同于 urlencode，它是用来传输二进制数据的（换言之所有的二进制字节都可以转化为base64编码，base64编码是为了和字节转换准备的，而不是和字符串转换做准备的）。
	// 可以看出urlencode才更适合url使用。urlencode可以转义url敏感字符而base64不可以。url的编码反而是字符串和字符串
	// https://segmentfault.com/q/1010000010231231
}
