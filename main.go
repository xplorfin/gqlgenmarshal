package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/xplorfin/gqlgenmarshal/generator"
)

var (
	typeNames        = flag.String("type", "", "comma-separated list of type names; must be set")
	output           = flag.String("output", "", "output file name; default srcdir/<type>_gqlgenmarshal.go")
	trimprefix       = flag.String("trimprefix", "", "trim the `prefix` from the generated constant names")
	lcasestringcheck = flag.Bool("checkstringlowercase", false, "If <type> implements fmt.Stringer, check the unmarshaled value against lowercase versions "+
		"of the type rather than the raw result of <type>.String()")
)

func Usage() {
	fmt.Fprintf(os.Stderr, "Usage of gqlgenmarshal:\n")
	fmt.Fprintf(os.Stderr, "\tgqlgenmarshal [flags] -type T [directory]\n")
	fmt.Fprintf(os.Stderr, "\tgqlgenmarshal [flags] -type T files... # Must be a single package\n")
	fmt.Fprintf(os.Stderr, "For more information, see:\n")
	fmt.Fprintf(os.Stderr, "Flags:\n")
	flag.PrintDefaults()
}

func main() {
	log.SetFlags(0)
	log.SetPrefix("gqlgenmarshal: ")
	flag.Usage = Usage
	flag.Parse()

	if len(*typeNames) == 0 {
		flag.Usage()
		os.Exit(2)
	}

	types := strings.Split(*typeNames, ",")

	// We accept either one directory or a list of files. Which do we have?
	args := flag.Args()
	if len(args) == 0 {
		// Default: process whole package in current directory.
		args = []string{"."}
	}

	// Parse the package once.
	var dir string

	if len(args) == 1 && isDirectory(args[0]) {
		dir = args[0]
	} else {
		dir = filepath.Dir(args[0])
	}

	fmt.Fprintf(os.Stderr, dir)

	gen := &generator.Generator{
		TrimPrefix:       *trimprefix,
		LCaseStringCheck: *lcasestringcheck,
	}

	gen.Generate(types, args, dir, output)
}

// isDirectory reports whether the named file is a directory.
func isDirectory(name string) bool {
	info, err := os.Stat(name)
	if err != nil {
		log.Fatal(err)
	}
	return info.IsDir()
}
