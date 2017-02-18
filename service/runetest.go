package main

import (
	"unicode/utf8"
	"fmt"
)

func main() {
	s := "haha"
	//s := "haha"
	ru, rucount := utf8.DecodeRuneInString(s)
	fmt.Println(ru, " ", rucount, " ", utf8.RuneLen(ru))
}
