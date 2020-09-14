package config

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Display prints out config file in an appropriate format
func Display(minDelay, maxDelay int, id, ip, port []string) {
	fmt.Println("---Configuration File Read---")
	fmt.Printf("Minimum Delay Time: %d(ms)\n", minDelay)
	fmt.Printf("Maximum Delay Time: %d(ms)\n", maxDelay)
	fmt.Printf("IDs: %v\n", id)
	fmt.Printf("IPs: %v\n", ip)
	fmt.Printf("Ports: %v\n", port)
	fmt.Println("-----------------------------")
}

//Extracts the min, max, and ports from config file and returns them
func Extract(str []string) (int, int, []string, []string, []string) {
	time := strings.Split(str[0], " ")
	min, err := strconv.Atoi(time[0])
	if err != nil {
		fmt.Println(err)
	}
	max, err := strconv.Atoi(time[1])
	if err != nil {
		fmt.Println(err)
	}

	var ids, ips, ports []string

	for i := 1; i < len(str); i++ {
		temp := strings.Split(str[i], " ")
		ids = append(ids, temp[0])
		ips = append(ips, temp[1])
		ports = append(ports, temp[2])
	}

	return min, max, ids, ips, ports
}

//Line by line, stores elements of config file
func ReadConfig() []string {
	file, err := os.Open("config.txt")
	if err != nil {
		fmt.Println(err)
	}

	scanner := bufio.NewScanner(file)
	var str []string //array to store line by line

	for scanner.Scan() {
		str = append(str, scanner.Text())
	}

	return str
}
