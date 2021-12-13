package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
	"sync"
)

var (
	matches []string
	lock    sync.Mutex
	wg      sync.WaitGroup
)

func fileSearch(root, filename string) {
	fmt.Println("Searching in", root)
	files, _ := ioutil.ReadDir(root)

	for _, file := range files {
		fullPath := filepath.Join(root, file.Name())
		if strings.Contains(file.Name(), filename) {
			// locks to avoid race condition to write on matches array
			lock.Lock()
			matches = append(matches, fullPath)
			lock.Unlock()
		}
		if file.IsDir() {
			wg.Add(1)
			go fileSearch(fullPath, filename)
		}
	}

	wg.Done()
}

func main() {
	wg.Add(1)
	go fileSearch("/Users", "README.md")
	wg.Wait()

	for _, file := range matches {
		fmt.Println("Matched", file)
	}
}
