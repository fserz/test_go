package main

import "fmt"

// src/io/io.go

// src/io/io.go

type Reader interface {
	Read(p []byte) (n int, err error)
}

type Writer interface {
	Write(p []byte) (n int, err error)
}

type Closer interface {
	Close() error
}

// ReadWriter 是组合Reader接口和Writer接口形成的新接口类型
type ReadWriter interface {
	Reader
	Writer
}

// ReadCloser 是组合Reader接口和Closer接口形成的新接口类型
type ReadCloser interface {
	Reader
	Closer
}

// WriteCloser 是组合Writer接口和Closer接口形成的新接口类型
type WriteCloser interface {
	Writer
	Closer
}

type InterfaceTest struct{}

func (i *InterfaceTest) Read(p []byte) (n int, err error) {
	fmt.Println("read method...")
	return
}

func (i *InterfaceTest) Close() error {
	return fmt.Errorf("interface err")
}
