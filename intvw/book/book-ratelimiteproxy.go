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
 * Complete the 'droppedRequests' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER_ARRAY requestTime as parameter.
 */

func droppedRequests(requestTime []int32) int32 {
	// Write your code here
	var countDorpped int32
	var countRequest int32
	var reqTimeAc int32
	var reqTimeCount int32
	for _, reqtime := range requestTime {
		countRequest++
		if reqtime == reqTimeAc {
			reqTimeCount++
		}
		reqTimeAc = reqtime
		if reqTimeCount >= 3 || countRequest > 20 {
			countDorpped++
		}

	}

	return countDorpped

}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	requestTimeCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	var requestTime []int32

	for i := 0; i < int(requestTimeCount); i++ {
		requestTimeItemTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		requestTimeItem := int32(requestTimeItemTemp)
		requestTime = append(requestTime, requestTimeItem)
	}

	result := droppedRequests(requestTime)

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
