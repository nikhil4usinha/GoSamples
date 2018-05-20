package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the diagonalDifference function below.
func diagonalDifference(a [][]int32) int32 {
	var difference int32
	var sum1 int32 = 0
	var sum2 int32 = 0
	var lengthArray int
	lengthArray = len(a)
	var i int = lengthArray - 1
	for index := 0; index < lengthArray; index++ {
		sum1 += a[index][index]
		sum2 += a[i][index]
		i = i - 1
	}
	if sum1 > sum2 {
		difference = sum1 - sum2
	}
	if sum2 >= sum1 {
		difference = sum2 - sum1
	}

	return difference

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

	var a [][]int32
	for i := 0; i < int(n); i++ {
		aRowTemp := strings.Split(readLine(reader), " ")

		var aRow []int32
		for _, aRowItem := range aRowTemp {
			aItemTemp, err := strconv.ParseInt(aRowItem, 10, 64)
			checkError(err)
			aItem := int32(aItemTemp)
			aRow = append(aRow, aItem)
		}

		if len(aRow) != int(int(n)) {
			panic("Bad input")
		}

		a = append(a, aRow)
	}

	result := diagonalDifference(a)

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
		panic(err)
	}
}
