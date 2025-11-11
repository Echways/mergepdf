package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/pdfcpu/pdfcpu/pkg/api"
)

const maxInputs = 10

func printUsage() {
	fmt.Println("Использование:")
	fmt.Println("  pdfmerge output.pdf input1.pdf [input2.pdf ... input10.pdf]")
	fmt.Println()
	fmt.Println("Пример:")
	fmt.Println("  pdfmerge merged.pdf file1.pdf file2.pdf file3.pdf")
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Ошибка: нужно указать файл-результат и хотя бы один входной PDF.")
		printUsage()
		os.Exit(1)
	}

	output := os.Args[1]
	inputs := os.Args[2:]

	if len(inputs) > maxInputs {
		fmt.Printf("Ошибка: максимум %d входных файлов, вы передали %d.\n", maxInputs, len(inputs))
		os.Exit(1)
	}

	for _, in := range inputs {
		if _, err := os.Stat(in); err != nil {
			fmt.Printf("Ошибка: входной файл '%s' не найден или недоступен: %v\n", in, err)
			os.Exit(1)
		}
	}

	output = normalizePath(output)

	if err := mergePDFs(output, inputs); err != nil {
		fmt.Printf("Ошибка при объединении PDF: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Успешно создан файл: %s\n", output)
}

func normalizePath(path string) string {
	if !filepath.IsAbs(path) {
		abs, err := filepath.Abs(path)
		if err == nil {
			return abs
		}
	}
	return path
}

func mergePDFs(output string, inputs []string) error {
	err := api.MergeCreateFile(inputs, output, false, nil)
	return err
}
