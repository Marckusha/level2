package main

import (
	"fmt"
	"level2/pkq/task2"
)

/*
"a4bc2d5e" => "aaaabccddddde"
"abcd" => "abcd"
"45" => "" (некорректная строка)
"" => ""
Дополнительное задание: поддержка escape - последовательностей
qwe\4\5 => qwe45 (*)
qwe\45 => qwe44444 (*)
qwe\\5 => qwe\\\\\ (*)
*/

func main() {
	var s string = "a4bc2d5e"
	res, _ := task2.Unpack(s)
	fmt.Println(res)
}
