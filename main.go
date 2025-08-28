package main

import (
	"fmt"
	"go-demo/mask"
	"os"
)

const defaultOut = "output.txt"

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Вы не ввели пути к файлам\n")
		os.Exit(2)
	}

	inputPath := os.Args[1]
	outputPath := defaultOut

	fmt.Printf("Команды %s", os.Args)

	if len(os.Args) >= 3 {
		outputPath = os.Args[2]
	}

	prod := mask.NewFileReader(inputPath)
	pres := mask.NewFilePresenter(outputPath)

	svc := mask.NewService(prod, pres)
	
	if err := svc.Run(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}

	fmt.Printf("OK: processed '%s' -> '%s'\n", inputPath, outputPath)
}