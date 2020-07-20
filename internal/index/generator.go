package index

import (
	"bytes"
	"log"
	"os"
	"path"
)

var logError = log.New(os.Stderr, "", log.Ltime)

var newLine = []byte("\n")
var sep = []byte(string(os.PathSeparator))

func Generate(n *Node) {

	b := bytes.Buffer{}
	// sort.Slice(files, func(i, j int) bool { return files[i].Name() < files[j].Name() })
	// sort.Slice(files, func(i, j int) bool { return files[i].Name() < files[j].Name() })

	p := []byte(n.Path)

	for _, d := range n.Files {
		b.Write(p)
		b.Write(sep)
		b.Write([]byte(d.Name()))
		if d.IsDir() {
			b.Write(sep)
		}
		b.Write(newLine)
	}

	writeToFile(n.Path, &b)
}

func writeToFile(p string, b *bytes.Buffer) {
	f, err := os.OpenFile(path.Join(p,"index.html"), os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		logError.Println(err)
	}
	_, err = b.WriteTo(f)
	if err != nil {
		logError.Println(err)
	}
	f.Close()
}
