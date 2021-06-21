package main

import (
	"flag"
	"log"
	"os"
	"path"

	"github.com/fyrchik/sortnet/formatter"
	"github.com/fyrchik/sortnet/sequence"
)

var (
	size     = flag.Int("n", 0, "size of slice to sort")
	outFile  = flag.String("out", "sort.go", "name of file to output")
	name     = flag.String("name", "Sort", "name of the function")
	argType  = flag.String("argtype", "", "argument type (must be a slice type)")
	sortType = flag.String("sort", "batcher", "sort type (either 'batcher' or 'bitonic')")
)

func main() {
	flag.Parse()

	var ss sequence.SwapSequence
	switch *sortType {
	case "batcher":
		ss = sequence.Batcher(*size)
	case "bitonic":
		ss = sequence.Bitonic(*size)
	default:
		log.Fatalf("unknown sort type: %s\n", *sortType)
	}

	pkg := path.Base(path.Dir(*outFile))
	if pkg == "." {
		var err error
		pkg, err = os.Getwd()
		if err != nil {
			log.Fatalf("can't get current working directory: %v", err)
		}
	}
	pkg = path.Base(pkg)

	err := formatter.WriteFile(ss, &formatter.Options{
		OutFile:      *outFile,
		FuncName:     *name,
		ArgumentType: *argType,
		PackageName:  pkg,
		Size:         *size,
	})
	if err != nil {
		log.Fatalln(err)
	}
}
