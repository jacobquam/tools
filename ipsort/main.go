package main

import (
	"bufio"
	"bytes"
	"fmt"
	"net"
	"os"
	"sort"
)

func main() {

	// open file into fp
	fp, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fp.Close()

	// New scanner for fp
	scanner := bufio.NewScanner(fp)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	realIPs := make([]net.IP, 0, len(lines))
	for _, ip := range lines {
		realIPs = append(realIPs, net.ParseIP(ip))
	}

	sort.Slice(realIPs, func(i, j int) bool {
		return bytes.Compare(realIPs[i], realIPs[j]) < 0
	})

	fmt.Println("Sorted IPs")
	for in, line := range realIPs {
		fmt.Println(in, line)
	}

}
