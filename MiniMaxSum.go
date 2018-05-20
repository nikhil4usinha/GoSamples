package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the miniMaxSum function below.
func miniMaxSum(arr []int32) {
	var maxNo int64 = int64(arr[0])
	var minNo int64 = int64(arr[0])
	var maxSum int64 = 0
	var minSum int64 = 0
	var indexMaxNo int
	var indexMinNo int

	for index := 0; index < len(arr); index++ {
		if int64(arr[index]) >= maxNo {
			maxNo = int64(arr[index])
			indexMaxNo = index
		}

		if int64(arr[index]) <= minNo {
			minNo = int64(arr[index])
			indexMinNo = index
		}
	}

	for index := 0; index < len(arr); index++ {
		if indexMaxNo != index {
			minSum += int64(arr[index])
		}
		if indexMinNo != index {
			maxSum += int64(arr[index])
		}
	}
	fmt.Println(minSum, maxSum)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	arrTemp := strings.Split(readLine(reader), " ")

	var arr []int32

	for i := 0; i < len(arrTemp); i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	miniMaxSum(arr)
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
