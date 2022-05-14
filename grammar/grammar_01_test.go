package grammar

import (
	"fmt"
	"testing"
)

//byte - 与uint8相同
//rune - 与int32相同
//uint - 32或64位
//int - 与uint大小相同
//uintptr - 无符号整数，用于存储指针值的未解释位
func TestGrammar01_00(t *testing.T) {
	var a byte = 'a'
	var b rune = 'b'
	var c uint = 1
	var d int = 2
	var e uintptr = 3
	fmt.Println(a, b, c, d, e)
}
