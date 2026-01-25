package interface_example

import "fmt"

type Reader interface {
	Read(p []byte) (n int, err error)
}

type Writer interface {
	Write(p []byte) (n int, err error)
}

type ReadWriter interface {
	Reader
	Writer
}

type File struct {
	Name string
}

func (f *File) Read(p []byte) (n int, err error) {
	fmt.Printf("Reading %d bytes from file: %s\n", len(p), f.Name)
	return len(p), nil
}

func (f *File) Write(p []byte) (n int, err error) {
	fmt.Printf("Writing %d bytes to file: %s\n", len(p), f.Name)
	return len(p), nil
}

func Process(data []byte, rw ReadWriter) {
	rw.Write(data)
	rw.Read(data)
}
