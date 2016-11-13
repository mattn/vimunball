package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	scan := bufio.NewScanner(os.Stdin)
	for scan.Scan() {
		if scan.Text() == "finish" {
			break
		}
	}
	if err := scan.Err(); err != nil {
		log.Fatal(err)
	}

	for scan.Scan() {
		line := scan.Text()
		if !strings.HasSuffix(line, "[[[1") {
			break
		}
		fname := line[:len(line)-5]
		os.MkdirAll(filepath.Dir(fname), 0644)
		scan.Scan()
		f, err := os.Create(fname)
		if err != nil {
			log.Fatal(err)
		}
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
