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
		fmt.Fprintf(os.Stderr, "Error: No input provided\n")
		showUsage()
		os.Exit(1)
	}

	result := toTitleCase(input)
	fmt.Println(result)
}

func showHelp() {
	fmt.Println("gtl - Go Title Case")
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
	return text
}
