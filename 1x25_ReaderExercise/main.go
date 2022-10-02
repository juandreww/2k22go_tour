package main

import "golang.org/x/tour/reader"

type MyReader struct{}
type MyReaderError bool

func (MyReaderError) Error() string {
	return "too short b capacity"
}

func (MyReader) Read(b []byte) (int, error) {
	if cap(b) < 1 {
		return 0, MyReaderError(true)
	}
	b[0] = 'A'
	return 1, nil
}

func main() {
	reader.Validate(MyReader{})
}
