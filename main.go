package main

import "fmt"

func main() {
	a1 := []string{"ejjjjmmtthh", "zxxuueeg", "aanlljrrrxx", "dqqqaaabbb", "oocccffuucccjjjkkkjyyyeehh"}
	a2 := []string{"bbbaaayddqbbrrrv"}
	fmt.Println(MxDifLg(a1, a2))

}

func MxDifLg(a1 []string, a2 []string) (result int) {
	if len(a1) == 0 || len(a2) == 0 {
		return -1
	}
	for _, x := range a1 {
		for _, y := range a2 {
			n := len(x) - len(y)
			if n < 0 {
				n *= -1
			}
			if n > result {
				result = n
			}

		}
	}
	return
}
