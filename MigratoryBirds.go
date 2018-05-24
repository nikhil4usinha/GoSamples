package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the migratoryBirds function below.
func migratoryBirds(ar []int32) int32 {
	var type1 []int32
	var type2 []int32
	var type3 []int32
	var type4 []int32
	var type5 []int32

	for index := 0; index < len(ar); index++ {
		if ar[index] == 1 {
			type1 = append(type1, ar[index])
		}
		if ar[index] == 2 {
			type2 = append(type2, ar[index])
		}
		if ar[index] == 3 {
			type3 = append(type3, ar[index])
		}
		if ar[index] == 4 {
			type4 = append(type4, ar[index])
		}
		if ar[index] == 5 {
			type5 = append(type5, ar[index])
		}
	}

	type1Count := len(type1)
	type2Count := len(type2)
	type3Count := len(type3)
	type4Count := len(type4)
	type5Count := len(type5)

	var max = 1
	var maxType = type1Count
	if maxType < type2Count {
		max = 2
		maxType = type2Count
	}
	if maxType < type3Count {
		max = 3
		maxType = type3Count
	}
	if maxType < type4Count {
		max = 4
		maxType = type4Count
	}
	if maxType < type5Count {
		max = 5
		maxType = type5Count
	}

	return int32(max)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	arCount, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)

	arTemp := strings.Split(readLine(reader), " ")

	var ar []int32

	for i := 0; i < int(arCount); i++ {
		arItemTemp, err := strconv.ParseInt(arTemp[i], 10, 64)
		checkError(err)
		arItem := int32(arItemTemp)
		ar = append(ar, arItem)
	}

	result := migratoryBirds(ar)

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
