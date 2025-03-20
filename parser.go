package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

// CompileToString converts a markdown file to LaTeX
func CompileToString(inputFile string) (string, error) {
	// Open input file
	inFile, err := os.Open(inputFile)
	if err != nil {
		return "", err
	}
	defer inFile.Close()

	scanner := bufio.NewScanner(inFile)
	inMathBlock := false
	var mathBlock []string
	body := ""

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		// Check for math block start/end with $$ (handle both $$ and $$...$$ cases)
		if strings.HasPrefix(line, "$$") && strings.HasSuffix(line, "$$") {
			// Inline $$...$$ block
			mathContent := strings.TrimSpace(line[2 : len(line)-2])
			body += "\\[\n" + mathContent + "\n\\]\n"
			continue
		} else if strings.HasPrefix(line, "$$") {
			// Start of a multiline $$ block
			inMathBlock = true
			if len(line) > 2 {
				mathBlock = append(mathBlock, strings.TrimSpace(line[2:]))
			}
			continue
		} else if strings.HasSuffix(line, "$$") && inMathBlock {
			// End of a multiline $$ block
			inMathBlock = false
			mathBlock = append(mathBlock, strings.TrimSpace(line[:len(line)-2]))
			body += "\\[\n" + strings.Join(mathBlock, "\n") + "\n\\]\n"
			mathBlock = nil
			continue
		}

		// Collect lines inside a math block
		if inMathBlock {
			mathBlock = append(mathBlock, line)
			continue
		}

		// Normal lines and inline maths ($...$) are handled here
		body += convertToLatex(line) + "\n"
	}

	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("error reading input file: %w", err)
	}
	return body, nil
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

	// Other conversions can be added here

	return line
}

func convertInlineMath(line string) string {
	// Regex to match inline math $...$
	re := regexp.MustCompile(`\$(.*?)\$`)
	return re.ReplaceAllString(line, `\\( $1 \\)`)
}
