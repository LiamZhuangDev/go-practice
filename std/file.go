package std

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

func FileRWTest() {
	// setup
	_, path, _, ok := runtime.Caller(0)
	if !ok {
		panic("Cannot retrieve current path")
	}
	dir := filepath.Dir(path)
	targetPath := filepath.Join(dir, "go-std-lib.txt")

	// Create and write content to a file
	// - Prefix '0o' stands for octal (base-8)
	// - '644' means permissions, which are split into three digits:
	// Digit	Who	    Meaning
	// 6	    Owner	read + write
	// 4	    Group	read
	// 4	    Others	read

	// Each digit is a sum of:
	// 4 → read (r)
	// 2 → write (w)
	// 1 → execute (x)

	// So:
	// 6 = 4 + 2 → rw-
	// 4 = 4 → r--
	// Final result: rw-r--r--
	if err := os.WriteFile(targetPath, []byte("Hello World!"), 0o644); err != nil {
		panic(err)
	}

	// Read content from the file
	content, err := os.ReadFile(targetPath)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Reads %s: %s\n", targetPath, content)

	// Create another file
	targetPath2 := filepath.Join(dir, "go-std-lib-2.txt")
	targetFile2, err := os.Create(targetPath2)
	if err != nil {
		panic(err)
	}
	defer targetFile2.Close()

	// Write by bufio.Writer
	writer := bufio.NewWriter(targetFile2) // NewWriter accepts io.Writer, and os.File implement this interface
	lines := []string{
		"line1: Hello World!",
		"line2: From bufio.Writer",
	}
	for _, line := range lines {
		if _, err := writer.WriteString(line + "\n"); err != nil {
			panic(err)
		}
	}
	if err := writer.Flush(); err != nil {
		panic(err)
	}
	fmt.Printf("Writes to %s: %s\n", targetPath2, lines)

	// Open and Read content by bufio.Scanner
	file, err := os.Open(targetPath2)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var scanWords = false
	if scanWords {
		scanner.Split(bufio.ScanWords) // default is bufio.ScanLines
	}
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
