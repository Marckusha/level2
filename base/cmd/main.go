package main

import (
	"fmt"
	"os"

	"github.com/beevik/ntp"
)

func main() {
	time, err := ntp.Time("0.beevik-ntp.pool.ntp.org")

	if err != nil {
		fmt.Fprintln(os.Stderr, "ERROR!", err)
		return
	}

	fmt.Println(time)

	args1 := []string{"a", "b"}
	args2 := []string{"c", "d"}
	copy(args1, args2)
	args1[0] = "rty"
	fmt.Println(args2)
	fmt.Println(args1)
}
