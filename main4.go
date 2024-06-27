package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

const (
	maxRetries      = 3
	downloadTimeout = 10 * time.Second
)

func loadFileNameListFromLocation(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

func downloadFile(ctx context.Context, index int, filename string) error {
	url := "https://lhit.ro/images/" + filename
	fmt.Printf("About to download file %d named %s\n", index, url)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return fmt.Errorf("failed to create request for file %d: %v", index, err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to download file %d: %v", index, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to download file %d: received status code %d", index, resp.StatusCode)
	}

	filepath := "./downloads/" + filename
	out, err := os.Create(filepath)
	if err != nil {
		return fmt.Errorf("failed to create file %d: %v", index, err)
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return fmt.Errorf("failed to save file %d: %v", index, err)
	}

	return nil
}

func main() {
	var wg sync.WaitGroup

	filenames, err := loadFileNameListFromLocation("./file_list.txt")
	if err != nil {
		log.Fatalf("Cannot read file list: %v", err)
	}

	// Ensure the downloads directory exists
	if err := os.MkdirAll("./downloads", os.ModePerm); err != nil {
		log.Fatalf("Failed to create downloads directory: %v", err)
	}

	for i, filename := range filenames {
		wg.Add(1)

		go func(i int, filename string) {
			defer wg.Done()

			var err error
			for attempt := 1; attempt <= maxRetries; attempt++ {
				// Create a new context with a 30-second timeout for each download attempt
				ctx, cancel := context.WithTimeout(context.Background(), downloadTimeout)

				err = downloadFile(ctx, i, filename)
				cancel() // Cancel the context to free resources

				if err == nil {
					break
				}

				log.Printf("Attempt %d: Error downloading file %d: %v", attempt, i, err)
				if attempt < maxRetries {
					time.Sleep(2 * time.Second) // Wait for 2 seconds before retrying
				}
			}

			if err != nil {
				log.Printf("Failed to download file %d after %d attempts: %v", i, maxRetries, err)
			}
		}(i, filename)
	}

	wg.Wait()
	log.Println("All downloads completed")
}
