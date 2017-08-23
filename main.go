package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func fatalIf(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	var err error
	var outdir, file string
	flag.StringVar(&outdir, "o", ".", "output directory")
	flag.StringVar(&file, "f", "", "vim ball file")
	flag.Parse()

	fatalIf(os.MkdirAll(outdir, 0755))

	var in *os.File
	if file == "" {
		in = os.Stdin
	} else {
		in, err = os.Open(file)
		fatalIf(err)
		defer in.Close()
	}

	scan := bufio.NewScanner(in)
	for scan.Scan() {
		if scan.Text() == "finish" {
			break
		}
	}
	fatalIf(scan.Err())

	for scan.Scan() {
		line := scan.Text()
		if !strings.HasSuffix(line, "[[[1") {
			break
		}
		fname := line[:len(line)-5]
		log.Println(fname)
		os.MkdirAll(filepath.Join(outdir, filepath.Dir(fname)), 0644)
		scan.Scan()
		f, err := os.Create(filepath.Join(outdir, fname))
		fatalIf(err)
		nl, _ := strconv.Atoi(scan.Text())
		for i := 0; i < nl; i++ {
			if !scan.Scan() {
				break
			}
			fmt.Fprintf(f, "%s\n", scan.Text())
		}
		f.Close()
	}
}
