package mask

import (
	"bufio"
	"os"
)

type FilePresenter struct {
	outPath string
}

func NewFilePresenter(outPath string) *FilePresenter {
	return &FilePresenter{outPath: outPath}
}

func (p *FilePresenter) present(lines []string) error {
	f, err := os.Create(p.outPath)
	if err != nil {
		return err
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	for i := range lines {
		if _, err := w.WriteString(lines[i] + "\n"); err != nil {
			return err
		}
	}
	return w.Flush()
}