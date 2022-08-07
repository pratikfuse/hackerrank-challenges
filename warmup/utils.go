package warmup

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func removeAtIndex(slice []int32, s int) []int32 {
	fmt.Println("removing at", s, slice)
	copy(slice[s:], slice[s+1:])
	slice[len(slice)-1] = 0
	return slice[:len(slice)-1]
}
func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
