package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	var inputFile string
	var outputFile string

	flag.StringVar(&inputFile, "input", "", "The markdown file to compile")
	flag.Parse()

	if inputFile == "" {
		fmt.Println("No file provided. Use the -input flag")
		os.Exit(1)
	}

	flag.StringVar(&inputFile, "output", "", "The latex file name")
	flag.Parse()

	if outputFile == "" {
		fmt.Println("No output file provided, using markdown name")
		base_name := strings.Split(inputFile, ".")[0]
		outputFile = base_name + ".tex"
	}

	// Open input file
	inFile, err := os.Open(inputFile)
	if err != nil {
		fmt.Println("Error opening input file:", err)
		return
	}
	defer inFile.Close()

	// Open output file
	outFile, err := os.Create(outputFile)
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer outFile.Close()

	writer := bufio.NewWriter(outFile)
	defer writer.Flush()

	// Write LaTeX document class and start
	// TODO - Make this a template
	writer.WriteString("\\documentclass{book}\n\\begin{document}\n")

	// Read input file line by line
	scanner := bufio.NewScanner(inFile)
	for scanner.Scan() {
		line := scanner.Text()
		latexLine := convertToLatex(line)
		writer.WriteString(latexLine + "\n")
	}

	// End the document
	writer.WriteString("\\end{document}\n")

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input file:", err)
	}
}

func convertToLatex(line string) string {
	// Header conversion
	if strings.HasPrefix(line, "# ") {
		return "\\chapter{" + strings.TrimPrefix(line, "# ") + "}"
	} else if strings.HasPrefix(line, "## ") {
		return "\\section{" + strings.TrimPrefix(line, "## ") + "}"
	} else if strings.HasPrefix(line, "### ") {
		return "\\subsection{" + strings.TrimPrefix(line, "### ") + "}"
	}

	// Inline math conversion
	line = convertInlineMath(line)

	// Block math conversion
	line = convertBlockMath(line)

	// Other conversions can be added here

	return line
}

func convertInlineMath(line string) string {
	re := regexp.MustCompile(`\$\$(.*?)\$\$`)
	return re.ReplaceAllString(line, `\[ $1 \]`)
}

func convertBlockMath(line string) string {
	re := regexp.MustCompile(`\$(.*?)\$`)
	return re.ReplaceAllString(line, `\( $1 \)`)
}
