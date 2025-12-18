package cli

import (
	"flag"
	"fmt"
	"io"
)

type Options struct {
	ArtName string
}

func Parse() *Options {
	if flag.Lookup("art") == nil {
		flag.String("art", "", "ASCII art to show (duck, cat, skull)")
	}

	if !flag.Parsed() {
		flag.Parse()
	}

	artName := ""
	f := flag.Lookup("art")
	if f != nil {
		artName = f.Value.String()
	}

	return &Options{
		ArtName: artName,
	}
}

func Args() []string {
	if !flag.Parsed() {
		flag.Parse()
	}
	return flag.Args()
}

func PrintASCII(w io.Writer, opts *Options) {
	var art Art
	var err error

	if opts.ArtName != "" {
		art, err = GetArt(opts.ArtName)
		if err != nil {
			art = RandomArt()
		}
	} else {
		art = RandomArt()
	}

	fmt.Fprintln(w, art.Color+art.Data+"\033[0m")
}
