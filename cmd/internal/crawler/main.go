package crawler

import (
	"log"
	"os"
	"path"
	"runtime"
	"sort"
	"sync"

	"github.com/mdaliyan/gondex/cmd/internal/index"
)

var wg = sync.WaitGroup{}

var queue = make(chan *index.Node, 1000)

var LogError = log.New(os.Stderr, "", log.Ltime)

var pool = sync.Pool{New: func() interface{} { return &index.Node{} }}

func Generate(Path string) {
	for i := 0; i < runtime.NumCPU(); i++ {
		go Processor()
	}
	wg.Add(1)
	go putToQueue(&index.Node{Path: Path})
	wg.Wait()
}

func NewNode() *index.Node {
	return pool.Get().(*index.Node)
}

func putToQueue(n *index.Node)  {
	queue <- n
}

func Processor() {
	for true {
		ProcessDirectory(<-queue)
	}
}

func ProcessDirectory(node *index.Node) {
	f, err := os.Open(node.Path)
	if err != nil {
		return
	}
	files, err := f.Readdir(-1)
	f.Close()
	if err != nil {
		LogError.Fatal(err)
	}
	sort.Slice(files, func(i, j int) bool {
		if (files[i].IsDir() && files[j].IsDir()) || (!files[i].IsDir() && !files[j].IsDir()){
			return files[i].Name() < files[j].Name()
		}
		return files[i].IsDir() && ! files[j].IsDir()
	})
	for _, f := range files {
		if ! f.IsDir() {
			break
		}
		wg.Add(1)
		newNode := NewNode()
		*newNode = index.Node{Path: path.Join(node.Path, f.Name()), Level: node.Level + 1}
		go putToQueue(newNode)
	}
	node.Files = files

	index.Generate(node)

	pool.Put(node)
	wg.Done()
}
