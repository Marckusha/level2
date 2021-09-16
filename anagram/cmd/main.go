package main

import (
	"fmt"
	"level2/anagram/pkg/anagram"
)

func main() {
	str := []string{"пятка", "яткап", "листок", "конверт", "пятак", "слиток", "столик", "тяпка", "каптер", "паркет", "отец"}
	m := anagram.GetMapAnagram(str)
	fmt.Println(m)
}
