package main

import (
	"bufio"
	_ "embed"
	"flag"
	"html/template"
	"io/fs"
	"log"
	"os"
)

//go:embed index.html.tmpl
var tmpl string

func main() {
	var title string
	var x string

	flag.StringVar(&title, "t", "LISTING", "Override default title")
	flag.StringVar(&x, "x", "", "Choose a different template")
	flag.Parse()

	if x != "" {
		bs, err := os.ReadFile(x)

		check(err)

		tmpl = string(bs)
	}

	var input []string
	var infos []fs.FileInfo

	if flag.NArg() == 0 {
		fi, err := os.Stdin.Stat()

		check(err)

		if fi.Size() == 0 {
			log.Fatal("listing: no input found")
		}

		scanner := bufio.NewScanner(os.Stdin)

		for scanner.Scan() {
			input = append(input, scanner.Text())
		}

		check(scanner.Err())
	} else {
		input = flag.Args()
	}

	for _, item := range input {
		fi, err := os.Stat(item)

		check(err)

		infos = append(infos, fi)
	}

	data := struct {
		List []fs.FileInfo
		Title string
	}{
		List: infos,
		Title: title,
	}

	t := template.Must(template.New("default").Parse(tmpl))

	check(t.Execute(os.Stdout, data))
}

func check(err error) {
	if err != nil {
		log.Fatal("listing: ", err)
	}
}
