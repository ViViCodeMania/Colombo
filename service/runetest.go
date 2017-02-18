package main

import (
	"unicode/utf8"
	"fmt"
)

func main() {
	const nihongo = "高性能linux服务器构建实战"
	fmt.Println(getStringLengthInTerminal(nihongo))

}

//func getStringLengthInTerminal(s string) int {
//	width := 0
//	for _, c := range s {
//		if utf8.RuneLen(c) > 2 {
//			width += 2
//		} else {
//			width += 1
//		}
//	}
//	return width
//}
