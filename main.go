package main

import (
	"bufio"
	_ "embed"
	"flag"
	"io/fs"
	"log"
	"os"
	"text/template"
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

		// https://stackoverflow.com/questions/22744443
		if (fi.Mode() & os.ModeCharDevice) == 0 {
			scanner := bufio.NewScanner(os.Stdin)

			for scanner.Scan() {
				input = append(input, scanner.Text())
			}

			check(scanner.Err())
		} else {
			log.Fatal("listing: no input found")
		}
	} else {
		input = flag.Args()
	}

	for _, item := range input {
		fi, err := os.Stat(item)

		check(err)

		infos = append(infos, fi)
	}

	data := struct {
		List  []fs.FileInfo
		Title string
	}{
		List:  infos,
		Title: title,
	}

	funcMap := map[string]interface{}{"list": list}
	t := template.Must(template.New("default").Funcs(template.FuncMap(funcMap)).Parse(tmpl))

	check(t.Execute(os.Stdout, data))
}

func check(err error) {
	if err != nil {
		log.Fatal("listing: ", err)
	}
}

func list(args ...interface{}) []interface{} {
	return args
}
