package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	line1 := strings.Split(readLine(reader), " ")

	sTemp, err := strconv.ParseInt(line1[0], 10, 64)
	checkError(err)
	totalOrders := int32(sTemp)

	tTemp, err := strconv.ParseInt(line1[1], 10, 64)
	checkError(err)
	orderNotEaten := int32(tTemp)

	orderArray := strings.Split(readLine(reader), " ")

	var orders []int32

	for i := 0; i < int(totalOrders); i++ {
		orderArrayTemp, err := strconv.ParseInt(orderArray[i], 10, 64)
		checkError(err)
		order := int32(orderArrayTemp)
		orders = append(orders, order)
	}

	line3 := readLine(reader)

	billTemp, err := strconv.ParseInt(line3, 10, 64)
	checkError(err)

	bill := int32(billTemp)

	CheckBonAppetite(totalOrders, orderNotEaten, orders, bill)

}

func CheckBonAppetite(totalOrders int32, orderNotEaten int32, orders []int32, bill int32) {
	var result string
	var totalBill int32
	for index := 0; index < len(orders); index++ {
		if index != int(orderNotEaten) {
			totalBill += orders[index]
		}
	}
	expectedBill := totalBill / 2
	billMisMatch := bill - expectedBill

	if billMisMatch == 0 {
		result = "Bon Appetit"
	} else {
		result = strconv.Itoa(int(billMisMatch))
	}

	fmt.Println(result)
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
