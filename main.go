package main

import (
	"fmt"
	"net"
	"os"
	"sync"
	"time"
)

var commonPorts = []int{21, 22, 23, 25, 53, 80, 110, 139, 143, 443, 445, 3306, 3389, 5432, 8080}

func scanPort(target string, port int, wg *sync.WaitGroup, results chan<- string) {
	defer wg.Done()
	address := fmt.Sprintf("%s:%d", target, port)
	conn, err := net.DialTimeout("tcp", address, 2*time.Second)
	if err == nil {
		conn.Close()
		results <- fmt.Sprintf("[OPEN] Port %d on %s", port, target)
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: vuln-scanner <target-ip/domain>")
		os.Exit(1)
	}
	target := os.Args[1]

	fmt.Printf("Starting Vulnerability Scan on %s...\n", target)
	fmt.Println("--------------------------------------------------")

	var wg sync.WaitGroup
	results := make(chan string, len(commonPorts))

	startTime := time.Now()

	for _, port := range commonPorts {
		wg.Add(1)
		go scanPort(target, port, &wg, results)
	}

	wg.Wait()
	close(results)

	for res := range results {
		fmt.Println(res)
	}

	fmt.Println("--------------------------------------------------")
	fmt.Printf("Scan completed in %v\n", time.Since(startTime))
}
