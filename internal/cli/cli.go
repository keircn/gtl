package cli

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/keircn/gtl/internal/titlecase"
	"github.com/keircn/gtl/pkg/version"
)

func Run() {
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

	result := titlecase.ToTitleCase(input)
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
	fmt.Println(version.BuildVersion())
	fmt.Println()
	fmt.Println("Author: keircn")
	fmt.Println("Report bugs at: https://github.com/keircn/gtl/issues")
}
