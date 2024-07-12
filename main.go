package main

import "fmt"

func main() {
	a1 := []string{"hoqq", "bbllkw", "oox", "ejjuyyy", "plmiis", "xxxzgpsssa", "xxwwkktt", "znnnnfqknaz", "qqquuhii", "dvvvwz"}
	a2 := []string{"cccooommaaqqoxii", "gggqaffhhh", "tttoowwwmmww"}
	fmt.Println(MxDifLg(a1, a2))

}

func MxDifLg(a1 []string, a2 []string) (result int) {
	for _, y := range a2 {
		for _, x := range a1 {
			n := len(y) - len(x)
			if n > result {
				result = n
			}

		}
	}
	return
}
