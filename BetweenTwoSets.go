package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the getTotalX function below.
 */
func getTotalX(a []int32, b []int32) int32 {
	var countOfNumbers int32
	var arrayA []int32
	var arrayB []int32
	var match bool
	var index int32
	for index = 0; index < 100; index++ {
		match = false
		for index1 := 0; index1 < len(a); index1++ {
			arrayAElement := a[index1]
			if (index+1)%arrayAElement == 0 {
				match = true
			} else {
				match = false
				break
			}
		}
		if match == true {

			arrayA = append(arrayA, index+1)

		}
	}

	for index := 0; index < len(arrayA); index++ {
		match = false
		for index1 := 0; index1 < len(b); index1++ {
			arrayBElement := b[index1]
			if arrayBElement%arrayA[index] == 0 {
				match = true
			} else {
				match = false
				break
			}

		}
		if match == true {
			arrayB = append(arrayB, arrayA[index])
		}
	}

	countOfNumbers = int32(len(arrayB))

	return countOfNumbers

}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	outputFile, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer outputFile.Close()

	writer := bufio.NewWriterSize(outputFile, 1024*1024)

	nm := strings.Split(readLine(reader), " ")

	nTemp, err := strconv.ParseInt(nm[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	mTemp, err := strconv.ParseInt(nm[1], 10, 64)
	checkError(err)
	m := int32(mTemp)

	aTemp := strings.Split(readLine(reader), " ")

	var a []int32

	for aItr := 0; aItr < int(n); aItr++ {
		aItemTemp, err := strconv.ParseInt(aTemp[aItr], 10, 64)
		checkError(err)
		aItem := int32(aItemTemp)
		a = append(a, aItem)
	}

	bTemp := strings.Split(readLine(reader), " ")

	var b []int32

	for bItr := 0; bItr < int(m); bItr++ {
		bItemTemp, err := strconv.ParseInt(bTemp[bItr], 10, 64)
		checkError(err)
		bItem := int32(bItemTemp)
		b = append(b, bItem)
	}

	total := getTotalX(a, b)

	fmt.Fprintf(writer, "%d\n", total)

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
