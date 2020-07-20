package index

import (
	"bytes"
	"log"
	"os"
	"path"
)

var Prefix = ""

var logError = log.New(os.Stderr, "", log.Ltime)

var newLine = []byte("\n")
var sep = []byte(string(os.PathSeparator))

func Generate(n *Node) {

	b := bytes.Buffer{}
	// sort.Slice(files, func(i, j int) bool { return files[i].Name() < files[j].Name() })
	// sort.Slice(files, func(i, j int) bool { return files[i].Name() < files[j].Name() })

	p := path.Join(Prefix, n.Path)

	for _, d := range n.Files {
		b.Write([]byte(path.Join(p, d.Name())))
		if d.IsDir() {
			b.Write(sep)
		}
		b.Write(newLine)
	}

	writeToFile(n.Path, &b)
}

func writeToFile(p string, b *bytes.Buffer) {
	index := path.Join(p, "index.html")
	f, err := os.OpenFile(index, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		logError.Println(err)
	}
	_, err = b.WriteTo(f)
	if err != nil {
		logError.Println(err)
	}
	f.Close()
}
