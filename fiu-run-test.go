package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
	"slices"
)

const (
	fileCountCst = 1
	//fileSizeCst  = 100 * 1024 // 100KB
	fileSizeCst = 1000 * 1024 // 1000KB = 1MB
	//fileSizeCst  = 10* 1000 * 1024 // 10,000KB = 10MB

	mpCst = 1

	debugLevelCst = 11
)

var (
	// Passed by "go build -ldflags" for the show version
	commit  string
	date    string
	version string

	debugLevel int
)

func main() {

	fc := flag.Int("fc", fileCountCst, "file count")
	fs := flag.Int("fs", fileSizeCst, "file size")

	// Maximum number of CPUs that can be executing simultaneously
	// https://golang.org/pkg/runtime/#GOMAXPROCS -> zero (0) means default
	goMaxProcs := flag.Int("goMaxProcs", mpCst, "goMaxProcs = https://golang.org/pkg/runtime/#GOMAXPROCS")

	v := flag.Bool("v", false, "show version")

	d := flag.Int("d", debugLevelCst, "debug level")

	flag.Parse()

	// Print version information passed in via ldflags in the Makefile
	if *v {
		log.Printf("xtcp commit:%s\tdate(UTC):%s\tversion:%s", commit, date, version)
		os.Exit(0)
	}

	fileCount := *fc
	fileSize := *fs

	debugLevel = *d

	if runtime.NumCPU() > *goMaxProcs {
		mp := runtime.GOMAXPROCS(*goMaxProcs)
		if debugLevel > 10 {
			log.Printf("Main runtime.GOMAXPROCS now:%d was:%d\n", *goMaxProcs, mp)
		}
	}

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
		if err := readBytesByteByByte(filename, fileSize, debugLevel); err != nil {
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
func readBytesByteByByte(filename string, size int, debugLevel int) error {
	// Open the file for reading
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	bigSlice := make([]byte, size)

	// Read one byte at a time
	for i := 0; i < size; i++ {
		byteSlice := make([]byte, 1)
		var n int
		var err error
		if n, err = file.Read(byteSlice); err != nil {
			return fmt.Errorf("failed to read byte from file: %w", err)
		}
		if debugLevel > 100 {
			fmt.Printf("n:%d:%x\n", n, byteSlice)
		}

		if debugLevel > 10 {
			bigSlice = slices.Concat(bigSlice, byteSlice)
		}
	}

	fmt.Printf("bigSlice:%x\n", bigSlice)

	return nil
}
