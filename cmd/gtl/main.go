package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

const version = "0.1.0"

func main() {
	flag.Usage = func() {
		showHelp()
		os.Exit(1)
	}

	var (
		helpFlag     = flag.Bool("help", false, "Show help information")
		helpFlagH    = flag.Bool("h", false, "Show help information")
		versionFlag  = flag.Bool("version", false, "Show version information")
		versionFlagV = flag.Bool("v", false, "Show version information")
	)

	flag.Parse()

	if *helpFlag || *helpFlagH {
		showHelp()
		return
	}

	if *versionFlag || *versionFlagV {
		showVersion()
		return
	}

	var input string

	if flag.NArg() > 0 {
		input = strings.Join(flag.Args(), " ")
	} else {
		stat, _ := os.Stdin.Stat()
		if (stat.Mode() & os.ModeCharDevice) != 0 {
			showHelp()
			return
		}

		scanner := bufio.NewScanner(os.Stdin)
		var lines []string
		for scanner.Scan() {
			lines = append(lines, scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			fmt.Fprintf(os.Stderr, "Error reading from stdin: %v\n", err)
			os.Exit(1)
		}
		input = strings.Join(lines, " ")
	}

	if strings.TrimSpace(input) == "" {
		showHelp()
		return
	}

	result := toTitleCase(input)
	fmt.Println(result)
}

func showHelp() {
	fmt.Println("gtl - Go Title Linter")
	fmt.Println("Transforms text into properly capitalized titles according to the Chicago Manual of Style.")
	fmt.Println()
	showUsage()
	fmt.Println()
	fmt.Println("Options:")
	fmt.Println("  -h, --help     Show this help message")
	fmt.Println("  -v, --version  Show version information")
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("  gtl \"the quick brown fox\"")
	fmt.Println("  echo \"the quick brown fox\" | gtl")
}

func showUsage() {
	fmt.Println("Usage:")
	fmt.Println("  gtl [options] [text]")
	fmt.Println("  echo \"text\" | gtl [options]")
}

func showVersion() {
	fmt.Printf("gtl version %s\n", version)
}

func toTitleCase(text string) string {
	smallWords := map[string]bool{
		"a": true, "an": true, "and": true, "as": true, "at": true, "but": true,
		"by": true, "for": true, "if": true, "in": true, "nor": true, "of": true,
		"on": true, "or": true, "the": true, "to": true, "up": true, "yet": true,
	}

	words := strings.Fields(text)
	if len(words) == 0 {
		return text
	}

	result := make([]string, len(words))

	for i, word := range words {
		lowerWord := strings.ToLower(word)

		if i == 0 || i == len(words)-1 {
			result[i] = strings.Title(lowerWord)
		} else if smallWords[lowerWord] {
			result[i] = lowerWord
		} else {
			result[i] = strings.Title(lowerWord)
		}
	}

	return strings.Join(result, " ")
}
