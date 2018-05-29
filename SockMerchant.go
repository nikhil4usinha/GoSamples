package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func contains(arr []int32, num int32) bool {
	for _, a := range arr {
		if a == num {
			return true
		}
	}
	return false
}

// Complete the sockMerchant function below.
func sockMerchant(n int32, ar []int32) int32 {
	var result int32
	var arrSum []int32

	var distinctArr []int32
	//10,20,20,10,30,20,10,10
	for index1 := 0; index1 < len(ar); index1++ {
		var sum int32
		if !contains(distinctArr, ar[index1]) {
			sum = 0
			number := ar[index1]
			for index := index1; index < len(ar); index++ {

				if ar[index] == number {
					sum++
				}

			}
			arrSum = append(arrSum, sum)
			distinctArr = append(distinctArr, ar[index1])

		}

	}

	for index := 0; index < len(arrSum); index++ {
		quotient := arrSum[index] / 2
		result += quotient

	}

	return result

}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	arTemp := strings.Split(readLine(reader), " ")

	var ar []int32

	for i := 0; i < int(n); i++ {
		arItemTemp, err := strconv.ParseInt(arTemp[i], 10, 64)
		checkError(err)
		arItem := int32(arItemTemp)
		ar = append(ar, arItem)
	}

	result := sockMerchant(n, ar)

	fmt.Fprintf(writer, "%d\n", result)

	writer.Flush()
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
		// panic(err)
	}
}
