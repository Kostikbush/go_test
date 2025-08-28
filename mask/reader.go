package mask

import (
	"bufio"
	"os"
)

type FileReader struct {
	inPath string
}

func NewFileReader(inPath string) *FileReader {
	return &FileReader{inPath: inPath}
}

func (p *FileReader) produce() ([]string, error) {
	f, err := os.Open(p.inPath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	lines := make([]string, 0, 128)
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
	if err := sc.Err(); err != nil {
		return nil, err
	}
	return lines, nil
}