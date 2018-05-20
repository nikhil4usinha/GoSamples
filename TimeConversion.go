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
 * Complete the timeConversion function below.
 */
func timeConversion(s string) string {
	var strTime24Hr string
	var strTime24HrArr []string
	if s == "12:00:00AM" {
		strTime24Hr = "00:00:00"
	}
	var timeArr []string = strings.Split(s, ":")
	var strTime string = timeArr[2][2:]

	var strHour string = timeArr[0]
	println(strHour)
	var strMin string = timeArr[1]
	var strSec string = timeArr[2][:2]

	intHour, err := strconv.Atoi(strHour)
	if strTime == "PM" {

		if intHour < 12 {

			intHour = intHour + 12
			strTime24HrArr = []string{strconv.Itoa(intHour), strMin, strSec}

			if err != nil {
				// handle error
				fmt.Println(err)
				os.Exit(2)
			}
		}
		if intHour >= 12 {
			strTime24HrArr = []string{strconv.Itoa(intHour), strMin, strSec}
		}
	}
	if strTime == "AM" {
		if intHour == 12 {
			strHour = "00"
		}
		strTime24HrArr = []string{strHour, strMin, strSec}
	}
	strTime24Hr = strings.Join(strTime24HrArr, ":")
	return strTime24Hr

}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	outputFile, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer outputFile.Close()

	writer := bufio.NewWriterSize(outputFile, 1024*1024)

	s := readLine(reader)

	result := timeConversion(s)

	fmt.Fprintf(writer, "%s\n", result)

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
