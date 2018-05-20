package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the plusMinus function below.
func plusMinus(arr []int32) {
	var countPositive int32 = 0
	var countNegative int32 = 0
	var countZero int32 = 0
	var lenArr int = len(arr)

	for index := 0; index < lenArr; index++ {
		if arr[index] > 0 {
			countPositive++
		}
		if arr[index] < 0 {
			countNegative++
		}
		if arr[index] == 0 {
			countZero++
		}
	}

	var fractionPositive float32 = float32(countPositive) / float32(lenArr)
	var fractionNegative float32 = float32(countNegative) / float32(lenArr)
	var fractionZero float32 = float32(countZero) / float32(lenArr)

	fmt.Println(fractionPositive)
	fmt.Println(fractionNegative)
	fmt.Println(fractionZero)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	arrTemp := strings.Split(readLine(reader), " ")

	var arr []int32

	for i := 0; i < int(n); i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	plusMinus(arr)
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
