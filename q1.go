package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

const maxThreads = 4

type DataSet struct {
	SectionID int
	Numbers   []int
}

var wg sync.WaitGroup

func SectionRead(parts []string, sectionID int, ch chan<- DataSet) {
	defer wg.Done()

	seen := make(map[int]bool)
	var unique []int

	for _, part := range parts {
		tokens := strings.Fields(part)
		for _, tok := range tokens {
			if num, err := strconv.Atoi(tok); err == nil {
				if !seen[num] {
					seen[num] = true
					unique = append(unique, num)
				}
			}
		}
	}

	ch <- DataSet{SectionID: sectionID, Numbers: unique}
}

func main() {
	file, err := os.Open("number.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	var parts []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parts = append(parts, scanner.Text())
	}

	partTotal := len(parts)
	SectionSize := (partTotal) / maxThreads

	ch := make(chan DataSet, maxThreads)

	for i := 0; i < maxThreads; i++ {
		start := i * SectionSize
		end := start + SectionSize
		if end > partTotal {
			end = partTotal
		}
		wg.Add(1)
		go SectionRead(parts[start:end], i, ch)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	results := make([]DataSet, maxThreads)
	for data := range ch {
		results[data.SectionID] = data
	}

	seenGlobal := make(map[int]bool)
	fmt.Println("Unique numbers in file order:")
	for _, ds := range results {
		for _, num := range ds.Numbers {
			if !seenGlobal[num] {
				seenGlobal[num] = true
				fmt.Print(num, " ")
			}
		}
	}
	fmt.Println()
}
