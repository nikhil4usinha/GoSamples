package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the solve function below.
func solve(year int32) string {
	daysInMonth := GetDaysInMonth(year)
	var day int32
	var month int32
	var sum int32
	for index := 0; index < len(daysInMonth)-1; index++ {
		sum += daysInMonth[index]
		if 256-sum < daysInMonth[index+1] {
			day = 256 - sum
			month = int32(index) + 2
			break
		}
	}
	var strDay string
	if day < 10 {
		strDay = "0" + strconv.Itoa(int(day))
	} else {
		strDay = strconv.Itoa(int(day))
	}

	var strMonth string
	if month < 10 {
		strMonth = "0" + strconv.Itoa(int(month))
	} else {
		strMonth = strconv.Itoa(int(month))
	}
	var dayInYear = strDay + "." + strMonth + "." + strconv.Itoa(int(year))
	//strconv.Itoa(int(day)) + "." + strconv.Itoa(int(month)) + "." + strconv.Itoa(int(year))

	return dayInYear
}

func GetDaysInMonth(year int32) [12]int32 {
	leapYear := CheckLeapYear(year)
	var daysInMonth [12]int32
	if year == 1918 {
		daysInMonth = [12]int32{31, 15, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	}
	if year < 1918 {
		if leapYear == true {
			daysInMonth = [12]int32{31, 29, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
		} else {
			daysInMonth = [12]int32{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
		}
	}
	if year > 1918 {
		if leapYear == true {
			daysInMonth = [12]int32{31, 29, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
		} else {
			daysInMonth = [12]int32{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
		}
	}

	return daysInMonth
}

func CheckLeapYear(year int32) bool {
	var leapYear = false
	if year <= 1918 {
		if year%4 == 0 {
			leapYear = true
		}
	} else if year > 1918 {
		if year%400 == 0 {
			leapYear = true
		}
		if year%4 == 0 && year%100 != 0 {
			leapYear = true
		}
	}
	return leapYear
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	yearTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	year := int32(yearTemp)

	result := solve(year)

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
