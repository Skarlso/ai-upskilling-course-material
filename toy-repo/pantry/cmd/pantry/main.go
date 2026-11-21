// Command pantry keeps track of what is in the kitchen.
package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"example.com/pantry/internal/item"
	"example.com/pantry/internal/report"
	"example.com/pantry/internal/store"
)

const usage = `pantry keeps track of what is in the kitchen.

Usage:
  pantry [--file path] add <amount> <name>      pantry add 500g flour
  pantry [--file path] remove <name> <amount>   pantry remove flour 200
  pantry [--file path] list

The pantry is stored in pantry.json in the current directory. Use --file or
the PANTRY_FILE environment variable to keep it somewhere else.
`

func main() {
	os.Exit(run(os.Args[1:], os.Stdout, os.Stderr))
}

// run is the whole program. It returns the exit code: 0 for success, 1 when
// the command failed, and 2 when the command line was wrong.
func run(args []string, stdout, stderr io.Writer) int {
	flags := flag.NewFlagSet("pantry", flag.ContinueOnError)
	flags.SetOutput(stderr)
	flags.Usage = func() { fmt.Fprint(stderr, usage) }
	file := flags.String("file", defaultFile(), "where the pantry is stored")
	if err := flags.Parse(args); err != nil {
		if errors.Is(err, flag.ErrHelp) {
			return 0
		}
		return 2
	}

	rest := flags.Args()
	if len(rest) == 0 {
		flags.Usage()
		return 2
	}

	s, err := store.Load(*file)
	if err != nil {
		return fail(stderr, err)
	}

	switch rest[0] {
	case "add":
		it, err := item.Parse(strings.Join(rest[1:], " "))
		if err != nil {
			return fail(stderr, err)
		}
		if err := s.Add(it); err != nil {
			return fail(stderr, err)
		}
		if err := s.Save(*file); err != nil {
			return fail(stderr, err)
		}
		have, _ := s.Get(it.Name)
		fmt.Fprintf(stdout, "Added. You now have %s.\n", have)
		return 0

	case "remove":
		// The name may be several words, so the amount is the last argument.
		if len(rest) < 3 {
			flags.Usage()
			return 2
		}
		name := strings.Join(rest[1:len(rest)-1], " ")
		qty, err := strconv.Atoi(rest[len(rest)-1])
		if err != nil {
			return fail(stderr, fmt.Errorf("amount must be a whole number, got %q", rest[len(rest)-1]))
		}
		if err := s.Remove(name, qty); err != nil {
			return fail(stderr, err)
		}
		if err := s.Save(*file); err != nil {
			return fail(stderr, err)
		}
		fmt.Fprintln(stdout, "Removed.")
		return 0

	case "list":
		items := s.List()
		fmt.Fprint(stdout, report.Table(items))
		fmt.Fprintln(stdout, report.Summary(items))
		return 0
	}

	fmt.Fprintf(stderr, "pantry: unknown command %q\n\n", rest[0])
	flags.Usage()
	return 2
}

func fail(stderr io.Writer, err error) int {
	fmt.Fprintln(stderr, "pantry:", err)
	return 1
}

func defaultFile() string {
	if path := os.Getenv("PANTRY_FILE"); path != "" {
		return path
	}
	return "pantry.json"
}
