package readwrite

import (
	"bufio"
	"errors"
	"io"
	"os"
	"path/filepath"
)

func Read(path string) ([]byte, error) {
	// open the file
	f, err := os.Open(filepath.Clean(path))
	if err != nil {
		return nil, err
	}
	defer f.Close()

	// create a bufio reader from the file
	r := bufio.NewReader(f)

	// start reading the file into the buffer
	buf := make([]byte, 1024)
	for {
		n, err := r.Read(buf)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}

			return nil, err
		}

		if n == 0 {
			break
		}
	}

	// return the buffer
	return buf, nil
}

func Write(path string, data []byte) (int, error) {
	// create the file
	f, err := os.Create(filepath.Clean(path))
	if err != nil {
		return 0, nil
	}
	defer f.Close()

	// create a new bufio writer for the file and write the data
	w := bufio.NewWriter(f)
	n, err := w.Write(data)
	if err != nil {
		return 0, err
	}
	w.Flush() // #nosec G104

	return n, nil
}
