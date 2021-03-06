package main

import (
	"bytes"
	"compress/gzip"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

var (
	inputFilename  = flag.String("input", "/Users/kwochiu/project/golang/ebitenAwesome/source/runner.png", "input filename")
	outputFilename = flag.String("output", "/Users/kwochiu/project/golang/ebitenAwesome/source/runner.go", "output filename")
	packageName    = flag.String("package", "source", "package name")
	varName        = flag.String("var", "Runner_png", "variable name")
	compress       = flag.Bool("compress", false, "use gzip compression")
	buildtags      = flag.String("buildtags", "", "build tags")
)

func write(w io.Writer, r io.Reader) error {
	if *compress {
		compressed := &bytes.Buffer{}
		cw, err := gzip.NewWriterLevel(compressed, gzip.BestCompression)
		if err != nil {
			return err
		}
		if _, err := io.Copy(cw, r); err != nil {
			return err
		}
		cw.Close()
		r = compressed
	}

	bs, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}

	if _, err := fmt.Fprintln(w, "// Code generated by file2byteslice. DO NOT EDIT."); err != nil {
		return err
	}
	if _, err := fmt.Fprintln(w, "// (gofmt is fine after generating)"); err != nil {
		return err
	}
	if *buildtags != "" {
		if _, err := fmt.Fprintln(w, "\n// +build "+*buildtags); err != nil {
			return err
		}
	}
	if _, err := fmt.Fprintln(w, ""); err != nil {
		return err
	}
	if _, err := fmt.Fprintln(w, "package "+*packageName); err != nil {
		return err
	}
	if _, err := fmt.Fprintln(w, ""); err != nil {
		return err
	}
	if _, err := fmt.Fprintf(w, "var %s = []byte(%q)\n", *varName, string(bs)); err != nil {
		return err
	}
	return nil
}

func run() error {
	var out io.Writer
	if *outputFilename != "" {
		f, err := os.Create(*outputFilename)
		if err != nil {
			return err
		}
		defer f.Close()
		out = f
	} else {
		out = os.Stdout
	}

	var in io.Reader
	if *inputFilename != "" {
		f, err := os.Open(*inputFilename)
		if err != nil {
			return err
		}
		defer f.Close()
		in = f
	} else {
		in = os.Stdin
	}

	if err := write(out, in); err != nil {
		return err
	}

	return nil
}

func main() {
	flag.Parse()
	if err := run(); err != nil {
		panic(err)
	}
}
