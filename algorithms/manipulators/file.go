package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("./server.log")
	if err != nil {
		fmt.Printf("%v", err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(f)
	max_latency := 0
	max_user := ""
	users := []string{}

	for scanner.Scan() {
		line := scanner.Text()

		spl := strings.Split(line, ",")

		user := spl[1]
		latency := strings.Replace(spl[3], "ms", "", -1)
		latency_number, _ := strconv.Atoi(latency)

		//fmt.Printf("%v\n", latency_number)

		if latency_number > max_latency {
			users = remove(users, max_user)
			max_latency = latency_number
			max_user = user
			users = append(users, max_user)
		} else if latency_number == max_latency {
			max_user = user
			users = append(users, max_user)
		}

	}
	fmt.Println(users)
}

func remove(slice []string, s string) []string {
	ind := 0
	for idx, v := range slice {
		if v == s {
			ind = idx
		}
	}
	if len(slice) <= 1 {
		slice = []string{}
	} else {
		slice = append(slice[:ind], slice[ind+1:]...)
	}

	return slice
}
