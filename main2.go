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

func downloadWithTimeout(index int, filename string, wg *sync.WaitGroup) {
	defer wg.Done()

	ctx, cancel := context.WithTimeout(context.Background(), 4*7500*time.Millisecond)
	defer cancel()

	err := downloadFile(ctx, index, filename)
	if err != nil {
		log.Printf("Error downloading file %d: %v", index, err)
	}
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
		go downloadWithTimeout(i, filename, &wg)
	}

	wg.Wait()
	log.Println("All downloads completed")
}

