package main

import (
	"errors"
	"fmt"
	"go/build"
	"io"
	"os"
)

func main() {
	if err := run(os.Args, os.Stdin, os.Stdout); err != nil {
		fmt.Fprintln(os.Stdout, err)
		os.Exit(1)
	}
}

func run(args []string, _ io.Reader, out io.Writer) error {
	if len(args) < 2 {
		return errors.New("Wrong number of arguments")
	}

	name := args[1]
	packageName, err := getPackageName()
	if err != nil {
		return err
	}

	helper := Helper{
		Name:        name,
		PackageName: packageName,
	}
	return helper.Render(out)
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
