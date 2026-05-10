package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/debkanchan/bark/internal/results"
	"github.com/debkanchan/bark/internal/scanner"
)

const (
	exitSuccess     = 0
	exitFound       = 1
	exitError       = 2
	defaultHookName = "pre-commit"
)

func main() {
	formatFlag := flag.String("format", "text", "Output format (text, json)")
	flag.StringVar(formatFlag, "f", "text", "Output format (text, json) - shorthand")

	pathFlag := flag.String("path", "", "Path to scan (optional, can be provided as argument)")
	flag.StringVar(pathFlag, "p", "", "Path to scan - shorthand")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Bark - Detect BARK comments in your code\n\n")
		fmt.Fprintf(os.Stderr, "Usage:\n")
		fmt.Fprintf(
			os.Stderr,
			"  bark [options] [path...]                      Scan for BARK comments\n",
		)
		fmt.Fprintf(
			os.Stderr,
			"  bark git-hook install [pre-commit|pre-push]   Install git hook (default: pre-commit)\n",
		)
		fmt.Fprintf(
			os.Stderr,
			"  bark git-hook uninstall [pre-commit|pre-push] Uninstall git hook (default: pre-commit)\n\n",
		)
		fmt.Fprintf(os.Stderr, "Arguments:\n")
		fmt.Fprintf(
			os.Stderr,
			"  path    Directories or files to scan (default: current directory)\n",
		)
		fmt.Fprintf(os.Stderr, "          Multiple paths can be provided to scan them all\n\n")
		fmt.Fprintf(os.Stderr, "Options:\n")
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\nExamples:\n")
		fmt.Fprintf(
			os.Stderr,
			"  bark                                          # Scan current directory\n",
		)
		fmt.Fprintf(
			os.Stderr,
			"  bark ./src                                    # Scan src directory\n",
		)
		fmt.Fprintf(
			os.Stderr,
			"  bark main.go utils.go lib/                    # Scan specific files and directories\n",
		)
		fmt.Fprintf(
			os.Stderr,
			"  bark -format json .                           # Scan current directory with JSON output\n",
		)
		fmt.Fprintf(
			os.Stderr,
			"  bark git-hook install                         # Install pre-commit hook (default)\n",
		)
		fmt.Fprintf(
			os.Stderr,
			"  bark git-hook install pre-push                # Install pre-push hook\n",
		)
		fmt.Fprintf(os.Stderr, "\nExit codes:\n")
		fmt.Fprintf(os.Stderr, "  0 - No BARK comments found\n")
		fmt.Fprintf(os.Stderr, "  1 - BARK comments found\n")
		fmt.Fprintf(os.Stderr, "  2 - Error occurred during scanning\n")
	}

	flag.Parse()

	if flag.NArg() > 0 && flag.Arg(0) == "git-hook" {
		if flag.NArg() < 2 {
			fmt.Fprintf(os.Stderr, "Error: git-hook requires a subcommand (install or uninstall)\n")
			fmt.Fprintf(os.Stderr, "Usage: bark git-hook install|uninstall [pre-commit|pre-push]\n")
			os.Exit(exitError)
		}

		subcommand := flag.Arg(1)
		hookName := defaultHookName
		if flag.NArg() >= 3 {
			hookName = flag.Arg(2)
		}
		if hookName != "pre-commit" && hookName != "pre-push" {
			fmt.Fprintf(os.Stderr, "Error: Unknown hook name '%s'\n", hookName)
			fmt.Fprintf(os.Stderr, "Valid hook names: pre-commit, pre-push\n")
			os.Exit(exitError)
		}

		switch subcommand {
		case "install":
			installGitHook(hookName)
		case "uninstall":
			uninstallGitHook(hookName)
		default:
			fmt.Fprintf(os.Stderr, "Error: Unknown git-hook subcommand '%s'\n", subcommand)
			fmt.Fprintf(os.Stderr, "Valid subcommands: install, uninstall\n")
			os.Exit(exitError)
		}
		return
	}

	var scanPaths []string
	if *pathFlag != "" {
		scanPaths = append(scanPaths, *pathFlag)
	}
	scanPaths = append(scanPaths, flag.Args()...)
	if len(scanPaths) == 0 {
		scanPaths = []string{"."}
	}

	var formatter results.Formatter
	switch *formatFlag {
	case "text":
		formatter = results.NewTextFormatter()
	case "json":
		formatter = results.NewJSONFormatter()
	default:
		fmt.Fprintf(os.Stderr, "Error: Invalid format '%s'. Use 'text' or 'json'.\n", *formatFlag)
		os.Exit(exitError)
	}

	for _, p := range scanPaths {
		_, err := os.Stat(p)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: Cannot access path '%s': %v\n", p, err)
			os.Exit(exitError)
		}
	}

	s := scanner.NewScanner()
	result := s.ScanPaths(scanPaths)

	output, err := formatter.Format(result)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error formatting output: %v\n", err)
		os.Exit(exitError)
	}

	fmt.Print(output)

	if len(result.GetErrors()) > 0 {
		os.Exit(exitError)
	}

	if result.HasFindings() {
		os.Exit(exitFound)
	}

	os.Exit(exitSuccess)
}
