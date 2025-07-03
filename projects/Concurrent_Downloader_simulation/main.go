package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func downloadFile(file string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Downloading: ", file)
	time.Sleep(time.Duration(rand.Intn(3)+1) * time.Second)
	fmt.Println("Download Complete: ", file)
}

func main() {
	rand.Seed(time.Now().UnixNano())

	files := []string{"file1.txt", "file2.txt", "file4.txt"}

	var wg sync.WaitGroup

	for _, f := range files {
		wg.Add(1)
		go downloadFile(f, &wg)
	}

	wg.Wait()
	fmt.Println("Congrats.....All downloads finished.")

}
