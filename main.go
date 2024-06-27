package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync" // for waitgroup and mutex
	"time"

	//"time" "context" "math/rand" "net/http"
	"os"
)

// var (
//   ctx context.Context
//   wg sync.WaitGroup
//   mu sync.Mutex
// )

func loadFileNameListFromLocation(url string) ([]string, error) {
	loader, err := os.Open("./file_list.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer loader.Close()

	// Create a new scanner
	scanner := bufio.NewScanner(loader)

	// Create a slice to hold the lines
	var lines []string

	// Read each line and append to the slice
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// Check for scanning errors
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// // Print the lines
	// for i, line := range lines {
	// 	fmt.Printf("Line %d: %s\n", i+1, line)
	// }

	return lines, nil
}

func downloadFile(ctx context.Context, index int, filename string, wg *sync.WaitGroup) error {
	defer wg.Done()
	fmt.Printf("About to download file %d named https://lhit.ro/images/%s\n", index, filename)
	// Get the data
	url := "https://lhit.ro/images/" + filename
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	filepath := "./downloads/" + filename
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}

func main() {

	var wg sync.WaitGroup

	bunch, err := loadFileNameListFromLocation("https://lhit.ro/images/")
	if err != nil {
		fmt.Errorf("Cannot read folder %w", err)
	}

	for i := 0; i < len(bunch); i++ {
		wg.Add(1)
		ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
		defer cancel()
		go func(i int, s string) {
			downloadFile(ctx, i, bunch[i], &wg)
		}(i, bunch[i])
		// fmt.Printf("File %d named %s\n", i, bunch[i])
	}
	wg.Wait()
	log.Println("All downloads completed")

	fmt.Println("All downloads completed")

	// for _, index := range bunch {
	// 	fmt.Println(index)
	// }

}
