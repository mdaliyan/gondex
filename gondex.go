package main

import (
	`os`

	`github.com/mdaliyan/gondex/internal/crawler`
	`github.com/mdaliyan/gondex/internal/index`
)

func main() {
	path := os.Args[1]
	if len(os.Args) > 2 {
		index.Prefix = os.Args[2]
	}
	crawler.Generate(path)
}
