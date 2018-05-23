package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the breakingRecords function below.
func breakingRecords(score []int32) []int32 {
	var highestScore int32 = score[0]
	var lowestScore int32 = score[0]

	var highestScoreCount int32 = 0
	var lowestScoreCount int32 = 0
	//10 5 20 20 4 5 2 25 1
	for index := 1; index < len(score); index++ {
		if score[index] > highestScore {
			highestScoreCount++
			highestScore = score[index]
		} else if score[index] < lowestScore {
			lowestScoreCount++
			lowestScore = score[index]
		}
	}

	report := []int32{highestScoreCount, lowestScoreCount}

	return report

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

	scoreTemp := strings.Split(readLine(reader), " ")

	var score []int32

	for i := 0; i < int(n); i++ {
		scoreItemTemp, err := strconv.ParseInt(scoreTemp[i], 10, 64)
		checkError(err)
		scoreItem := int32(scoreItemTemp)
		score = append(score, scoreItem)
	}

	result := breakingRecords(score)

	for i, resultItem := range result {
		fmt.Fprintf(writer, "%d", resultItem)

		if i != len(result)-1 {
			fmt.Fprintf(writer, " ")
		}
	}

	fmt.Fprintf(writer, "\n")

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
