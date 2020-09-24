package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"sync"
)

var fname string
var nS string
var n int
var wg sync.WaitGroup

func worker(i int, queue <-chan string) {
	defer wg.Done()
	for cmd := range queue {
		fmt.Printf("#%d %s\n", i, cmd)
		c := exec.Command("sh", "-c", cmd)
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr

		s, _ := os.Getwd()
		c.Dir = s
		err := c.Run()
		if err != nil {
			log.Printf("Error: %#v\n", err)
		}
	}
}

func main() {
	var err error
	flag.StringVar(&fname, "file", "", "file name")
	flag.StringVar(&nS, "n", "1", "num of processes")
	flag.Parse()
	n, err = strconv.Atoi(nS)
	if err != nil {
		panic(err)
	}

	queue := make(chan string)
	bytes, err := ioutil.ReadFile(fname)

	content := string(bytes)
	bytes = nil
	if err != nil {
		panic(err)
	}

	for i := 0; i < n; i++ {
		wg.Add(1)
		go worker(i+1, queue)
	}
	lines := strings.Split(content, "\n")
	for _, line := range lines {
		line := strings.TrimSpace(line)
		if strings.HasPrefix(line, "#") || line == "" {
			continue
		}
		queue <- line
	}
	close(queue)
	wg.Wait()
}
