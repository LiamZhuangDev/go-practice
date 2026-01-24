package panic

import (
	"fmt"
	"os"
)

func ProcessFile(path string) (err error) {
	// Defer a function to handle any panic that occurs
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("error processing file %s: %v", path, r)
		}
	}()

	file, err := os.Open(path)
	if err != nil {
		return err
	}

	// Ensure the file is closed after processing
	defer func() {
		fmt.Printf("Closing file: %s\n", path)
		file.Close()
	}()

	parseFile(file)

	return nil
}

func parseFile(f *os.File) {
	// Simulate a panic during parsing
	panic("simulated parsing error")
}
