package warmup

import (
	"fmt"
	"hackerrank/algorithm/sorting"
)

/*
 * Complete the 'sockMerchant' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER n
 *  2. INTEGER_ARRAY ar
 */
/**
 n = 7
ar = [1,2,1,2,1,3,2]
*/

/**
reassign to the array to remove the  element at index for array
*/
func sockMerchant(n int32, ar []int32) int32 {

	var pairsCount int32 = 0

	// Write your code here
	sorting.BubbleSort[int32](ar)
	fmt.Println(ar)

	for i := 0; i < int(n); i++ {

		fmt.Println(i)
		if ar[i] == ar[i+1] {
			pairsCount++
			if i < int(n) {
				i++
			}
		}
	}
	return pairsCount
}

func SockMerchantProblem() {
	arr := []int32{20, 20, 30, 30, 405, 505, 405}
	pairs := sockMerchant(int32(len(arr)), arr)
	fmt.Println(pairs)
	//reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
	//
	//stdout, err := os.Create("random")
	//checkError(err)
	//
	//defer checkError(stdout.Close())
	//
	//writer := bufio.NewWriterSize(stdout, 16*1024*1024)
	//
	//nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	//checkError(err)
	//n := int32(nTemp)
	//
	//arTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")
	//
	//var ar []int32
	//
	//for i := 0; i < int(n); i++ {
	//	arItemTemp, err := strconv.ParseInt(arTemp[i], 10, 64)
	//	checkError(err)
	//	arItem := int32(arItemTemp)
	//	ar = append(ar, arItem)
	//}
	//
	//result := sockMerchant(n, ar)
	//
	//_, err = fmt.Fprintf(writer, "%d\n", result)
	//
	//checkError(err)
	//
	//err = writer.Flush()
	//checkError(err)
}
