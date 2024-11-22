package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	var inputPath string
	var outputPath string
	var latexTemplateFile = "template.tex"

	flag.StringVar(&inputPath, "src", "", "The markdown file to compile")
	flag.StringVar(&outputPath, "out", "", "The latex file name")
	flag.Parse()

	if inputPath == "" {
		log.Fatal("No file provided")
	}

	if outputPath == "" {
		base_name := strings.Split(inputPath, ".")[0]
		outputPath = base_name + ".tex"
		fmt.Println("No output file provided, using markdown name:", outputPath)
	}

	basePath, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	inputPath = filepath.Join(basePath, inputPath)

	// Read the LaTeX template
	templateContent, err := os.ReadFile(latexTemplateFile)
	if err != nil {
		log.Fatal("Error reading LaTeX template file:", err)
	}

	if _, err := os.Stat(outputPath); err == nil {
		log.Fatal("output file already exists", err)
	}

	// Create the output file to write to
	outputFile, err := os.Create(outputPath)
	if err != nil {
		log.Fatal("Error creating output file:", err)
	}
	defer outputFile.Close()

	writer := bufio.NewWriter(outputFile)
	defer writer.Flush()

	// Write the template up to \begin{document}
	writer.WriteString(string(templateContent))

	// Read and write the markdown file to LaTex
	latexBody, err := CompileToString(inputPath)
	if err != nil {
		// Input file didn't open correctly
		log.Fatal("Error opening input file:", err)
	}
	writer.WriteString(latexBody + "\n")

	// End the document
	writer.WriteString("\\end{document}\n")
}
