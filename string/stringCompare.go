package main

import "fmt"

func optimalSolution(string1, string2 string) bool {

	p1 := len(string1) - 1
	p2 := len(string2) - 1

	for p1 >= 0 || p2 >= 0 {

		if string1[p1] == '#' || string2[p2] == '#' {
			if string1[p1] == '#' {
				backCount := 2

				for backCount > 0 {
					p1--
					backCount--

					if string1[p1] == '#' {
						backCount += 2
					}
				}
			}

			if string2[p2] == '#' {
				backCount := 2

				for backCount > 0 {
					p2--
					backCount--

					if string2[p2] == '#' {
						backCount += 2
					}
				}
			}
		} else {

			if string1[p1] != string2[p2] {
				return false
			}

			p1--
			p2--
		}
	}
	return true
}

func backspaceCompare(s string, t string) bool {
	p1 := len(s) - 1
	p2 := len(t) - 1

	for p1 >= 0 || p2 >= 0 {

		if s[p1] == '#' || t[p2] == '#' {
			if s[p1] == '#' {
				backCount := 2

				for backCount > 0 {
					p1--
					backCount--

					if p1 > 0 && s[p1] == '#' {
						backCount += 2
					}
				}
			}

			if t[p2] == '#' {
				backCount := 2

				for backCount > 0 {
					p2--
					backCount--

					if p2 > 0 && t[p2] == '#' {
						backCount += 2
					}
				}
			}
		} else {

			if s[p1] != t[p2] {
				return false
			}

			p1--
			p2--
		}
	}
	return true
}
func main() {

	isEqual := backspaceCompare("ab##", "c#d#")

	if isEqual {
		fmt.Println("The strings are equal")
	} else {
		fmt.Println("The strings are not equal")
	}
}
