package warmup

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
func repeatedString(s string, n int64) int64 {
	var count int64 = 0

	stringLength := int64(len(s))

	numRepetitions := n / stringLength
	remainder := n % stringLength
	for _, character := range s {
		if character == 'a' {
			count++
		}
	}

	count = count * numRepetitions
	for i := 0; i < int(remainder); i++ {
		if s[i] == 'a' {
			count++
		}
	}
	return count
}

func RepeatedString() {

	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create("result")

	checkError(err)

	defer func(stdout *os.File) {
		err := stdout.Close()
		if err != nil {
			checkError(err)
		}
	}(stdout)

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	s := readLine(reader)

	n, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	result := repeatedString(s, n)

	_, err = fmt.Fprintf(writer, "%d\n", result)

	checkError(err)

	err = writer.Flush()
	checkError(err)
}
