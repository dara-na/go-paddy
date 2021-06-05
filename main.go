package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/pflag"
)

var files = pflag.StringArrayP("files", "f", []string{}, "files to fix")

func main() {
	pflag.Parse()

	if len(*files) > 0 {
		for _, f := range *files {
			abs, err := filepath.Abs(f)
			if err != nil {
				log.Fatal(err)
			}
			source, err := os.ReadFile(abs)
			if err != nil {
				log.Fatal(err)
			}
			s, err := eachStructs(b2s(source), structTypeSorter)
			if err != nil {
				log.Fatal(err)
			}

			if err := os.WriteFile(abs, s2b(s), 0633); err != nil {
				log.Fatal(err)
			}
		}
	}
}
