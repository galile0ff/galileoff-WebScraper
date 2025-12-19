package cli

import (
	"flag"
	"fmt"
	"io"
	"strings"
	"time"
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

	lines := strings.Split(art.Data, "\n")
	colorCode := art.Color

	fmt.Fprint(w, colorCode)
	for _, line := range lines {
		fmt.Fprintln(w, line)
		time.Sleep(30 * time.Millisecond)
	}
	fmt.Fprint(w, "\033[0m")
}
