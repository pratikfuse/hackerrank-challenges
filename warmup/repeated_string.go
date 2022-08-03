package warmup

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

/*
 * Complete the 'repeatedString' function below.
 *
 * The function is expected to return a LONG_INTEGER.
 * The function accepts following parameters:
 *  1. STRING s
 *  2. LONG_INTEGER n
 */

func repeatedString(s string, n int) int64 {
	// Write your code here

	countForA := 0

	str := s

	sLength := len(s)

	//strIndex := 0

	for _, character := range str {
		if character == 'a' {
			fmt.Println(string(character))
			countForA++
		}

		strLength := len(str)

		if strLength < n {

			if strLength+sLength <= n {
				str += s

			} else {
				str += s[:n-strLength]
			}
		}
	}
	fmt.Println(str)

	fmt.Println(countForA)
	return int64(countForA)
}

func RepeatedStrings() {

	repeatedString("aba", 13)
	//
	//reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
	//
	//stdout, err := os.Create("result")
	//
	//if err != nil {
	//	panic(err)
	//}
	//checkError(err)
	//
	//defer func(stdout *os.File) {
	//	err := stdout.Close()
	//	if err != nil {
	//
	//	}
	//}(stdout)
	//
	//writer := bufio.NewWriterSize(stdout, 16*1024*1024)
	//
	//s := readLine(reader)
	//
	//n, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	//checkError(err)
	//
	//result := repeatedString(s, n)
	//
	//_, err = fmt.Fprintf(writer, "%d\n", result)
	//if err != nil {
	//	return
	//}
	//
	//err = writer.Flush()
	//if err != nil {
	//	return
	//}
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
