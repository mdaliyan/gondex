package index

import "os"

type Node struct {
	Level int
	Path string
	Files []os.FileInfo
}
