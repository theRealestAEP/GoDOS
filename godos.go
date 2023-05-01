package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"sync"
	"time"
)

const (
	targetURL         = "http://localhost:8080" // Replace with the URL you want to test
	workers           = 10                      // Number of concurrent workers
	requestsPerWorker = 100                     // Number of requests each worker sends
)

var (
	mode     string
	targetIP string
)

func init() {
	flag.StringVar(&mode, "mode", "http", "Choose the test mode: http, tcp, or udp")
	flag.StringVar(&targetIP, "target", "localhost:8080", "Target IP address and port for TCP/UDP tests")
	flag.Parse()
}

func customRequest(url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "Custom-Load-Tester/1.0")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func tcpWorker(workerID int) {
	defer wg.Done()

	conn, err := net.Dial("tcp", targetIP)
	if err != nil {
		fmt.Printf("Worker %d: connection failed: %v\n", workerID, err)
		return
	}
	defer conn.Close()

	for j := 0; j < requestsPerWorker; j++ {
		_, err := conn.Write([]byte("hello"))
		if err != nil {
			fmt.Printf("Worker %d: write %d failed: %v\n", workerID, j, err)
			break
		}
		fmt.Printf("Worker %d: write %d completed\n", workerID, j)
		time.Sleep(time.Millisecond * 1)
	}
}

func udpWorker(workerID int) {
	defer wg.Done()

	conn, err := net.Dial("udp", targetIP)
	if err != nil {
		fmt.Printf("Worker %d: connection failed: %v\n", workerID, err)
		return
	}
	defer conn.Close()

	for j := 0; j < requestsPerWorker; j++ {
		_, err := conn.Write([]byte("hello"))
		if err != nil {
			fmt.Printf("Worker %d: write %d failed: %v\n", workerID, j, err)
			break
		}
		fmt.Printf("Worker %d: write %d completed\n", workerID, j)
		time.Sleep(time.Millisecond * 1)
	}
}

var wg sync.WaitGroup

func main() {
	for i := 0; i < workers; i++ {
		wg.Add(1)

		switch mode {
		case "http":
			go func(workerID int) {
				defer wg.Done()

				for j := 0; j < requestsPerWorker; j++ {
					resp, err := customRequest(targetURL)
					if err != nil {
						fmt.Printf("Worker %d: request %d failed: %v\n", workerID, j, err)
						continue
					}

					_, _ = ioutil.ReadAll(resp.Body)
					_ = resp.Body.Close()

					fmt.Printf("Worker %d: request %d completed with status %s\n", workerID, j, resp.Status)
					time.Sleep(time.Millisecond * 1)
				}
			}(i)

		case "tcp":
			go tcpWorker(i)
		case "udp":
			go udpWorker(i)
		default:
			fmt.Printf("Invalid mode: %s. Supported modes are http, tcp, and udp.\n", mode)
			return
		}
		
		}
		
		wg.Wait()
		fmt.Println("Load test completed")
		}
		
