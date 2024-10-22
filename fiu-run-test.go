package main

import (
	"crypto/rand"
	"fmt"
	"log"
	"os"
)

const (
	fileCount = 10
	fileSize  = 100 * 1024 // 100KB
)

func main() {
	// Write random bytes to files byte by byte
	for i := 0; i < fileCount; i++ {
		filename := fmt.Sprintf("file_%d.bin", i)
		if err := writeRandomBytesByteByByte(filename, fileSize); err != nil {
			log.Fatalf("Failed to write file %s: %v", filename, err)
		}
		fmt.Printf("Written file %s\n", filename)
	}

	// Read files byte by byte
	for i := 0; i < fileCount; i++ {
		filename := fmt.Sprintf("file_%d.bin", i)
		if err := readBytesByteByByte(filename, fileSize); err != nil {
			log.Fatalf("Failed to read file %s: %v", filename, err)
		}
		fmt.Printf("Read file %s\n", filename)
	}
}

// writeRandomBytesByteByByte writes random bytes to a file one byte at a time
func writeRandomBytesByteByByte(filename string, size int) error {
	// Open the file for writing
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	// Write one random byte at a time
	for i := 0; i < size; i++ {
		byteSlice := make([]byte, 1)
		if _, err := rand.Read(byteSlice); err != nil {
			return fmt.Errorf("failed to generate random byte: %w", err)
		}
		if _, err := file.Write(byteSlice); err != nil {
			return fmt.Errorf("failed to write byte to file: %w", err)
		}
	}

	return nil
}

// readBytesByteByByte reads the content of a file one byte at a time
func readBytesByteByByte(filename string, size int) error {
	// Open the file for reading
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	// Read one byte at a time
	for i := 0; i < size; i++ {
		byteSlice := make([]byte, 1)
		if _, err := file.Read(byteSlice); err != nil {
			return fmt.Errorf("failed to read byte from file: %w", err)
		}
	}

	return nil
}
