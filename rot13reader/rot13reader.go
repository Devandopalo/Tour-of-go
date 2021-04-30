package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (reader *rot13Reader) Read(buffer []byte) (int, error) {
	lenth, err := reader.r.Read(buffer)

	for i := 0; i < lenth; i++ {
		if buffer[i] == ' ' {
			continue
		} else if (buffer[i] >= 'a' && buffer[i] < 'n') || (buffer[i] >= 'A' && buffer[i] < 'N') {
			buffer[i] += 13
		} else if (buffer[i] > 'm' && buffer[i] <= 'z') || (buffer[i] > 'M' || buffer[i] <= 'Z') {
			buffer[i] -= 13
		}
	}
	return lenth, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
