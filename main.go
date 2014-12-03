package main

import (
	"fmt"
	"go/build"
	"io"
	"os"

	"github.com/docopt/docopt-go"
)

var usage = `
context-helpers

generate getter/setter functions for accessing a value via gorilla/context.

Usage:
	context-helpers <type>
	context-helpers -h | --help
	context-helpers --version

Options:
	-h --help     Show this screen.
	-v --version  Show version.
`

func main() {
	if err := run(os.Args, os.Stdin, os.Stdout); err != nil {
		fmt.Fprintln(os.Stdout, err)
		os.Exit(1)
	}
}

func run(args []string, _ io.Reader, out io.Writer) error {
	options, err := docopt.Parse(usage, args[1:], true, "0.0.1", false)
	if err != nil {
		return err
	}

	name := options["<type>"].(string)
	packageName, err := getPackageName()
	if err != nil {
		return err
	}

	helper := Helper{
		Name:        name,
		PackageName: packageName,
	}
	f, err := os.Create(helper.FileName())
	if err != nil {
		return err
	}
	defer f.Close()

	return helper.Render(f)
}

func getPackageName() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	p, err := build.ImportDir(dir, build.AllowBinary)
	if err != nil {
		return "", err
	}
	return p.Name, nil
}
