package warmup

import (
	"fmt"
)

/*
 * Complete the 'countingValleys' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER steps
 *  2. STRING path
 */

func countingValleys(steps int32, path string) int32 {
	// Write your code here
	var numOfValleys int32
	var pathArray = []rune(path)
	hikerLevel := 0

	for _, character := range pathArray {
		if character == 'U' {
			hikerLevel += 1
		}

		if character == 'D' {
			hikerLevel -= 1
		}

		if hikerLevel == 0 && character == 'U' {
			numOfValleys++
		}
	}

	return numOfValleys
}

func CountingValleys() {
	//reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	//stepsTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	//checkError(err)
	//steps := int32(stepsTemp)
	//
	//path := readLine(reader)

	result := countingValleys(8, "UDDDUDUU")

	fmt.Println(result)
}
