package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/zhhink/jt/jsont"
)

var (
	inputFile   string
	outputItems string
	mode        string
	help        bool
)

func init() {
	flag.BoolVar(&help, "h", false, "this help")

	flag.StringVar(&inputFile, "i", "", "json file")
	flag.StringVar(&outputItems, "o", "", "items need to output")
	flag.StringVar(&mode, "m", "", "run mode")

	flag.Usage = usage
}

func main() {
	flag.Parse()

	if help {
		flag.Usage()
	}

	jsonT := jsont.JSONT{
		JSONFileName: inputFile,
	}

	jsonT.FilterItemsFromJSONFile(outputItems)
}

func usage() {
	fmt.Fprintf(os.Stderr, `jt version: jt/0.0.1
Usage: nginx [-hvVm] [-i inputfile] [-o outputItems]

Options:
`)
	flag.PrintDefaults()
}
